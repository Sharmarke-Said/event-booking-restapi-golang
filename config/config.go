package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JwtSecret string
	AppEnv    string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AppConfig = Config{
		JwtSecret: os.Getenv("JWT_SECRET"),
		AppEnv:    os.Getenv("APP_ENV"),
	}
}