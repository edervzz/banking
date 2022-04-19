package service

import "banking/utils"

type PaymitemCreateRequest struct {
	AccountId int     `json:"accountId"`
	TAmount   float64 `json:"tamount"`
	TransType string  `json:"transType"`
	Concept   string  `json:"concept"`
}

type PaymitemCreateResponse struct {
	DocumentId int `json:"documentId"`
}

type PaymitemService interface {
	Create(*PaymitemCreateRequest) (*PaymitemCreateResponse, *utils.AppMess)
}
