package wsmanager

import (
	"errors"
	"fmt"
	"github.com/glamostoffer/ValinorChat/internal/model"
	"github.com/gofiber/websocket/v2"
	"log/slog"
	"sync"
)

type Manager struct {
	connections map[int][]*websocket.Conn
	mu          *sync.RWMutex
	log         *slog.Logger
}

func New(log *slog.Logger) *Manager {
	return &Manager{
		connections: make(map[int][]*websocket.Conn),
		mu:          &sync.RWMutex{},
		log:         log,
	}
}

func (ws *Manager) RegisterConnection(request model.RegisterConnectionRequest) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	ws.connections[request.UserID] = append(ws.connections[request.UserID], request.Conn)

	return
}

func (ws *Manager) UnregisterConnection(request model.UnregisterConnectionRequest) {
	log := ws.log.With(slog.String("op", "ws.UnregisterConnection"))

	ws.mu.Lock()
	defer ws.mu.Unlock()

	if connections, ok := ws.connections[request.UserID]; ok {
		for i, conn := range connections {
			if conn == request.Conn {
				err := conn.Close()
				if err != nil {
					log.Error(err.Error())
				}

				ws.connections[request.UserID] = append(connections[:i], connections[i+1:]...)

				break
			}
		}
	}
}

func (ws *Manager) Broadcast(message []byte) (joinedErr error) {
	log := ws.log.With(slog.String("op", "ws.Broadcast"))

	ws.mu.Lock()
	defer ws.mu.Unlock()

	for key := range ws.connections {
		for _, conn := range ws.connections[key] {
			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Error(err.Error())

				err = errors.New(
					fmt.Sprintf("%s; ws.writeMessage(message: %s)",
						err.Error(),
						string(message),
					),
				)

				joinedErr = errors.Join(joinedErr, err)
				continue
			}
		}
	}

	return joinedErr
}

func (ws *Manager) SendByIDWebsocketMessage(userID int, message []byte) (joinedErr error) {
	log := ws.log.With(slog.String("op", "ws.SendByIDWebsocketMessage"))

	ws.mu.Lock()
	defer ws.mu.Unlock()

	if connections, ok := ws.connections[userID]; ok {
		for _, conn := range connections {
			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Error(err.Error())

				err = errors.New(
					fmt.Sprintf("%s; ws.writeMessage(message: %s)",
						err.Error(),
						string(message),
					),
				)

				joinedErr = errors.Join(joinedErr, err)

				continue
			}
		}

	}

	return joinedErr
}
