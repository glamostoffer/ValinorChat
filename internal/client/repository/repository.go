package repository

import (
	"github.com/glamostoffer/ValinorChat/pkg/pg_connector"
	"log/slog"
)

type ClientRepository struct {
	db  *pg_connector.Connector
	log *slog.Logger
}

func New(db *pg_connector.Connector, log *slog.Logger) *ClientRepository {
	pg := &ClientRepository{
		db:  db,
		log: log,
	}

	return pg
}
