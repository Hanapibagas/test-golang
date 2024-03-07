package data

import (
	"Test-Golang/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"default:null;unique"`
	Password string
	Role     string
}

func (u User) ModelToCoreLogin() user.UserCore {
	return user.UserCore{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
	}
}

func (u User) ModelToCoreRegister() user.UserCore {
	return user.UserCore{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
	}
}
