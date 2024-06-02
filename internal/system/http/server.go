package http

import (
	"context"
	"github.com/glamostoffer/ValinorChat/internal/client/delivery/ws"
	"github.com/glamostoffer/ValinorChat/pkg/errlist"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"log"
	"log/slog"
	"time"
)

type FiberServer struct {
	cfg       Config
	fiber     *fiber.App
	logger    *slog.Logger
	wsHandler ws.Handler
}

func New(
	cfg Config,
	logger *slog.Logger,
	wsHandlers ws.Handler,
) *FiberServer {
	return &FiberServer{
		cfg: cfg,
		fiber: fiber.New(fiber.Config{
			BodyLimit:             100 * 1024,
			ReadBufferSize:        100 * 1024,
			WriteBufferSize:       100 * 1024,
			DisableStartupMessage: true,
		}),
		logger:    logger,
		wsHandler: wsHandlers,
	}
}

func (f *FiberServer) Start(_ context.Context) error {
	f.fiber.Use(cors.New(cors.Config{
		AllowOrigins:     f.cfg.AllowOrigins,
		AllowCredentials: true,
		AllowHeaders:     f.cfg.AllowHeaders,
		ExposeHeaders:    f.cfg.ExposeHeaders,
	}))

	if f.cfg.PProfEnabled {
		f.fiber.Use(pprof.New())
	}

	f.mapHandlers()

	go func() {
		if err := f.fiber.Listen(f.cfg.Host + ":" + f.cfg.Port); err != nil {
			log.Fatal(err.Error())
		}
	}()

	return nil
}

func (f *FiberServer) Stop(_ context.Context) error {
	errCh := make(chan error)
	go func() {
		err := f.fiber.Shutdown()
		errCh <- err
	}()

	select {
	case err := <-errCh:
		return err
	case <-time.After(f.cfg.StopTimeout):
		return errlist.ErrStopTimeout
	}
}

func (f *FiberServer) mapHandlers() {
	router := f.fiber.Group("websocket")
	ws.MapWsHandlers(router, f.wsHandler)
}
