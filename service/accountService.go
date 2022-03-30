package service

import "banking/utils"

type CreateAccountRequest struct {
	CustomerId  int     `json:"customerId"`
	AccountType string  `json:"account_type"`
	Balance     float64 `json:"balance"`
}

type CreateAccountResponse struct {
	AccountId int `json:"accountId"`
}

type GetBalanceRequest struct {
	AccountId int `json:"accountId"`
}

type GetBalanceResponse struct {
	Balance float64 `json:"balance"`
}

type LockAccountRequest struct {
	AccountId int `json:"accountId"`
}

type LockAccountResponse struct{}

type UnlockAccountRequest struct {
	AccountId int `json:"accountId"`
}

type UnlockAccountResponse struct{}

type AccountService interface {
	CreateAccount(*CreateAccountRequest) (*CreateAccountResponse, *utils.AppMess)
	GetBalance(*GetBalanceRequest) (*GetBalanceResponse, *utils.AppMess)
	LockAccount(*LockAccountRequest) (*LockAccountResponse, *utils.AppMess)
	UnlockAccount(*UnlockAccountRequest) (*UnlockAccountResponse, *utils.AppMess)
}
