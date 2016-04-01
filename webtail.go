package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func writer(ws *websocket.Conn) {
	ws.WriteMessage(websocket.TextMessage, []byte("hello"))
	ws.Close()
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	go writer(ws)
}

func main() {
	http.HandleFunc("/ws", serveWs)
	http.ListenAndServe(":8081", nil)
}
