package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/health", HealthHandler).Methods("GET")

	http.ListenAndServe(":8080", mux)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
