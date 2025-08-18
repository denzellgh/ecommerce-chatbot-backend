package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type OllamaClient struct {
	baseURL    string
	httpClient *http.Client
	model      string
}

func NewOllamaClient(baseURL string, modelName string) *OllamaClient {

	return &OllamaClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 800 * time.Second,
		},
		model: modelName,
	}
}

func (c *OllamaClient) GenerateResponse(ctx context.Context, req ChatRequest) (*ChatResponse, error) {
	prompt := c.buildPrompt(req)

	ollamaReq := OllamaRequest{
		Model:  c.model,
		Prompt: prompt,
		Stream: false,
		Think:  true,
	}

	jsonData, err := json.Marshal(ollamaReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/api/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ollama API error: %s", string(body))
	}

	var ollamaResp OllamaResponse
	if err := json.Unmarshal(body, &ollamaResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	chatResp := &ChatResponse{
		Response:  ollamaResp.Response,
		SessionID: req.SessionID,
		Model:     c.model,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	return chatResp, nil
}

func (c *OllamaClient) buildPrompt(req ChatRequest) string {
	systemPrompt := `You are a helpful sales assistant for Makers Tech, a technology ecommerce company.

	Your role is to help customers find products, check inventory, compare items, and provide recommendations based on the available products.

	IMPORTANT GUIDELINES:
	- Always be conversational, friendly, and professional
	- Use the provided inventory data to give accurate, specific information
	- If asked about stock, give exact numbers and mention specific models/brands
	- When products are low in stock (5 or fewer), mention this as urgency
	- Suggest alternatives if requested items are out of stock
	- Keep responses concise but informative (2-3 sentences max)
	- Always mention specific product names, brands, and prices when relevant
	- If you don't have information about a product, be honest about limitations

	`

	if len(req.Products) > 0 {
		systemPrompt += "CURRENT INVENTORY:\n"
		for _, product := range req.Products {
			systemPrompt += fmt.Sprintf("- %s %s: $%.2f (Stock: %d) - %s\n",
				product.Brand, product.Name, product.Price, product.Stock, product.Description)
		}
		systemPrompt += "\n"
	}

	if req.Context != "" {
		systemPrompt += "CONVERSATION CONTEXT:\n" + req.Context + "\n\n"
	}

	systemPrompt += "CUSTOMER QUESTION: " + req.Message + "\n\n"
	systemPrompt += "RESPONSE (be helpful and specific):"

	return systemPrompt
}

func (c *OllamaClient) HealthCheck(ctx context.Context) error {
	resp, err := http.Get(c.baseURL + "/api/tags")
	if err != nil {
		return fmt.Errorf("ollama is not running: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ollama health check failed with status: %d", resp.StatusCode)
	}

	var tags struct {
		Models []struct {
			Name string `json:"name"`
		} `json:"models"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &tags); err != nil {
		return err
	}

	for _, model := range tags.Models {
		if model.Name == c.model {
			return nil
		}
	}

	return fmt.Errorf("model %s not found, run: ollama pull %s", c.model, c.model)
}
