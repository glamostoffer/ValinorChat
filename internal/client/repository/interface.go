package repository

import (
	"context"
	"github.com/glamostoffer/ValinorChat/internal/model"
)

type Repository interface {
	Room
	Message
}

type Room interface {
	CreateRoom(ctx context.Context, hostID int64, name string) (roomID int64, err error)
	GetRooms(ctx context.Context, hostID int64) (rooms []model.Room, err error)
	AddClientToRoom(ctx context.Context, clientID int64, roomID int64) (err error)
	RemoveClientFromRoom(ctx context.Context, clientID int64, roomID int64) (err error)
}

type Message interface {
	CreateMessage(ctx context.Context, clientID, roomID int64, message string) (err error)
	GetMessagesFromRoom(ctx context.Context, roomID int64) ([]model.Message, error)
}
