package main

import (
	"bufio"
	"flag"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var (
	stdIn    bool
	filePath string
)

func writer(ws *websocket.Conn) {
	defer ws.Close()
	var r *bufio.Reader
	if stdIn {
		r = bufio.NewReader(os.Stdin)
	} else {
		f, _ := os.Open(filePath)
		r = bufio.NewReader(f)
		defer f.Close()
	}
	for {
		p, err := r.ReadBytes('\n')

		if err != io.EOF {
			ws.WriteMessage(websocket.TextMessage, p)
		}
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

	flag.Parse()
	if flag.NArg() < 1 {
		stdIn = true
	} else {
		filePath = flag.Args()[0]
	}

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	http.ListenAndServe(":8081", nil)
}
