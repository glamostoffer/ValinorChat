package repository

import (
	"context"
	"errors"
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

	rows, err := r.db.QueryContext(ctx, queryGetMessagesFromRoom, roomID)
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
		message := model.Message{}

		if scanErr := rows.Scan(&message); scanErr != nil {
			log.Error("can't scan value from row to room model", err.Error())
			return nil, scanErr
		}

		messages = append(messages, message)
	}

	return messages, nil
}
