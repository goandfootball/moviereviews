package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func welcolme(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world!")
}

func Router() http.Handler {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", welcolme).Methods("GET")

	return r
}

func main() {
	http.ListenAndServe("localhost:8080", Router())
}
