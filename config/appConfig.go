package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
	AppSecret  string
}

func SetUpEnv() (cfg AppConfig, err error) {

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("HTTP_PORT")

	print("==asd")
	print(httpPort)
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env varaiables not found")
	}

	Dsn := os.Getenv("DSN")
	if len(Dsn) < 1 {
		return AppConfig{}, errors.New("env varaiables not found")
	}

	log.Println("Database: %v\n", Dsn)

	AppSecret := os.Getenv("APPSECRET")
	if len(AppSecret) < 1 {
		return AppConfig{}, errors.New("app secret not found")
	}

	return AppConfig{
			ServerPort: httpPort,
			Dsn:        Dsn,
			AppSecret:  AppSecret,
		},
		nil

}
