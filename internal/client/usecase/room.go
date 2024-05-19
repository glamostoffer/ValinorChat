package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorChat/internal/model"
)

func (uc *ClientUseCase) CreateRoom(ctx context.Context, roomName string, ownerID int64) (roomID int64, err error) {
	// мда...

	return uc.repo.CreateRoom(ctx, ownerID, roomName)
}

func (uc *ClientUseCase) GetListOfRooms(ctx context.Context, clientID int64) (rooms []model.Room, err error) {

	return uc.repo.GetRooms(ctx, clientID)
}

func (uc *ClientUseCase) AddClientToRoom(ctx context.Context, roomID, clientID int64) (err error) {

	return uc.repo.AddClientToRoom(ctx, clientID, roomID)
}

func (uc *ClientUseCase) RemoveClientFromRoom(ctx context.Context, roomID, clientID int64) (err error) {

	return uc.repo.RemoveClientFromRoom(ctx, roomID, clientID)
}
