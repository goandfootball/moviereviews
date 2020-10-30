package data

import (
	"fmt"

	"github.com/goandfootball/test-api/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getConnection() (*gorm.DB, error) {
	var (
		//dialect  string
		host     string
		user     string
		password string
		dbName   string
		port     string
		sslMode  string
		//timeZone string

		err error
	)
	/*
		dialect, err = configs.GetEnv("POSTGRES_DIALECT")
		if err != nil {
			fmt.Println(err)
		}
	*/
	host, err = configs.GetEnv("POSTGRES_HOST")
	if err != nil {
		fmt.Println(err)
	}
	user, err = configs.GetEnv("POSTGRES_USER")
	if err != nil {
		fmt.Println(err)
	}
	password, err = configs.GetEnv("POSTGRES_PASSWORD")
	if err != nil {
		fmt.Println(err)
	}
	dbName, err = configs.GetEnv("POSTGRES_DBNAME")
	if err != nil {
		fmt.Println(err)
	}
	port, err = configs.GetEnv("POSTGRES_PORT")
	if err != nil {
		fmt.Println(err)
	}
	sslMode, err = configs.GetEnv("POSTGRES_SSLMODE")
	if err != nil {
		fmt.Println(err)
	}
	/*
		timeZone, err = configs.GetEnv("TIMEZONE")
		if err != nil {
			fmt.Println(err)
		}
	*/

	dsn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v", host, port, user, dbName, password, sslMode)
	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
