package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/goandfootball/test-api/configs"
	"github.com/goandfootball/test-api/internal/server"
)

// Welcolme handler
func Welcolme(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	/*
		err := godotenv.Load(os.ExpandEnv("C:/workspaces/Go/src/github.com/goandfootball/test-api/.env.development.local"))
		if err != nil {
			fmt.Println(err)
		}
	*/

	// new router
	r := mux.NewRouter()
	// Routers consist of a path and a handler function
	// Welcolme path
	r.HandleFunc("/", Welcolme).Methods("GET")

	port, err := configs.GetEnv("PORT") //os.LookupEnv("PORT")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(port)
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
