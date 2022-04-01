package service

import (
	"banking/domain"
	"banking/utils"
	"fmt"
	"hash/fnv"
	"regexp"
)

type AuthServiceInterface struct {
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
				Data: `
			At least one upper case 
			At least one lower case 
			At least one digit
			At least one special character, (?=.*?[#?!@$%^&*-])
			Minimum eight in length 8`,
			},
			nil
	}
	h := fnv.New32a()
	h.Write([]byte(req.Password))
	hashedPwd := h.Sum32()

	newUser := domain.User{
		Username:       req.Username,
		HashedPassword: string(rune(hashedPwd)),
		Email:          req.Email,
	}

	fmt.Println(newUser)

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
func (s AuthServiceInterface) Login(*RegisterRequest) (*RegisterResponse, *utils.AppMess)
func (s AuthServiceInterface) Verify(*RegisterRequest) (*RegisterResponse, *utils.AppMess)

// constructor
func NewAuthServiceInterface(repo domain.UserRepository) *AuthServiceInterface {
	return &AuthServiceInterface{
		repo,
	}
}

// aux methods
func checkPassword(p string) error {
	_, err := regexp.MatchString(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`, p)
	if err != nil {
		return err
	}

	return nil
}
