package data

import (
	"log"
	"sync"

	"gorm.io/gorm"
)

var (
	data *Data
	once sync.Once
)

// Data manages the connection to the database.
type Data struct {
	Db *gorm.DB
}

func initDb() {
	db, err := getConnection()
	if err != nil {
		log.Panic(err)
	}

	data = &Data{
		Db: db,
	}
}

func New() *Data {
	once.Do(initDb)

	return data
}
