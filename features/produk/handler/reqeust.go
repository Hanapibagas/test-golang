package handler

import "Test-Golang/features/produk"

type ProductRequest struct {
	Name   string `json:"name"`
	Harga  string `json:"harga"`
	UserID uint   `json:"user_id"`
}

type ProductUpdateRequest struct {
	Name  string `json:"name"`
	Harga string `json:"harga"`
}

func RequestProductToCore(input ProductUpdateRequest) produk.ProductCore {
	return produk.ProductCore{
		Name:  input.Name,
		Harga: input.Harga,
	}
}
