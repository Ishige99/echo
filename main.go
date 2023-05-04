package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", TestHandler).Methods(http.MethodGet) // Hello World!

	r.HandleFunc("/", GetHomeViewHandler).Methods(http.MethodGet)       // Home
	r.HandleFunc("/user", GetUserViewHandler).Methods(http.MethodGet)   // User details page
	r.HandleFunc("/log", GetLogFormViewHandler).Methods(http.MethodGet) // Log form page
	r.HandleFunc("/log", PostLogHandler).Methods(http.MethodPost)       // Log post

	log.Println("Server start at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r)) // Start server
}
