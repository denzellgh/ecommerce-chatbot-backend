package ai

import "time"

type ChatRequest struct {
	Message   string           `json:"message"`
	Context   string           `json:"context,omitempty"`
	SessionID string           `json:"session_id,omitempty"`
	Products  []ProductContext `json:"products,omitempty"`
}

type ProductContext struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
	Specs       string  `json:"specs"`
}

type ChatResponse struct {
	Response  string `json:"response"`
	SessionID string `json:"session_id"`
	Model     string `json:"model"`
	Timestamp string `json:"timestamp"`
}

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
	Think  bool   `json:"think"`
}

type OllamaResponse struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Response  string    `json:"response"`
	Done      bool      `json:"done"`
	Context   []int     `json:"context,omitempty"`
}
