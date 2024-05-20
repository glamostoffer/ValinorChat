package ws

import (
	"context"
	"github.com/glamostoffer/ValinorChat/internal/model"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan *model.Message
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan *model.Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Start(_ context.Context) error {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *Hub) Stop(_ context.Context) error {
	for client := range h.clients {
		delete(h.clients, client)
		close(client.send)
	}

	return nil
}
