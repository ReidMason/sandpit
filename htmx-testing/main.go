package main

import (
	"log"
	"net/http"
	"text/template"
)

var classes = ""

func main() {
	http.HandleFunc("/", getIndex)

	http.HandleFunc("/setTheme1", setTheme1)
	http.HandleFunc("/getThing", getThing)
	// http.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
	// 	serveWs(w, r)
	// })

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func setTheme1(w http.ResponseWriter, _ *http.Request) {
	classes = "bg-blue-900 rounded p-4"
	w.Header().Set("HX-Trigger", "theme-update")
	w.WriteHeader(http.StatusOK)
}

func getThing(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("<div>Hello!</div>"))
}

func getIndex(w http.ResponseWriter, _ *http.Request) {
	templ := template.Must(template.ParseFiles("index.html"))
	templ.Execute(w, nil)
}
