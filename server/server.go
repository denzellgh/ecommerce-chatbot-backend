package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"ecommerce/go/chatbot/ai"
	"ecommerce/go/chatbot/database"
	"ecommerce/go/chatbot/handlers"
	middlewareLocal "ecommerce/go/chatbot/middleware"
)

type Config struct {
	Port        string
	DatabaseUrl string
	OllamaUrl   string
	ModelName   string
}

func NewServer(config *Config) {

	db, err := database.StartPostgresDB(config.DatabaseUrl)

	if err != nil {
		log.Fatal("Error with database connection:", err)
	}
	defer db.Close()

	aiClient := ai.NewOllamaClient(config.OllamaUrl, config.ModelName)

	chatHandler := handlers.NewChatHandler(db, aiClient, config.ModelName)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(800 * time.Second))
	r.Use(middlewareLocal.CORS())

	// Routes
	r.Route("/api", func(r chi.Router) {
		r.Post("/chat", handlers.AdaptHandler(chatHandler.HandleChatQuery))
		r.Get("/health", handlers.AdaptHandler(chatHandler.HandleHealthCheck))
	})

	fmt.Println("Starting server on port:", config.Port)
	http.ListenAndServe(config.Port, r)
}
