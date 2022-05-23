package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	m := mux.NewRouter()

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`Welcome to server`))
	}).Methods(http.MethodGet)

	m.HandleFunc("/host", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		name, _ := os.Hostname()
		w.Write([]byte(name))
	}).Methods(http.MethodGet)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", m)
}
