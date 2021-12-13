package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigType string

const (
	Dev  ConfigType = "development"
	Test            = "test"
	Prod            = "production"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     int
	DB_HOST     string
	DB_NAME     string
}

// get configurations default with 'development' environment
func GetConfig(configType ...ConfigType) Configuration {
	configuration := Configuration{}
	env := Dev
	if len(configType) > 0 {
		env = configType[0]
	}
	fileName := fmt.Sprintf("./config.%s.json", env)

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&configuration); err != nil {
		panic(err)
	}
	return configuration
}
