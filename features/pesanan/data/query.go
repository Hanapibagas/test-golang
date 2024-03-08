package data

import (
	"Test-Golang/features/pesanan"

	"gorm.io/gorm"
)

type orderQuery struct {
	db *gorm.DB
}

func (repo *orderQuery) SearchOrderByQuery(query string) ([]pesanan.OrderCore, error) {
	var pesananData []pesanan.OrderCore

	tx := repo.db.Table("orders").Where("status LIKE ?", "%"+query+"%").Find(&pesananData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var pesananDataCore []pesanan.OrderCore
	for _, value := range pesananData {
		var pesananCore = pesanan.OrderCore{
			ID:         value.ID,
			UserID:     value.UserID,
			ProductsID: value.ProductsID,
			Tanggal:    value.Tanggal,
			Status:     value.Status,
		}
		pesananDataCore = append(pesananDataCore, pesananCore)
	}
	return pesananDataCore, nil
}

func (repo *orderQuery) SelectAll() ([]pesanan.OrderCore, error) {
	var orderData []Order

	tx := repo.db.Find(&orderData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var ordereDataCore []pesanan.OrderCore
	for _, value := range orderData {
		var OrderCore = pesanan.OrderCore{
			ID:         value.ID,
			UserID:     value.UserID,
			ProductsID: value.ProductsID,
			Tanggal:    value.Tanggal,
			Status:     value.Status,
		}
		ordereDataCore = append(ordereDataCore, OrderCore)
	}

	return ordereDataCore, nil
}

func NewOrder(db *gorm.DB) pesanan.OrderDataInterface {
	return &orderQuery{
		db: db,
	}
}
