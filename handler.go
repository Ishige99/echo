package main

import (
	"html/template"
	"log"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hit /hello")

	var response = []byte("Hello World!")
	w.Write(response)
}

func GetIndexViewHandler(w http.ResponseWriter, r *http.Request) {
	indexPageResponse, err := template.ParseFiles("./view/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = indexPageResponse.Execute(w, nil) // index.htmlの実行
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUserViewHandler(w http.ResponseWriter, r *http.Request) {
	userPageResponse, err := template.ParseFiles("./view/user.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = userPageResponse.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
