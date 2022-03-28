package utils

import "net/http"

type AppMess struct {
	Code    int
	Message string
}

func NewInternalError(m string) *AppMess {
	return &AppMess{
		Code:    http.StatusInternalServerError,
		Message: m,
	}
}

func NewBadRequest(m string) *AppMess {
	return &AppMess{
		Code:    http.StatusBadRequest,
		Message: m,
	}
}

func NewCreated(m string) *AppMess {
	return &AppMess{
		Code:    http.StatusCreated,
		Message: m,
	}
}

func NewOK(m string) *AppMess {
	return &AppMess{
		Code:    http.StatusOK,
		Message: m,
	}
}
