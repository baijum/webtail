package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	ws.WriteMessage(websocket.TextMessage, []byte("Hello"))
	ws.Close()
}

func main() {
	http.HandleFunc("/ws", serveWs)
	http.ListenAndServe(":8081", nil)
}
