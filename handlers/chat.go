package handlers

import (
	"context"
	"database/sql"
	"encoding/json"

	"log"
	"net/http"
	"strings"
	"time"

	"ecommerce/go/chatbot/ai"
	"ecommerce/go/chatbot/database"
	"ecommerce/go/chatbot/helpers"
	"ecommerce/go/chatbot/models"
)

type ChatHandler struct {
	db        *sql.DB
	aiClient  *ai.OllamaClient
	modelName string
}

type ChatRequest struct {
	Message   string `json:"message"`
	SessionID string `json:"session_id"`
}

type ChatResponse struct {
	Response  string `json:"response"`
	SessionID string `json:"session_id"`
	Timestamp string `json:"timestamp"`
	Model     string `json:"model"`
}

type HealthCheckResponse struct {
	Status    string `json:"status"`
	AiModel   string `json:"ai_model"`
	Database  string `json:"database"`
	AiService string `json:"ai_service"`
}

func AdaptHandler(h func(http.ResponseWriter, *http.Request) *models.ApiResponse) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		resp := h(w, r)

		w.WriteHeader(resp.StatusCode)
		json.NewEncoder(w).Encode(resp)
	}
}

func NewChatHandler(db *sql.DB, aiClient *ai.OllamaClient, modelName string) *ChatHandler {
	return &ChatHandler{
		db:        db,
		aiClient:  aiClient,
		modelName: modelName,
	}
}

func (h *ChatHandler) HandleChatQuery(w http.ResponseWriter, r *http.Request) *models.ApiResponse {
	var req ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return &models.ApiResponse{
			Data:       nil,
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    "Invalid JSON",
		}
	}

	if strings.TrimSpace(req.Message) == "" {
		return &models.ApiResponse{
			Data:       nil,
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    "Message cannot be empty",
		}

	}

	if req.SessionID == "" {
		req.SessionID = helpers.GenerateSessionID()
	}

	products, err := h.getRelevantProducts(req.Message)
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		// Continue with empty products - AI can still respond
	}

	aiReq := ai.ChatRequest{
		Message:   req.Message,
		SessionID: req.SessionID,
		Products:  products,
	}

	ctx, cancel := context.WithTimeout(r.Context(), 180*time.Second)
	defer cancel()

	aiResp, err := h.aiClient.GenerateResponse(ctx, aiReq)
	if err != nil {
		log.Printf("AI generation error: %v", err)
		http.Error(w, "Sorry, I'm having trouble processing your request right now", http.StatusInternalServerError)
		return &models.ApiResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    "Sorry, I'm having trouble processing your request right now",
		}
	}

	chatResponse := &ChatResponse{
		Response:  aiResp.Response,
		SessionID: req.SessionID,
		Timestamp: time.Now().Format(time.RFC3339),
		Model:     h.modelName,
	}

	response := models.ApiResponse{
		Data:       chatResponse,
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "Success",
	}

	log.Printf("Chat [%s]: User: %s | Bot: %s", req.SessionID, req.Message, aiResp.Response)

	return &response
}

func (h *ChatHandler) getRelevantProducts(query string) ([]ai.ProductContext, error) {
	var args []interface{}

	sqlQuery := database.GetProductsQuery(query)

	rows, err := h.db.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []ai.ProductContext
	for rows.Next() {
		var p ai.ProductContext
		var specs sql.NullString

		err := rows.Scan(&p.ID, &p.Name, &p.Brand, &p.Category,
			&p.Price, &p.Stock, &p.Description, &specs)
		if err != nil {
			log.Printf("Error scanning product: %v", err)
			continue
		}

		if specs.Valid {
			p.Specs = specs.String
		}

		products = append(products, p)
	}

	return products, nil
}

func (h *ChatHandler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) *models.ApiResponse {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.aiClient.HealthCheck(ctx); err != nil {
		response := &HealthCheckResponse{
			Status:    "error",
			AiService: err.Error(),
			AiModel:   h.modelName,
			Database:  "No check",
		}
		return &models.ApiResponse{
			Data:       response,
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Error AI service",
			Success:    false,
		}
	}

	if err := h.db.Ping(); err != nil {
		response := &HealthCheckResponse{
			Status:    "error",
			AiService: "Success",
			AiModel:   h.modelName,
			Database:  err.Error(),
		}
		return &models.ApiResponse{
			Data:       response,
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Error db connection",
			Success:    false,
		}
	}

	response := &HealthCheckResponse{
		Status:    "healthy",
		AiService: "Success",
		AiModel:   h.modelName,
		Database:  "connected",
	}
	return &models.ApiResponse{
		Data:       response,
		StatusCode: http.StatusOK,
		Message:    "Success",
		Success:    true,
	}
}

func (ch *ChatHandler) HandleGetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := `SELECT id, name, description FROM categories WHERE deleted_at IS NULL ORDER BY name`
	rows, err := ch.db.Query(query)
	if err != nil {
		http.Error(w, `{"error": "Failed to get categories"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []map[string]string
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			continue
		}
		categories = append(categories, map[string]string{
			"id":          id,
			"name":        name,
			"description": description,
		})
	}

	json.NewEncoder(w).Encode(categories)
}

func (ch *ChatHandler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := `
        SELECT id, name, brand, category_id, price, stock_quantity, description, COALESCE(specs::text, '{}')
        FROM products 
        WHERE deleted_at IS NULL AND stock_quantity > 0
        ORDER BY name
    `
	rows, err := ch.db.Query(query)
	if err != nil {
		http.Error(w, `{"error": "Failed to get products"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []models.Products
	for rows.Next() {
		var p models.Products
		err := rows.Scan(&p.ID, &p.Name, &p.Brand, &p.CategoryID, &p.Price, &p.StockQuantity, &p.Description, &p.Specs)
		if err != nil {
			continue
		}
		products = append(products, p)
	}

	json.NewEncoder(w).Encode(products)
}
