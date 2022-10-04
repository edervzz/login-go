package service

import (
	"fmt"
	"login-go/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserService struct {
	repository domain.IUserRepository
}

func (user *UserService) CreateUser(request *CreateUserRequest) (*CreateUserResponse, *ErrorMessage) {
	// 1. Request validation
	err := CreateUserRequestValidator(*request)
	if err != nil {
		return &CreateUserResponse{}, &ErrorMessage{Code: "CreateUserRequestValidator", Message: err.Error()}
	}
	fmt.Println("Request validated")
	// 2. Biz Validations
	userFound, _ := user.repository.Get(request.Username)
	if userFound != nil {
		return &CreateUserResponse{}, &ErrorMessage{Code: "CreateUserRequestBizValidator", Message: "user already exist"}
	}
	fmt.Println("Business rules validated")
	// 3. Save entity in repository
	user.repository.Create(&domain.User{
		Username:    request.Username,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
	})
	fmt.Println("Entity saved")

	return &CreateUserResponse{}, nil
}

func (user *UserService) LoginUser(request *LoginUserRequest) (*LoginUserResponse, *ErrorMessage) {
	// 1. Request validation
	err := LoginUserRequestValidator(*request)
	if err != nil {
		return &LoginUserResponse{}, &ErrorMessage{Code: "LoginUserRequestValidator", Message: err.Error()}
	}
	fmt.Println("Request validated")
	// 2. Find user
	userFound, _ := user.repository.Get(request.Username)
	if userFound == nil || userFound.Password != request.Password {
		return &LoginUserResponse{}, &ErrorMessage{Code: "UserNotFound", Message: "user and password do not exist"}
	}
	fmt.Println("Business rules validated")
	// 3. Create token
	claims := jwt.MapClaims{
		"username": userFound.Username,
		"email":    userFound.Email,
		"exp":      time.Now().Add(time.Second * 3600),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("SECRET"))
	fmt.Println("Token created")

	return &LoginUserResponse{
		Username: userFound.Username,
		Token:    signedToken,
	}, nil
}

func NewUserService(r domain.IUserRepository) *UserService {
	return &UserService{
		repository: r,
	}
}
