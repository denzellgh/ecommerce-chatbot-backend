package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/lib/pq"
)

type RecommendationService struct {
	db *sql.DB
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	CategoryID  string  `json:"category_id"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock_quantity"`
	Description string  `json:"description"`
	Specs       string  `json:"specs"`
}

type UserPreference struct {
	ID                  string   `json:"id"`
	UserSession         string   `json:"user_session"`
	PreferredCategories []string `json:"preferred_categories"`
	PriceRangeMin       *float64 `json:"price_range_min"`
	PriceRangeMax       *float64 `json:"price_range_max"`
	PreferredBrands     []string `json:"preferred_brands"`
}

type RecommendationResponse struct {
	HighlyRecommended []Product `json:"highly_recommended"`
	Recommended       []Product `json:"recommended"`
	Other             []Product `json:"other"`
	Message           string    `json:"message"`
}

func NewRecommendationService(db *sql.DB) *RecommendationService {
	return &RecommendationService{db: db}
}

func (rs *RecommendationService) GetRecommendations(userSession string) (*RecommendationResponse, error) {
	prefs, err := rs.getUserPreferences(userSession)
	if err != nil {
		log.Printf("Error getting user preferences: %v", err)
		return rs.getGeneralRecommendations()
	}

	products, err := rs.getAllProducts()
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %v", err)
	}

	highly := []Product{}
	recommended := []Product{}
	other := []Product{}

	for _, product := range products {
		score := rs.calculateRecommendationScore(product, prefs)

		if score >= 80 {
			highly = append(highly, product)
		} else if score >= 50 {
			recommended = append(recommended, product)
		} else {
			other = append(other, product)
		}
	}

	return &RecommendationResponse{
		HighlyRecommended: highly,
		Recommended:       recommended,
		Other:             other,
		Message:           fmt.Sprintf("Found %d highly recommended products based on your preferences", len(highly)),
	}, nil
}

func (rs *RecommendationService) getUserPreferences(userSession string) (*UserPreference, error) {
	query := `
		SELECT id, user_session, preferred_categories, price_range_min, price_range_max, preferred_brands
		FROM user_preferences 
		WHERE user_session = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC 
		LIMIT 1
	`

	var prefs UserPreference
	var preferredCategories pq.StringArray
	var preferredBrands pq.StringArray

	err := rs.db.QueryRow(query, userSession).Scan(
		&prefs.ID,
		&prefs.UserSession,
		&preferredCategories,
		&prefs.PriceRangeMin,
		&prefs.PriceRangeMax,
		&preferredBrands,
	)

	if err != nil {
		return nil, err
	}

	prefs.PreferredCategories = []string(preferredCategories)
	prefs.PreferredBrands = []string(preferredBrands)

	return &prefs, nil
}

func (rs *RecommendationService) getAllProducts() ([]Product, error) {
	query := `
		SELECT id, name, brand, category_id, price, stock_quantity, description, COALESCE(specs::text, '{}')
		FROM products 
		WHERE deleted_at IS NULL AND stock_quantity > 0
		ORDER BY stock_quantity DESC, price ASC
	`

	rows, err := rs.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Brand, &p.CategoryID, &p.Price, &p.Stock, &p.Description, &p.Specs)
		if err != nil {
			log.Printf("Error scanning product: %v", err)
			continue
		}
		products = append(products, p)
	}

	return products, nil
}

func (rs *RecommendationService) calculateRecommendationScore(product Product, prefs *UserPreference) int {
	score := 0

	if prefs.PreferredCategories != nil {
		for _, cat := range prefs.PreferredCategories {
			if cat == product.CategoryID {
				score += 40
				break
			}
		}
	}

	if prefs.PreferredBrands != nil {
		for _, brand := range prefs.PreferredBrands {
			if strings.EqualFold(brand, product.Brand) {
				score += 30
				break
			}
		}
	}

	if prefs.PriceRangeMin != nil && prefs.PriceRangeMax != nil {
		if product.Price >= *prefs.PriceRangeMin && product.Price <= *prefs.PriceRangeMax {
			score += 20
		}
	}

	if product.Stock > 5 {
		score += 10
	} else if product.Stock > 0 {
		score += 5
	}

	return score
}

func (rs *RecommendationService) getGeneralRecommendations() (*RecommendationResponse, error) {
	products, err := rs.getAllProducts()
	if err != nil {
		return nil, err
	}

	highly := []Product{}
	recommended := []Product{}
	other := []Product{}

	for _, product := range products {
		if product.Stock > 10 && product.Price < 500 {
			highly = append(highly, product)
		} else if product.Stock > 5 {
			recommended = append(recommended, product)
		} else {
			other = append(other, product)
		}
	}

	return &RecommendationResponse{
		HighlyRecommended: highly,
		Recommended:       recommended,
		Other:             other,
		Message:           "General recommendations based on popular products",
	}, nil
}

func (rs *RecommendationService) SaveUserPreferences(userSession string, categories []string, brands []string, minPrice, maxPrice *float64) error {
	prefID := fmt.Sprintf("pref_%d", len(userSession)*12345)

	query := `
		INSERT INTO user_preferences (id, user_session, preferred_categories, preferred_brands, price_range_min, price_range_max)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (user_session) 
		DO UPDATE SET 
			preferred_categories = $3,
			preferred_brands = $4,
			price_range_min = $5,
			price_range_max = $6,
			updated_at = NOW()
	`

	_, err := rs.db.Exec(query, prefID, userSession, pq.Array(categories), pq.Array(brands), minPrice, maxPrice)
	return err
}

func (ch *ChatHandler) HandleRecommendations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userSession := r.Header.Get("X-User-Session")
	if userSession == "" {
		userSession = fmt.Sprintf("user_%d", r.Context().Value("requestTime"))
	}

	recService := NewRecommendationService(ch.db)
	recommendations, err := recService.GetRecommendations(userSession)
	if err != nil {
		log.Printf("Error getting recommendations: %v", err)
		http.Error(w, `{"error": "Failed to get recommendations"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(recommendations)
}

func (ch *ChatHandler) HandleSavePreferences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userSession := r.Header.Get("X-User-Session")
	if userSession == "" {
		userSession = fmt.Sprintf("user_%d", r.Context().Value("requestTime"))
	}

	var reqBody struct {
		Categories []string `json:"categories"`
		Brands     []string `json:"brands"`
		MinPrice   *float64 `json:"min_price"`
		MaxPrice   *float64 `json:"max_price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	recService := NewRecommendationService(ch.db)
	err := recService.SaveUserPreferences(userSession, reqBody.Categories, reqBody.Brands, reqBody.MinPrice, reqBody.MaxPrice)
	if err != nil {
		log.Printf("Error saving preferences: %v", err)
		http.Error(w, `{"error": "Failed to save preferences"}`, http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"message": "Preferences saved successfully"}`))
}
