package routes

import (
	_authData "Test-Golang/features/user/data"
	_authHandler "Test-Golang/features/user/handler"
	_authService "Test-Golang/features/user/service"
	"Test-Golang/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hashService := encrypts.NewHashService()

	authData := _authData.NewUser(db)
	autService := _authService.NewUser(authData, hashService)
	authHandler := _authHandler.NewUser(autService)

	// login
	e.POST("/register", authHandler.RegisterUser)
	e.POST("/login", authHandler.LoginUser)
}
