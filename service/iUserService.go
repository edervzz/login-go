package service

type CreateUserRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type CreateUserResponse struct{}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type ErrorMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type IUserService interface {
	CreateUser(*CreateUserRequest) (*CreateUserResponse, *ErrorMessage)
	LoginUser(*LoginUserRequest) (*LoginUserResponse, *ErrorMessage)
}
