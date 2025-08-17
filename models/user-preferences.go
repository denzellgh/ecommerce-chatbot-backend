package models

import "time"

type UserPreferences struct {
	ID                  string    `json:"id"`
	UserSession         string    `json:"user_session"`
	PreferredCategories []string  `json:"preferred_categories"`
	PriceRangeMin       *float64  `json:"price_range_min"`
	PriceRangeMax       *float64  `json:"price_range_max"`
	PreferredBrands     []string  `json:"preferred_brands"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           time.Time `json:"deleted_at"`
}
