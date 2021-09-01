package main

import (
	"fmt"
	"github.com/sergeygvozdev08101993/go-js-chatapp/pkg/websocket"
	"net/http"
)

// serveWs устанавливает соединение по протоколу WebSocket.
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {

	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: ws,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

// setupRoutes определяет маршруты и соответствующие им обработчики.
func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("chat app is running...")
	setupRoutes()
	http.ListenAndServe(":3030", nil)
}
