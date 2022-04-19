package service

import (
	"banking/utils"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
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

type User interface {
	Register(*RegisterRequest) (*RegisterResponse, *utils.AppMess)
	Login(*LoginRequest) (*LoginResponse, *utils.AppMess)
}
