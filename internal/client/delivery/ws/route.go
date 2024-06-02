package ws

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func MapWsHandlers(r fiber.Router, h Handler) {
	r.Get("/:token", h.ClientMiddleware, websocket.New(h.HandleWsClientChatConnection))
}
