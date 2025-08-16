package database

import (
	"context"
	"database/sql"
	"ecommerce/go/chatbot/models"
	"log"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return &PostgresRepository{
		db: db,
	}, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}

func (repo *PostgresRepository) ListProducts(ctx context.Context) ([]*models.Products, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, brand, created_at FROM products WHERE deleted_at IS NULL")

	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()

		if err != nil {
			log.Fatal("Error closing database connection", err)
		}
	}()

	var products []*models.Products

	for rows.Next() {
		var product = models.Products{}
		if err = rows.Scan(&product.Id, &product.Name, &product.Brand, &product.CreatedAt); err == nil {
			products = append(products, &product)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil

}
