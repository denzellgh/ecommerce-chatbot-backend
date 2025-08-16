package main

import (
	"ecommerce/go/chatbot/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	if PORT == "" {
		log.Fatal("port is required")
	}

	if DATABASE_URL == "" {
		log.Fatal("database url is required")
	}

	server.NewServer(&server.Config{
		Port:        PORT,
		DatabaseUrl: DATABASE_URL,
	})
}
