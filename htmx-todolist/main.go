package main

import (
	"log"
	"net/http"
	"text/template"
)

type Todo struct {
	id        int
	userId    int
	Title     string
	completed bool
}

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		todos := map[string][]Todo{
			"Todos": {
				{
					id:        1,
					userId:    1,
					Title:     "delectus aut autem",
					completed: false,
				},
				{
					id:        2,
					userId:    1,
					Title:     "quis ut nam facilis et officia qui",
					completed: false,
				},
			},
		}

		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, todos)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
