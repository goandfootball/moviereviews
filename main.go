package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/goandfootball/test-api/internal/server"
)

// Welcolme handler
func Welcolme(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	port := os.Getenv("PORT")

	// new router
	r := mux.NewRouter()
	// Routers consist of a path and a handler function
	// Welcolme path
	r.HandleFunc("/", Welcolme).Methods("GET")

	serv, err := server.New(port, r)
	if err != nil {
		fmt.Println(err)
	}

	// start the server.
	err = serv.Start()
	if err != nil {
		fmt.Println(err)
	}
}
