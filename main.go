package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

  r.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets"))).Methods("GET")

	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}
