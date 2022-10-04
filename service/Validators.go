package service

import (
	"errors"
	"unicode"
)

func PasswordValidator(password string) error {
	if len(password) < 6 || len(password) > 12 {
		return errors.New("password must have a length between 6 and 16")
	}
	isUpper := false
	isLower := false
	isNumber := false
	isSpecial := false
	for _, v := range password {
		switch {
		case unicode.IsUpper(v):
			isUpper = true
		case unicode.IsLower(v):
			isLower = true
		case unicode.IsNumber(v):
			isNumber = true
		case v == '@' || v == '$' || v == '&':
			isSpecial = true
		}
	}
	if !isUpper || !isLower || !isNumber || !isSpecial {
		return errors.New("password must must have at least 1 upper-case, 1 lower-case, 1 number and 1 special character(@ $ &)")
	}
	return nil
}

func CreateUserRequestValidator(r CreateUserRequest) error {
	// 1. Validate empty fields
	if r.Username == "" || r.Email == "" || r.Password == "" || r.PhoneNumber == "" {
		return errors.New("username, email, password and phone number must not be empty")
	}
	// 2. Password validation
	passwordValidatorError := PasswordValidator(r.Password)
	if passwordValidatorError != nil {
		return passwordValidatorError
	}
	// 3. Phone number validation
	if len(r.PhoneNumber) != 10 {
		return errors.New("phone number must have a length 10")
	}
	for _, v := range r.PhoneNumber {
		switch {
		case !unicode.IsNumber(v):
			return errors.New("phone number format error")
		}
	}

	return nil
}

func LoginUserRequestValidator(r LoginUserRequest) error {
	// 1. Validate empty fields
	if r.Username == "" || r.Password == "" {
		return errors.New("username, and password must not be empty")
	}
	// 2. Password validation
	passwordValidatorError := PasswordValidator(r.Password)
	if passwordValidatorError != nil {
		return passwordValidatorError
	}
	return nil
}
