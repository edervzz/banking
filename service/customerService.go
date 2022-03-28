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

type CustomerService interface {
	CreateCustomer(*CustomerCreateRequest) (*CustomerCreateResponse, *utils.AppMess)
}
