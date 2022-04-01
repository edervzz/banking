package service

import "banking/utils"

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisterResponse struct {
	Data string `json:"data"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Auth interface {
	Register(*RegisterRequest) (*RegisterResponse, *utils.AppMess)
	Login(*RegisterRequest) (*RegisterResponse, *utils.AppMess)
	Verify(*RegisterRequest) (*RegisterResponse, *utils.AppMess)
}
