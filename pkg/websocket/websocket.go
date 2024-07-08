package websocket

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joskeiner/go-myChat/internal/entities"
)

type UpgradeConn struct {
	websocket.Upgrader
}

type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	broadcast chan *entities.Message
}

func NewUpgrade() UpgradeConn {
	return UpgradeConn{
		websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *entities.Message),
	}
}

func (s *WebSocketServer) HandlerWebsocket(conn *websocket.Conn) {
	s.clients[conn] = true

	defer func() {
		delete(s.clients, conn)
		conn.Close()
	}()

	for {
		_, text, err := conn.ReadMessage()
		log.Printf("value: %v", string(text))
		if err != nil {
			log.Panicf("Error Read %s", err)
			break
		}
		var message entities.Message
		if err := json.Unmarshal(text, &message); err != nil {
			log.Fatalln("Error converting: ")
		}
		s.broadcast <- &message
		log.Println(message)
	}
}

func (s *WebSocketServer) HandleMessage() {
	for {
		msg := <-s.broadcast
		for client := range s.clients {
			err := client.WriteMessage(websocket.TextMessage, getMessageTemple(msg))
			if err != nil {
				log.Printf("Write Error %v", err)
				client.Close()
				delete(s.clients, client)
			}
		}
	}
}

func getMessageTemple(msg *entities.Message) []byte {
	basePath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	tmpl, err := template.ParseFiles(basePath + "/internal/static/views/message.html")
	if err != nil {
		log.Fatalf("template parsing : %s", err)
	}
	var renderedMessage bytes.Buffer
	err = tmpl.Execute(&renderedMessage, msg)
	if err != nil {
		log.Fatalf("template execution :%s", err)
	}
	return renderedMessage.Bytes()
}
