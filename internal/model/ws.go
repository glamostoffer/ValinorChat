package model

import "github.com/gofiber/websocket/v2"

type (
	RegisterConnectionRequest struct {
		UserID int
		Conn   *websocket.Conn
	}
	UnregisterConnectionRequest struct {
		UserID int
		Conn   *websocket.Conn
	}
)
