package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", TestHandler).Methods(http.MethodGet) // Hello World!

	log.Println("Server start at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r)) // Start server
}
