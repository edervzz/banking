package app

import (
	"banking/logger"
	"banking/service"
	"banking/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Customer
type CustomerHandler struct {
	service service.CustomerService
}

func (h *CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var req service.CustomerCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		appError := utils.NewBadRequest("cannot decode request")
		logger.Warn(err.Error())
		w.WriteHeader(appError.Code)
		json.NewEncoder(w).Encode(appError.Message)
		return
	}

	res, appError := h.service.CreateCustomer(&req)
	if appError != nil {
		logger.Warn(appError.Message)
		w.WriteHeader(appError.Code)
		json.NewEncoder(w).Encode(appError.Message)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *CustomerHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	req := service.CustomerGetRequest{
		CustomerId: id,
	}
	res, appErr := h.service.GetCustomer(&req)
	if appErr != nil {
		logger.Warn(appErr.Message)
		w.WriteHeader(appErr.Code)
		json.NewEncoder(w).Encode(appErr.Message)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func NewCustomerHandler(service service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		service,
	}
}

// Migrations
type MigrationHandler struct {
	service service.MigrationService
}

func (mh MigrationHandler) Migrations(w http.ResponseWriter, r *http.Request) {
	result := mh.service.Prepare()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func NewMigrationHandler(service service.MigrationService) MigrationHandler {
	return MigrationHandler{
		service,
	}
}
