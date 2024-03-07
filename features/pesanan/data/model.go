package data

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     uint
	ProductsID uint
	Tanggal    string
	Status     string
}

type Product struct {
	gorm.Model
	Name   string
	Harga  string
	UserID uint
	Orders []Order `gorm:"foreignKey:ProductsID"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"default:null;unique"`
	Password string
	Role     string
	Orders   []Order   `gorm:"foreignKey:UserID"`
	Products []Product `gorm:"foreignKey:UserID"`
}
