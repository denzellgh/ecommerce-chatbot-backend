package database

import (
	"context"
	"database/sql"
	repository "ecommerce/go/chatbot/Repository"
	"ecommerce/go/chatbot/models"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func StartPostgresDB(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connected")

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connected")

	return &PostgresRepository{
		db: db,
	}, nil
}

func StartRepositoryDB(DatabaseUrl string) {

	repo, err := NewPostgresRepository(DatabaseUrl)

	if err != nil {
		log.Fatal("Error with database connection", err)
	}

	repository.SetRepository(repo)
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
