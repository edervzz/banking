package service

import (
	"banking/domain"
	"banking/logger"
	"banking/utils"
	"crypto/sha1"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserServiceInterface struct {
	repo domain.UserRepository
}

func (s AuthServiceInterface) Register(req *RegisterRequest) (*RegisterResponse, *utils.AppMess) {
	if req.Username == "" || req.Password == "" {
		return &RegisterResponse{
				Data: "must send username and password",
			},
			nil
	}

	err := checkPassword(req.Password)
	if err != nil {
		return &RegisterResponse{
				Data: `At least one upper case. At least one lower case. At least one digit. At least one special character, (?=.*?[#?!@$%^&*-]). Minimum eight in length 8`,
			},
			nil
	}
	h := sha1.New()
	h.Write([]byte(req.Password))
	hashedPwd := h.Sum(nil) // TODO: change nil with key secret

	newUser := domain.User{
		Username:       req.Username,
		HashedPassword: fmt.Sprintf("%x", hashedPwd),
		Email:          req.Email,
		Role:           req.Role,
	}

	err = s.repo.Create(&newUser)
	if err != nil {
		appMess := utils.AppMess{
			Code:    400,
			Message: "cannot create user",
		}
		return nil, &appMess
	}

	return nil, nil
}
func (s AuthServiceInterface) Login(req *LoginRequest) (*LoginResponse, *utils.AppMess) {
	h := sha1.New()
	h.Write([]byte(req.Password))
	hashedPwd := h.Sum(nil)
	user := domain.User{
		Username:       req.Username,
		HashedPassword: fmt.Sprintf("%x", hashedPwd),
	}

	err := s.repo.Find(&user)
	if err != nil {
		appMess := utils.AppMess{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		logger.Info(appMess.Message)
		return nil, &appMess
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
		"exp":      time.Now().AddDate(0, 3, 0),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("SECRET"))

	if err != nil {
		appMess := utils.AppMess{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		logger.Info(appMess.Message)
		return nil, &appMess
	}

	response := LoginResponse{
		Token: signedToken,
	}
	return &response, nil
}

// constructor
func NewUserServiceInterface(repo domain.UserRepository) *AuthServiceInterface {
	return &AuthServiceInterface{
		repo,
	}
}
