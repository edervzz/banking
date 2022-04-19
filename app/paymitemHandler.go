package app

import (
	"banking/service"
	"encoding/json"
	"net/http"
)

type PaymitemHandler struct {
	service service.PaymitemService
}

func (h PaymitemHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var req service.PaymitemCreateRequest
	json.NewDecoder(r.Body).Decode(&req)
	response, appErr := h.service.Create(&req)
	if appErr != nil {
		w.WriteHeader(appErr.Code)
		json.NewEncoder(w).Encode(appErr.Message)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func NewPaymitemHandler(service service.PaymitemService) *PaymitemHandler {
	return &PaymitemHandler{
		service,
	}
}
