package usecase

import (
	"github.com/glamostoffer/ValinorChat/internal/client/repository"
	"github.com/glamostoffer/ValinorChat/pkg/tx_manager"
	"log/slog"
)

type ClientUseCase struct {
	repo repository.Repository
	tx   *tx_manager.TxManager
	log  *slog.Logger
}

func New(
	repo *repository.ClientRepository,
	tx *tx_manager.TxManager,
	log *slog.Logger,
) *ClientUseCase {
	uc := &ClientUseCase{
		repo: repo,
		tx:   tx,
		log:  log,
	}

	return uc
}
