package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/goandfootball/test-api/configs"
	"github.com/goandfootball/test-api/internal/data"
	"github.com/goandfootball/test-api/internal/server"
)

// Welcome handler
func Welcome(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, world!")
	if err != nil {
		fmt.Println("error trying to print greet")
	}
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
	// Welcome path
	r.HandleFunc("/", Welcome).Methods("GET")

	port, err := configs.GetEnv("SERVER_PORT") //os.LookupEnv("PORT")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(port)
	serv, err := server.New(port, r)
	if err != nil {
		fmt.Println(err)
	}

	// connection to the database.
	d := data.New()
	sqlDB, errDB := d.Db.DB()
	if errDB != nil {
		fmt.Println(errDB)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
	/*
		if err := d.Db.Ping(); err != nil {
			log.Fatal(err)
		}
	*/
	// start the server.
	err = serv.Start()
	if err != nil {
		fmt.Println(err)
	}
}
