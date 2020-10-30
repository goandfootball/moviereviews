package configs

import (
	"errors"

	"github.com/spf13/viper"
)

func GetEnv(key string) (string, error) {
	// Change path for development, testing or production configurationss
	viper.SetConfigFile("C:/workspaces/Go/src/github.com/goandfootball/test-api/development-local.env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		return "", err
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		return "", errors.New("not environment variable defined")
	}

	return value, nil
}
