package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorChat/internal/model"
)

type UseCase interface {
	Message
	Room
}

type Message interface {
	SaveMessage(ctx context.Context, message model.Message) (err error)
	GetMessages(ctx context.Context, roomID int64) (messages []model.Message, err error)
	GetAllMessages(ctx context.Context) (messages []model.Message, err error)
}

type Room interface {
	CreateRoom(ctx context.Context, roomName string, ownerID int64) (roomID int64, err error)
	GetListOfRooms(ctx context.Context, clientID int64) (rooms []model.Room, err error)
	AddClientToRoom(ctx context.Context, roomID, clientID int64) (err error)
	RemoveClientFromRoom(ctx context.Context, roomID, clientID int64) (err error)
}
