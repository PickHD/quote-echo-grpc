package configs

import (
	"os"
	"strconv"
)

type (
	Config struct {
		ServerPort int
		DBHost     string
		DBUsername string
		DBPassword string
		DBPort     int
		DBName     string
	}
)

func NewConfig() *Config {
	serverPortInt, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	dbPortInt, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	return &Config{
		ServerPort: serverPortInt,
		DBHost:     os.Getenv("DB_HOST"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBPort:     dbPortInt,
		DBName:     os.Getenv("DB_NAME"),
	}
}
