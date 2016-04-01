package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func writer(ws *websocket.Conn) {
	defer ws.Close()

	for {
		ws.WriteMessage(websocket.TextMessage, []byte("hello"))
		time.Sleep(2 * time.Second)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, _ := template.ParseFiles("webtail.html")
	var v = struct {
		Host string
	}{
		r.Host,
	}
	t.Execute(w, &v)
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	go writer(ws)
}

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	http.ListenAndServe(":8081", nil)
}
