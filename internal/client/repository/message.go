package repository

import (
	"context"
	"github.com/glamostoffer/ValinorChat/internal/model"
	"github.com/glamostoffer/ValinorChat/pkg/errlist"
	"log/slog"
	"time"
)

func (r *ClientRepository) CreateMessage(ctx context.Context, clientID, roomID int64, message string) (err error) {
	log := r.log.With(slog.String("op", "message_repo.CreateMessage"))

	result, err := r.db.ExecContext(ctx, queryCreateMessage, roomID, clientID, message, time.Now().Unix())
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

func (r *ClientRepository) GetMessagesFromRoom(ctx context.Context, roomID int64) (messages []model.Message, err error) {
	log := r.log.With(slog.String("op", "message_repo.GetMessagesFromRoom"))

	messages = make([]model.Message, 0)

	err = r.db.SelectContext(ctx, &messages, queryGetMessagesFromRoom, roomID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return messages, nil
}

func (r *ClientRepository) GetAllMessages(ctx context.Context) (messages []model.Message, err error) {
	log := r.log.With(slog.String("op", "message_repo.GetAllMessages"))

	messages = make([]model.Message, 0)

	err = r.db.SelectContext(ctx, &messages, queryGetMessages)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return messages, nil
}
