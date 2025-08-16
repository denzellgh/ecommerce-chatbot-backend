package handlers

import (
	repository "ecommerce/go/chatbot/Repository"
	"ecommerce/go/chatbot/models"
	"encoding/json"
	"log"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) *models.ApiResponse

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	serverResponse := h(w, r)

	w.WriteHeader(serverResponse.StatusCode)
	json.NewEncoder(w).Encode(serverResponse)
}

func ListProductsHandler(w http.ResponseWriter, r *http.Request) *models.ApiResponse {
	posts, err := repository.ListProducts(r.Context())

	if err != nil {
		log.Println(err)
		return &models.ApiResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    "error al decodificar peticion",
			Error:      err.Error(),
		}
	}

	return &models.ApiResponse{
		Data:       posts,
		Message:    "SUCCESS",
		StatusCode: http.StatusOK,
		Error:      "",
	}

}
