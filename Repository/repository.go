package repository

import (
	"context"
	"ecommerce/go/chatbot/models"
)

type Repository interface {
	ListProducts(ctx context.Context) ([]*models.Products, error)
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func Close() error {
	return implementation.Close()
}

func ListProducts(ctx context.Context) ([]*models.Products, error) {
	return implementation.ListProducts(ctx)
}
