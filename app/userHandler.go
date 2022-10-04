package app

import (
	"encoding/json"
	"login-go/service"
	"net/http"
)

type UserHandler struct {
	service service.IUserService
}

func (h UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	// 1. Set JSON content-type
	w.Header().Add("Content-Type", "application/json")
	// 2. Decode body to request type
	request := service.CreateUserRequest{}
	json.NewDecoder(r.Body).Decode(&request)
	// 3. Call service/command
	response, err := h.service.CreateUser(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	// 4. Send response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	// 1. Set JSON content-type
	w.Header().Add("Content-Type", "application/json")
	// 2. Decode body to request type
	request := service.LoginUserRequest{}
	json.NewDecoder(r.Body).Decode(&request)
	// 3. Call service/command
	response, err := h.service.LoginUser(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	// 4. Send response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func NewUserHandler(s service.IUserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}
