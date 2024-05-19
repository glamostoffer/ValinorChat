package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorChat/internal/model"
)

type UseCase interface {
	//Message
	Room
}

type Message interface {
}

type Room interface {
	CreateRoom(ctx context.Context, roomName string, ownerID int64) (roomID int64, err error)
	GetListOfRooms(ctx context.Context, clientID int64) (rooms []model.Room, err error)
	AddClientToRoom(ctx context.Context, roomID, clientID int64) (err error)
	RemoveClientFromRoom(ctx context.Context, roomID, clientID int64) (err error)
}
