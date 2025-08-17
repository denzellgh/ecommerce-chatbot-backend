package models

type ApiResponse struct {
	Data       any    `json:"data"`
	StatusCode int    `json:"statusCode,omitempty"`
	Message    string `json:"message,omitempty"`
	Success    bool   `json:"success,omitempty"`
}

type ApiResponseGeneric[T any] struct {
	Data       *T     `json:"data"`
	StatusCode int    `json:"statusCode,omitempty"`
	Message    string `json:"message,omitempty"`
	Success    bool   `json:"success,omitempty"`
}
