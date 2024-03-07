package handler

import "Test-Golang/features/user"

type UserResponRegister struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

type UsereResponById struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func CoreResponUserById(data user.UserCore) UsereResponById {
	return UsereResponById{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
		Role:  data.Role,
	}
}
