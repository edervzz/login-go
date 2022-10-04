package service

import (
	"login-go/infrastructure"
	"reflect"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {

	userService := NewUserService(infrastructure.NewUserRepositoryStub())

	type args struct {
		request *CreateUserRequest
	}
	tests := []struct {
		name         string
		user         *UserService
		args         args
		response     *CreateUserResponse
		errorMessage *ErrorMessage
	}{
		{
			"OK-Create user",
			userService,
			args{
				request: &CreateUserRequest{
					Username:    "ederv",
					Email:       "qwerty@gmail.com",
					PhoneNumber: "5511223344",
					Password:    "Eder123@",
				},
			},
			&CreateUserResponse{},
			nil,
		},
		{
			"ERR-Duplicated user",
			userService,
			args{
				request: &CreateUserRequest{
					Username:    "ederv",
					Email:       "qwerty@gmail.com",
					PhoneNumber: "5511223344",
					Password:    "Eder123@",
				},
			},
			&CreateUserResponse{},
			&ErrorMessage{
				Code:    "CreateUserRequestBizValidator",
				Message: "user already exist",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, errorMessage := tt.user.CreateUser(tt.args.request)
			if !reflect.DeepEqual(response, tt.response) {
				t.Errorf("UserService.CreateUser() response = %v, want %v", response, tt.response)
			}
			if !reflect.DeepEqual(errorMessage, tt.errorMessage) {
				t.Errorf("UserService.CreateUser() errorMessage = %v, want %v", errorMessage, tt.errorMessage)
			}
		})
	}
}

func TestUserService_LoginUser(t *testing.T) {
	userService := NewUserService(infrastructure.NewUserRepositoryStub())
	userService.CreateUser(&CreateUserRequest{Username: "ederv", Email: "qwerty@gmail.com", PhoneNumber: "5511223344", Password: "Eder123@"})

	type args struct {
		request *LoginUserRequest
	}
	tests := []struct {
		name  string
		user  *UserService
		args  args
		want  *LoginUserResponse
		want1 *ErrorMessage
	}{
		{
			"OK-Login",
			userService,
			args{
				request: &LoginUserRequest{
					Username: "ederv",
					Password: "Eder123@",
				},
			},
			&LoginUserResponse{
				Username: "ederv",
				Token:    "--------",
			},
			nil,
		},
		{
			"ERR-Login",
			userService,
			args{
				request: &LoginUserRequest{
					Username: "eder",
					Password: "Eder123@",
				},
			},
			nil,
			&ErrorMessage{
				"UserNotFound",
				"user and password do not exist",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := tt.user.LoginUser(tt.args.request)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UserService.LoginUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
