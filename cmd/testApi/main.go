package main

import (
	"fmt"
	"github.com/goandfootball/test-api/configs"
	"github.com/goandfootball/test-api/internal/data"
	"github.com/goandfootball/test-api/internal/server"
	"log"
)

func main() {
	port, err := configs.GetEnv("SERVER_PORT")
	if err != nil {
		fmt.Println(err)
	}

	serv, err := server.New(port)
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

	err = serv.Start()
	if err != nil {
		fmt.Println(err)
	}
}
