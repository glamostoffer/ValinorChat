package tx_manager

import (
	"context"
	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/glamostoffer/ValinorChat/pkg/pg_connector"
)

type TxManager struct {
	db *pg_connector.Connector
	*manager.Manager
}

func New(db *pg_connector.Connector) *TxManager {
	return &TxManager{
		db: db,
	}
}

func (tx *TxManager) Start(ctx context.Context) error {
	tx.Manager = manager.Must(trmsqlx.NewDefaultFactory(tx.db.DB))

	return nil
}

func (tx *TxManager) Stop(ctx context.Context) error {
	return nil
}
