package data

import (
	"Test-Golang/app/middlewares"
	"Test-Golang/features/user"
	"errors"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

func (repo *userQuery) Login(email string, password string) (data *user.UserCore, err error) {
	var user User
	tx := repo.db.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := user.ModelToCoreLogin()

	return &result, nil
}

func (repo *userQuery) Register(input user.UserCore) (data *user.UserCore, token string, err error) {
	inputRegisterGorm := User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     "user",
	}

	tx := repo.db.Create(&inputRegisterGorm)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, "", errors.New("insert failed, row affected = 0")
	}

	var authGorm User
	tx = repo.db.Where("email = ?", input.Email).First(&authGorm)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	result := authGorm.ModelToCore()

	generatedToken, err := middlewares.CreateToken(int(result.ID))
	if err != nil {
		return nil, "", err
	}

	return &result, generatedToken, nil
}
