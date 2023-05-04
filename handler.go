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

func GetHomeViewHandler(w http.ResponseWriter, r *http.Request) {
	homePageResponse, err := template.ParseFiles("./public/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = homePageResponse.Execute(w, nil) // index.htmlの実行
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUserViewHandler(w http.ResponseWriter, r *http.Request) {
	userPageResponse, err := template.ParseFiles("./public/user.html")
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

func GetLogFormViewHandler(w http.ResponseWriter, r *http.Request) {
	logFormPageResponse, err := template.ParseFiles("./public/log_form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = logFormPageResponse.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostLogHandler(w http.ResponseWriter, r *http.Request) {
	var printResponse = []byte("form post done.")
	w.Write(printResponse)
}
