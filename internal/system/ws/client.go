package ws

import (
	"context"
	"github.com/glamostoffer/ValinorChat/internal/model"
	"github.com/gorilla/websocket"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan *model.Message
}

func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	return &Client{
		hub:  hub,
		conn: conn,
		send: make(chan *model.Message),
	}
}

func (c *Client) Start(_ context.Context) error {
	go c.writePump()
	go c.readPump()

	return nil
}

func (c *Client) Stop(_ context.Context) error {
	return c.conn.Close()
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// log the error if it's not a normal closure
			}
			break
		}
		c.hub.broadcast <- &model.Message{Content: string(message)}
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.conn.WriteJSON(message)
			if err != nil {
				// log the error
				return
			}
		}
	}
}
