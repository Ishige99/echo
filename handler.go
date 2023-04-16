package main

import (
	"log"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hit /hello")

	var response = []byte("Hello World!")
	w.Write(response)
}
