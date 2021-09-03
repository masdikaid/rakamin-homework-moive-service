package config

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
)

type Config struct {
	AppName     string
	AppPort     int
	LogLevel    string
	Environment string
	JWTSecret   string
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Name     string
}

func Init() *Config {
	defaultEnv := ".env"

	if err := gotenv.Load(defaultEnv); err != nil {
		log.Fatal("failed load .env")
	}

	log.SetOutput(os.Stdout)

	appConfig := &Config{
		AppName:     GetString("APP_NAME"),
		AppPort:     GetInt("APP_PORT"),
		LogLevel:    GetString("LOG_LEVEL"),
		Environment: GetString("ENVIRONMENT"),
		JWTSecret:   GetString("JWT_SECRET"),
		DB_Username: GetString("DB_USERNAME"),
		DB_Password: GetString("DB_PASSWORD"),
		DB_Host:     GetString("DB_HOST"),
		DB_Port:     GetString("DB_PORT"),
		DB_Name:     GetString("DB_NAME"),
	}

	return appConfig
}
