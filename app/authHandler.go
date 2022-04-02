package app

import (
	"banking/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	service service.AuthService
}

func (a AuthHandler) Authorization() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")
			fmt.Println(currentRoute.GetName())
			fmt.Println(currentRouteVars)
			fmt.Println(authHeader)
			err := a.service.Verify(currentRoute.GetName(), currentRouteVars, authHeader)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)

		})
	}
}

// constructor
func NewAuthHandler(s service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}
