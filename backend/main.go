package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool { return true },
}

// reader получает сообщения передаваемые по WebSocket соединению,
// и передает это же сообщение обратно клиенту.
func reader(conn *websocket.Conn) {

	for {

		messageType, data, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(data))

		if err := conn.WriteMessage(messageType, data); err != nil {
			log.Println(err)
			return
		}
	}
}

// serveWs устанавливает соединение по протоколу WebSocket.
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	reader(ws)
}

// setupRoutes определяет маршруты и соответствующие им обработчики.
func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "simple server")
	})

	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("chat app is running...")
	setupRoutes()
	http.ListenAndServe(":3030", nil)
}
