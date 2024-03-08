package routes

import (
	"Test-Golang/app/middlewares"
	_pesananData "Test-Golang/features/pesanan/data"
	_pesananHandler "Test-Golang/features/pesanan/handler"
	_pesannanService "Test-Golang/features/pesanan/service"
	_productData "Test-Golang/features/produk/data"
	_producthHandler "Test-Golang/features/produk/handler"
	_productService "Test-Golang/features/produk/service"
	_authData "Test-Golang/features/user/data"
	_authHandler "Test-Golang/features/user/handler"
	_authService "Test-Golang/features/user/service"
	"Test-Golang/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hashService := encrypts.NewHashService()

	// user
	authData := _authData.NewUser(db)
	autService := _authService.NewUser(authData, hashService)
	authHandler := _authHandler.NewUser(autService)

	// produk
	productData := _productData.NewProduk(db)
	productService := _productService.NewProduk(productData)
	productHandler := _producthHandler.NewProduct(productService)

	// pesnan
	pesananData := _pesananData.NewOrder(db)
	pesananService := _pesannanService.NewOrder(pesananData)
	pesnanHandler := _pesananHandler.NewOrder(pesananService)

	// login
	e.POST("/register", authHandler.RegisterUser)
	e.POST("/login", authHandler.LoginUser)
	e.GET("/user", authHandler.GetById, middlewares.JWTMiddleware())
	e.PUT("/update-password", authHandler.UpdatePassword, middlewares.JWTMiddleware())

	// product
	e.POST("/product", productHandler.Create, middlewares.JWTMiddleware())
	e.PUT("/product/:product_id", productHandler.Update, middlewares.JWTMiddleware())
	e.GET("/product", productHandler.GetAll, middlewares.JWTMiddleware())
	e.DELETE("/product/:product_id", productHandler.Delete, middlewares.JWTMiddleware())

	e.GET("/order", pesnanHandler.GetAll, middlewares.JWTMiddleware())
	e.GET("/order", pesnanHandler.SearchOrderByQuery, middlewares.JWTMiddleware())
}
