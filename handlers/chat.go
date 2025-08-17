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

func NewChatHandler(db *sql.DB, aiClient *ai.OllamaClient, modelName string) *ChatHandler {
	return &ChatHandler{
		db:        db,
		aiClient:  aiClient,
		modelName: modelName,
	}
}

func (h *ChatHandler) HandleChatQuery(w http.ResponseWriter, r *http.Request) {
	var req ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Message) == "" {
		http.Error(w, "Message cannot be empty", http.StatusBadRequest)
		return
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
		return
	}

	response := ChatResponse{
		Response:  aiResp.Response,
		SessionID: req.SessionID,
		Timestamp: time.Now().Format(time.RFC3339),
		Model:     h.modelName,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		return
	}

	log.Printf("Chat [%s]: User: %s | Bot: %s", req.SessionID, req.Message, aiResp.Response)
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

func (h *ChatHandler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.aiClient.HealthCheck(ctx); err != nil {
		response := map[string]string{
			"status":     "error",
			"ai_service": err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := h.db.Ping(); err != nil {
		response := map[string]string{
			"status":   "error",
			"database": err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]string{
		"status":   "healthy",
		"ai_model": h.modelName,
		"database": "connected",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
