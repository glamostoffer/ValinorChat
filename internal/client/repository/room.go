package repository

import (
	"context"
	"errors"
	"github.com/glamostoffer/ValinorChat/internal/model"
	"github.com/glamostoffer/ValinorChat/pkg/errlist"
	"github.com/lib/pq"
	"log/slog"
)

func (r *ClientRepository) CreateRoom(
	ctx context.Context,
	hostID int64,
	name string,
) (roomID int64, err error) {
	log := r.log.With(slog.String("op", "room_repo.CreateRoom"))

	err = r.db.GetContext(ctx, &roomID, queryCreateRoom, name, hostID)
	if err != nil {
		log.Error("failed to create room", err.Error())
		return -1, err
	}

	return roomID, nil
}

func (r *ClientRepository) GetRooms(ctx context.Context, hostID int64) (rooms []model.Room, err error) {
	log := r.log.With(slog.String("op", "room_repo.GetRooms"))

	rows, err := r.db.QueryContext(ctx, queryGetRooms, hostID)
	defer func() {
		if rows != nil {
			err = errors.Join(err, rows.Close())
		}
	}()

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	for rows.Next() {
		room := model.Room{}
		clientIDs := pq.Int64Array{}

		if scanErr := rows.Scan(&room.ID, &room.Name, &room.OwnerID, &clientIDs); scanErr != nil {
			log.Error("can't scan value from row to room model", err.Error())
			return nil, scanErr
		}

		room.ClientIDs = clientIDs

		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r *ClientRepository) AddClientToRoom(ctx context.Context, clientID int64, roomID int64) (err error) {
	log := r.log.With(slog.String("op", "room_repo.AddClientToRoom"))

	result, err := r.db.ExecContext(ctx, queryAddClientToRoom, clientID, roomID)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	if rows != 1 {
		log.Error(errlist.InvalidAffectedRowsCount)
		return errlist.ErrInvalidAffectedRowsCount
	}

	return nil
}

func (r *ClientRepository) RemoveClientFromRoom(ctx context.Context, clientID int64, roomID int64) (err error) {
	log := r.log.With(slog.String("op", "room_repo.RemoveClientFromRoom"))

	result, err := r.db.ExecContext(ctx, queryRemoveClientFromRoom, clientID, roomID)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	if rows != 1 {
		log.Error(errlist.InvalidAffectedRowsCount)
		return errlist.ErrInvalidAffectedRowsCount
	}

	return nil
}
