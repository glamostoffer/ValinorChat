package usecase

import (
	"github.com/glamostoffer/ValinorChat/internal/client/repository"
	"github.com/glamostoffer/ValinorChat/pkg/tx_manager"
	authclient "github.com/glamostoffer/ValinorProtos/auth"
	"log/slog"
)

type ClientUseCase struct {
	repo repository.Repository
	tx   *tx_manager.TxManager
	log  *slog.Logger
	auth *authclient.Connector
}

func New(
	repo *repository.ClientRepository,
	tx *tx_manager.TxManager,
	log *slog.Logger,
	auth *authclient.Connector,
) *ClientUseCase {
	uc := &ClientUseCase{
		repo: repo,
		tx:   tx,
		log:  log,
		auth: auth,
	}

	return uc
}
