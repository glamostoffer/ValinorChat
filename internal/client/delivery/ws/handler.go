package ws

import (
	"context"
	"encoding/json"
	"github.com/glamostoffer/ValinorChat/internal/client/usecase"
	"github.com/glamostoffer/ValinorChat/internal/model"
	"github.com/glamostoffer/ValinorChat/internal/system/wsmanager"
	"github.com/glamostoffer/ValinorChat/pkg/constants"
	authclient "github.com/glamostoffer/ValinorProtos/auth"
	authproto "github.com/glamostoffer/ValinorProtos/auth/client_auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/tidwall/sjson"
	"log"
	"time"
)

type Handler struct {
	manager *wsmanager.Manager
	auth    *authclient.Connector
	uc      usecase.UseCase
}

func New(
	manager *wsmanager.Manager,
	auth *authclient.Connector,
	uc usecase.UseCase,
) *Handler {
	return &Handler{
		manager: manager,
		auth:    auth,
		uc:      uc,
	}
}

func (h *Handler) ClientMiddleware(c *fiber.Ctx) (err error) {
	ctx := c.Context()

	authToken := c.Params("token")
	session, err := h.auth.ClientAuth.ClientAuth(
		ctx,
		&authproto.ClientAuthRequest{
			AccessToken: authToken,
		},
	)
	if err != nil {
		return
	}

	c.Locals(constants.UserIDKey, session.UserID)

	return c.Next()
}

func (h *Handler) HandleWsClientChatConnection(c *websocket.Conn) {
	ctx := context.Background()

	var userID int
	var err error

	defer func() {
		if err != nil {
			_ = c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		}

		h.manager.UnregisterConnection(
			model.UnregisterConnectionRequest{
				UserID: userID,
				Conn:   c,
			},
		)
	}()

	user, ok := c.Locals(constants.UserIDKey).(int64)
	if !ok {
		return
	}

	userID = int(user)

	h.manager.RegisterConnection(
		model.RegisterConnectionRequest{
			UserID: userID,
			Conn:   c,
		},
	)

	c.SetPingHandler(func(appData string) error {
		return nil
	})

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			return
		}

		dto := model.MessageDTO{}

		err = json.Unmarshal(message, &dto)
		if err != nil {
			return
		}

		msg := model.Message{
			ClientID: user,
			RoomID:   dto.RoomID,
			Content:  dto.Content,
			SentAt:   time.Now().Unix(),
		}

		err = h.uc.SaveMessage(context.Background(), msg)
		if err != nil {
			log.Printf("failed to save message: %s", err.Error())
		}

		write, err := json.Marshal(msg)

		details, err := h.auth.ClientAuth.GetClientDetails(ctx, &authproto.GetClientDetailsRequest{
			ClientID: msg.ClientID,
		})

		write, err = sjson.SetBytes(write, "username", details.GetUsername())

		if err != nil {
			log.Printf("failed to marshal message: %s", err.Error())
		}

		err = h.manager.Broadcast(write)
		if err != nil {
			log.Printf("failed to broadcast message: %s", err.Error())
		}

	}
}
