package service

import (
	"banking/domain"
	"errors"

	"github.com/golang-jwt/jwt"
)

type accessTokenClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Exp      string `json:"exp"`
	jwt.StandardClaims
}
type AuthServiceInterface struct {
	repo domain.UserRepository
}

func (s AuthServiceInterface) Verify(routeName string, routeVars map[string]string, token string) error {
	jwtToken, err := jwt.ParseWithClaims(
		token,
		&accessTokenClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("SECRET"), nil
		})
	if err != nil {
		return err
	}

	routesRoles := getRouteRoles()

	claims := jwtToken.Claims.(*accessTokenClaim)
	if routes, ok := routesRoles[claims.Role]; jwtToken.Valid && ok {
		for _, v := range routes {
			if v == routeName {
				return nil
			}
		}
		return errors.New("wrong route-role")
	}

	return nil
}

// constructor
func NewAuthServiceInterface(repo domain.UserRepository) *AuthServiceInterface {
	return &AuthServiceInterface{
		repo,
	}
}

// aux methods
func getRouteRoles() (routeRoles map[string][]string) {
	routeRoles = make(map[string][]string)
	routeRoles["BKK010"] = []string{
		"GetCustomer", "GetAccount",
		"CreatePaymitem",
	}
	routeRoles["ADMIN"] = []string{
		"CreateCustomer", "CreateAccount",
		"LockAccount", "UnlockAccount",
	}
	return
}

func checkPassword(p string) error {
	// _, err := regexp.MatchString(`/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[#$@!%&*?])[A-Za-z\d#$@!%&*?]{8,30}$/`, p)
	// if err != nil {
	// 	return err
	// }
	return nil
}
