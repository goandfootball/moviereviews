package data

import (
	"fmt"
	"io/ioutil"

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

	dsn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v", host, port, user, dbName, password, sslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func dbMigration(db *gorm.DB) error {
	// 202010312232 TODO: fix migration error
	envUsersTable, errUsersTable := configs.GetEnv("TEST_USERS_TABLE_PATH")
	if errUsersTable != nil {
		return errUsersTable
	}
	envDummyData, errDummy := configs.GetEnv("TEST_USERS_DUMMY_DATA")
	if errDummy != nil {
		return errDummy
	}

	bUsersTable, errReadUsers := ioutil.ReadFile(envUsersTable)
	if errReadUsers != nil {
		return errReadUsers
	}

	bDummyData, errReadDummy := ioutil.ReadFile(envDummyData)
	if errReadDummy != nil {
		return errReadDummy
	}

	sqlDb, errDB := db.DB()
	if errDB != nil {
		return errDB
	}

	rowsUsrTab, errQueryUsers := sqlDb.Query(string(bUsersTable))
	if errQueryUsers != nil {
		return errQueryUsers
	}
	rowsDummyDat, errQueryDummy := sqlDb.Query(string(bDummyData))
	if errQueryDummy != nil {
		return errQueryDummy
	}

	errCloUsers := rowsUsrTab.Close()
	if errCloUsers != nil {
		return errCloUsers
	}
	errCloDummy := rowsDummyDat.Close()
	if errCloDummy != nil {
		return errCloDummy
	}

	return nil
}
