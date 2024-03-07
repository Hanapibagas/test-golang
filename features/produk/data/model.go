package data

import (
	"Test-Golang/features/produk"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name   string
	Harga  string
	UserID uint
	// Users  User
}

func CoreCreateToModel(input produk.ProductCore) Product {
	return Product{
		Name:   input.Name,
		Harga:  input.Harga,
		UserID: input.UserID,
	}
}

func CoreUpdateToModel(input produk.ProductCore) Product {
	return Product{
		Name:  input.Name,
		Harga: input.Harga,
	}
}
