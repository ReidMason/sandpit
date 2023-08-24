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
	http.HandleFunc("/", getIndex)

	http.HandleFunc("/todos", getTodos)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getIndex(w http.ResponseWriter, _ *http.Request) {
	templ := template.Must(template.ParseFiles("index.html"))
	templ.Execute(w, nil)
}

func getTodos(w http.ResponseWriter, _ *http.Request) {
	todos := []Todo{
		{
			id:        1,
			userId:    1,
			Title:     "<script>alert('test')</script>",
			completed: false,
		},
		{
			id:        2,
			userId:    1,
			Title:     "quis ut nam facilis et officia qui",
			completed: false,
		},
	}

	templ := template.Must(template.ParseFiles("templates/todos.html"))
	err := templ.Execute(w, todos)
	if err != nil {
		log.Println(err)
	}
}
