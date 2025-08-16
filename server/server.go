package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	repository "ecommerce/go/chatbot/Repository"
	"ecommerce/go/chatbot/database"
	"ecommerce/go/chatbot/handlers"
	middlewareLocal "ecommerce/go/chatbot/middleware"
)

type Config struct {
	Port        string
	DatabaseUrl string
}

func NewServer(config *Config) {

	StartDB(config.DatabaseUrl)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(18 * time.Second))
	r.Use(middlewareLocal.CORS())

	r.Method("POST", "/chat", handlers.Handler(handlers.ListProductsHandler))

	fmt.Println("Starting server on port:", config.Port)
	http.ListenAndServe(config.Port, r)

}

func StartDB(DatabaseUrl string) {

	repo, err := database.NewPostgresRepository(DatabaseUrl)

	if err != nil {
		log.Fatal("Error with database connection", err)
	}

	repository.SetRepository(repo)
}
