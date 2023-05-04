package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", TestHandler).Methods(http.MethodGet) // Hello World!

	r.HandleFunc("/", GetIndexViewHandler).Methods(http.MethodGet)    // Home
	r.HandleFunc("/user", GetUserViewHandler).Methods(http.MethodGet) // User details page
	r.HandleFunc("/post", GetPostViewHandler).Methods(http.MethodGet) // Post page

	log.Println("Server start at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r)) // Start server
}
