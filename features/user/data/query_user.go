package data

import (
	"Test-Golang/app/middlewares"
	"Test-Golang/features/user"
	"errors"

	"golang.org/x/crypto/bcrypt"
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
	registerUser := User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     "user",
	}

	tx := repo.db.Create(&registerUser)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, "", errors.New("insert failed, row affected = 0")
	}

	var authUser User
	tx = repo.db.Where("email = ?", input.Email).First(&authUser)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	result := authUser.ModelToCoreRegister()

	generatedToken, err := middlewares.CreateToken(int(result.ID))
	if err != nil {
		return nil, "", err
	}

	return &result, generatedToken, nil
}

func (repo *userQuery) CheckPassword(savedPassword string, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(savedPassword), []byte(inputPassword))
	return err == nil
}

func (repo *userQuery) GetById(id uint) (*user.UserCore, error) {
	var userData User
	tx := repo.db.First(&userData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	respunUser := &user.UserCore{
		ID:    userData.ID,
		Name:  userData.Name,
		Email: userData.Email,
		Role:  userData.Role,
	}

	return respunUser, nil
}

func (repo *userQuery) UpdatePassword(id uint, input user.AuthCorePassword) error {
	userInput := User{
		Password: input.Password,
	}

	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(&userInput)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("edit failed, row affected = 0")
	}

	return nil
}
