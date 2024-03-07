package data

import (
	"Test-Golang/features/pesanan/data"
	"Test-Golang/features/produk"
	"errors"

	"gorm.io/gorm"
)

type productQuery struct {
	db *gorm.DB
}

func NewProduk(db *gorm.DB) produk.ProductDataInterface {
	return &productQuery{
		db: db,
	}
}

func (repo *productQuery) Create(input produk.ProductCore) error {
	newProduct := CoreCreateToModel(input)

	tx := repo.db.Create(&newProduct)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *productQuery) Delete(id uint) error {
	tx := repo.db.Delete(&Product{}, id)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("delete failed, row affected = 0")
	}

	return nil
}

func (repo *productQuery) Edit(id uint, input produk.ProductCore) error {
	product := CoreUpdateToModel(input)

	tx := repo.db.Model(&data.Product{}).Where("id = ?", id).Updates(product)

	if tx.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

func (repo *productQuery) SelectAll() ([]produk.ProductCore, error) {
	var productData []Product

	tx := repo.db.Find(&productData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var productDataCore []produk.ProductCore
	for _, value := range productData {
		var productCore = produk.ProductCore{
			Id:     value.ID,
			Name:   value.Name,
			Harga:  value.Harga,
			UserID: value.UserID,
		}
		productDataCore = append(productDataCore, productCore)
	}

	return productDataCore, nil
}
