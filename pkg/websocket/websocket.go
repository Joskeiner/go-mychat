package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/joskeiner/go-myChat/internal/entities"
)

type webSocketServer struct {
	clients   map[*websocket.Conn]bool
	broadcast chan *entities.Message
}

func NewWebSocketServer() *webSocketServer {
	return &webSocketServer{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *entities.Message),
	}
}
