package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Password() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env File!!")
	}

	return os.Getenv("DB_PASSWORD")
}

func Host() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env File!!")
	}

	return os.Getenv(("DB_HOST"))
}

func Port() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env File!!")
	}

	return os.Getenv(("DB_PORT"))
}

func DbName() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env File!!")
	}

	return os.Getenv(("DB_NAME"))
}

func Username() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env File!!")
	}

	return os.Getenv(("DB_USERNAME"))
}
