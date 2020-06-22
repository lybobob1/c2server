package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/implants", getImplantHandler).Methods("GET")
	r.HandleFunc("/callback", callbackHandler).Methods("GET", "POST")
	r.HandleFunc("/task/add", taskAddHandler).Methods("POST")

	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")
	return r
}
