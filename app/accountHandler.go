package app

import (
	"banking/logger"
	"banking/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var req service.CreateAccountRequest
	json.NewDecoder(r.Body).Decode(&req)
	res, appErr := h.service.CreateAccount(&req)
	if appErr != nil {
		w.WriteHeader(appErr.Code)
		logger.Info(appErr.Message)
		json.NewEncoder(w).Encode(appErr.Message)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
	return
}

func (h AccountHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var req service.GetBalanceRequest

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	req.AccountId = id
	if err != nil {
		logger.Info("request error:" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	res, appErr := h.service.GetBalance(&req)
	if appErr != nil {
		w.WriteHeader(appErr.Code)
		logger.Info(appErr.Message)
		json.NewEncoder(w).Encode(appErr.Message)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

func (h AccountHandler) Lock(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var req service.LockAccountRequest
	acct, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Info("cannot read account")
		json.NewEncoder(w).Encode("cannot read account")
		return
	}
	req.AccountId = acct
	_, appErr := h.service.LockAccount(&req)
	if appErr != nil {
		w.WriteHeader(appErr.Code)
		logger.Info(appErr.Message)
		json.NewEncoder(w).Encode(appErr.Message)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (h AccountHandler) Unlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var req service.UnlockAccountRequest
	acct, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Info("cannot read account")
		json.NewEncoder(w).Encode("cannot read account")
		return
	}
	req.AccountId = acct
	_, appErr := h.service.UnlockAccount(&req)
	if appErr != nil {
		w.WriteHeader(appErr.Code)
		logger.Info(appErr.Message)
		json.NewEncoder(w).Encode(appErr.Message)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func NewAccountHandler(s service.AccountService) *AccountHandler {
	h := AccountHandler{
		service: s,
	}
	return &h
}
