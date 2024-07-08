package main

import (
	"log"
	"net/http"

	"github.com/joskeiner/go-myChat/internal/config"
	"github.com/joskeiner/go-myChat/internal/server"
	"github.com/joskeiner/go-myChat/pkg/websocket"
)

func main() {
	run()
}

// this function will execute the server and the dependecies
func run() {
	router := http.NewServeMux()
	adr, basePath := config.LoandingDeps()

	upgreade := websocket.NewUpgrade()
	server := server.NewServer(adr, router)

	connWebSocket := websocket.NewWebSocketServer()
	// tmpl := template.Must(template.ParseFiles(basePath + "/internal/static/views/index.html"))
	// server static files
	// fs := http.FileServer(http.Dir(basePath + "/internal/static/"))

	// router test
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, basePath+"/internal/static/views/index.html")
	})

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgreade.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
		}
		connWebSocket.HandlerWebsocket(conn)
	})
	go connWebSocket.HandleMessage()
	server.Start()
}
