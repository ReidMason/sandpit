package main

import (
	"bytes"
	"htmx-testing/internal/board"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var interval = 1000

var ws *websocket.Conn

type Todo struct {
	id        int
	userId    int
	Title     string
	completed bool
}

func main() {
	http.HandleFunc("/", getIndex)

	http.HandleFunc("/setInterval", setInterval)
	http.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r)
	})

	go sendStuff()

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func setInterval(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	intervalString := r.Form.Get("interval")
	newInterval, err := strconv.Atoi(intervalString)
	if err != nil {
		w.Write([]byte("Failed to update interval"))
		log.Println(err)
		return
	}

	interval = newInterval

	w.Write([]byte("Interval updated"))
}

func sendStuff() {
	var boardData *board.Board
	for {
		if ws != nil {
			if boardData == nil {
				boardData = board.New(100)
			}

			data := boardData.Display()
			// boardData.Iter()
			for boardData.Iter() {

			}
			// for i := 0; i < 5; i++ {
			// 	boardData.Iter()
			// }

			templ := template.Must(template.ParseFiles("templates/time.html"))
			w := bytes.NewBuffer(make([]byte, 0))

			err := templ.Execute(w, struct{ Data [][]board.TileDisplay }{Data: data})
			if err != nil {
				log.Println(err)
			}

			ws.WriteMessage(1, w.Bytes())
		}
		time.Sleep(time.Millisecond * time.Duration(interval))
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, _ = upgrader.Upgrade(w, r, nil)
}

func getIndex(w http.ResponseWriter, _ *http.Request) {
	templ := template.Must(template.ParseFiles("index.html"))
	templ.Execute(w, nil)
}
