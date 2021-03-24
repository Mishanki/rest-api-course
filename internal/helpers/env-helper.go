package helpers

import (
	"github.com/joho/godotenv"
	"log"
)

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
