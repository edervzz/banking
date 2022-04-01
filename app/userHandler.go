package app

import (
	"banking/logger"
	"banking/service"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	service service.Auth
}

func (h UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	request := service.RegisterRequest{}
	json.NewDecoder(r.Body).Decode(&request)
	response, errApp := h.service.Register(&request)
	if errApp != nil {
		w.WriteHeader(errApp.Code)
		logger.Warn(errApp.Message)
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func NewUserHandler(s service.Auth) UserHandler {
	return UserHandler{
		service: s,
	}
}
