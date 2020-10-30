package data

import (
	"fmt"
	"github.com/goandfootball/test-api/configs"
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

	migration, errEnv := configs.GetEnv("POSTGRES_MIGRATION")
	if errEnv != nil {
		fmt.Println(errEnv)
	}

	if migration == "true" {
		errMig := dbMigration(db)
		if errMig != nil {
			fmt.Println("error on migration database")
		}
	}

	data = &Data{
		Db: db,
	}
}

func New() *Data {
	once.Do(initDb)

	return data
}
