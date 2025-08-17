package models

import "time"

type Products struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	Brand          string    `json:"brand"`
	Category_Id    string    `json:"category_id"`
	Price          float64   `json:"price"`
	Stock_Quantity int       `json:"stock_quantity"`
	Description    string    `json:"description"`
	Specs          string    `json:"specs"` // Warning
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}
