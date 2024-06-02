package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorChat/internal/model"
	authproto "github.com/glamostoffer/ValinorProtos/auth/client_auth"
)

func (uc *ClientUseCase) SaveMessage(ctx context.Context, message model.Message) (err error) {
	return uc.repo.CreateMessage(ctx, message.ClientID, message.RoomID, message.Content)
}

func (uc *ClientUseCase) GetMessages(ctx context.Context, roomID int64) (messages []model.Message, err error) {
	messages, err = uc.repo.GetMessagesFromRoom(ctx, roomID)
	if err != nil {
		return nil, err
	}

	for i, msg := range messages {
		details, err := uc.auth.ClientAuth.GetClientDetails(ctx, &authproto.GetClientDetailsRequest{
			ClientID: msg.ClientID,
		})
		if err != nil {
			return nil, err
		}

		messages[i].Username = details.GetUsername()
	}

	return messages, nil
}

func (uc *ClientUseCase) GetAllMessages(ctx context.Context) (messages []model.Message, err error) {
	return uc.repo.GetAllMessages(ctx)
}
