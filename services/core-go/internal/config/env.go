package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err:= godotenv.Load()

	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

}

func GetPort() string {
	port:= os.Getenv("PORT")
	if port =="" {
		return "3000"
	}

	return port
}