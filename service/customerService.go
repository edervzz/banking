package service

import "banking/utils"

type CustomerCreateRequest struct {
	Fullname string `json:"fullname"`
	City     string `json:"city"`
	Zipcode  string `json:"zipcode"`
}

type CustomerCreateResponse struct {
	CustomerId int `json:"id"`
}

type CustomerGetRequest struct {
	CustomerId int `json:"id"`
}

type CustomerGetResponse struct {
	CustomerId int    `json:"id"`
	Fullname   string `json:"fullname"`
	City       string `json:"city"`
	Zipcode    string `json:"zipcode"`
}

type CustomerService interface {
	CreateCustomer(*CustomerCreateRequest) (*CustomerCreateResponse, *utils.AppMess)
	GetCustomer(*CustomerGetRequest) (*CustomerGetResponse, *utils.AppMess)
}
