package models

import "time"

type UserPreferences struct {
	Id                   string    `json:"id"`
	User_Session         string    `json:"user_session"`
	Preferred_Categories string    `json:"preferred_categories"` /// Warning
	Price_Range_Min      float64   `json:"price_range_min"`
	Price_Range_Max      float64   `json:"price_range_max"`
	Preferred_Brands     string    `json:"preferred_brands"` // Warning
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	DeletedAt            time.Time `json:"deleted_at"`
}
