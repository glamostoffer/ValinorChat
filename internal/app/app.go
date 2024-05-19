package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/glamostoffer/ValinorChat/internal/client/repository"
	"github.com/glamostoffer/ValinorChat/internal/client/usecase"
	"github.com/glamostoffer/ValinorChat/internal/config"
	"github.com/glamostoffer/ValinorChat/internal/system/grpc"
	"github.com/glamostoffer/ValinorChat/pkg/constants"
	"github.com/glamostoffer/ValinorChat/pkg/pg_connector"
	"github.com/glamostoffer/ValinorChat/pkg/tx_manager"

	"log/slog"
)

type (
	App struct {
		cfg        config.Config
		components []component
		log        *slog.Logger
	}
	component struct {
		Service Lifecycle
		Name    string
	}
	Lifecycle interface {
		Start(ctx context.Context) error
		Stop(ctx context.Context) error
	}
)

func New(cfg config.Config, logger *slog.Logger) *App {
	return &App{
		cfg: cfg,
		log: logger,
	}
}

func (a *App) Start(ctx context.Context) error {
	log := a.log.With(slog.String("op", "app.Start"))

	pg := pg_connector.New(a.cfg.Postgres)

	tx := tx_manager.New(pg)

	clientRepo := repository.New(pg, a.log)

	clientUC := usecase.New(clientRepo, tx, a.log)

	grpcServer := grpc.NewServer(a.cfg, clientUC)

	a.components = append(
		a.components,
		component{pg, "postgres"},
		component{tx, "tx manager"},
		component{grpcServer, "grpc server"},
	)

	okChan := make(chan struct{})
	errChan := make(chan error)

	go func() {
		var err error
		for _, c := range a.components {
			log.Info(constants.FmtStarting, slog.Any("name", c.Name))

			err = c.Service.Start(context.Background())
			if err != nil {
				log.Error(constants.FmtErrOnStarting, c.Name, err.Error())
				errChan <- errors.New(
					fmt.Sprintf("%s %s: %s", constants.FmtCannotStart, c.Name, err.Error()),
				)

				return
			}
		}
		okChan <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return errors.New("start timeout")
	case err := <-errChan:
		return err
	case <-okChan:
		log.Info("application started!")
		return nil
	}
}

func (a *App) Stop(ctx context.Context) error {
	log := a.log.With(slog.String("op", "app.Stop"))
	okChan := make(chan struct{})
	errChan := make(chan error)

	go func() {
		var err error
		for i := len(a.components) - 1; i >= 0; i-- {
			log.Info(
				constants.FmtStopping,
				slog.Any("name", a.components[i].Name),
			)

			err = a.components[i].Service.Stop(ctx)
			if err != nil {
				log.Error(constants.FmtErrOnStopping, a.components[i].Name, err.Error())
				errChan <- errors.New(
					fmt.Sprintf(
						"%s %s: %s",
						constants.FmtCannotStop,
						a.components[i].Name,
						err.Error(),
					),
				)

				return
			}
		}
		okChan <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return errors.New("stop timeout")
	case err := <-errChan:
		return err
	case <-okChan:
		log.Info("application stopped!")
		return nil
	}
}
