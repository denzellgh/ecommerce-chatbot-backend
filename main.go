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
	OLLAMA_URL := os.Getenv("OLLAMA_URL")
	MODEL_NAME := os.Getenv("MODEL_NAME")

	if PORT == "" {
		log.Fatal("port is required")
	}

	if DATABASE_URL == "" {
		log.Fatal("database url is required")
	}

	if OLLAMA_URL == "" {
		OLLAMA_URL = "http://localhost:11434"
	}

	if MODEL_NAME == "" {
		log.Fatal("model name is required")
	}

	server.NewServer(&server.Config{
		Port:        PORT,
		DatabaseUrl: DATABASE_URL,
		OllamaUrl:   OLLAMA_URL,
		ModelName:   MODEL_NAME,
	})
}
