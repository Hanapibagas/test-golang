package handler

import "Test-Golang/features/user"

type UserRequestRegister struct {
	Name     string
	Email    string
	Password string
	Role     string
}

type UserRequestLogin struct {
	Email    string
	Password string
}

func RequestUserRegisterToCore(input UserRequestRegister) user.UserCore {
	role := "user"
	if input.Role != "" {
		role = input.Role
	}
	return user.UserCore{
		ID:       0,
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     role,
	}
}
