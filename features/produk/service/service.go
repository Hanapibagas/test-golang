package service

import (
	"Test-Golang/features/produk"
	"errors"

	"github.com/go-playground/validator/v10"
)

type productService struct {
	produkData produk.ProductDataInterface
	validate   *validator.Validate
}

func NewProduk(repo produk.ProductDataInterface) produk.ProductServiceInterface {
	return &productService{
		produkData: repo,
		validate:   validator.New(),
	}
}

func (service *productService) Create(input produk.ProductCore) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	err := service.produkData.Create(input)
	return err
}

func (service *productService) Delete(id uint) error {
	err := service.produkData.Delete(id)
	return err
}

func (service *productService) Edit(id uint, input produk.ProductCore) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	err := service.produkData.Edit(id, input)
	return err
}

func (service *productService) SelectAll() ([]produk.ProductCore, error) {
	result, err := service.produkData.SelectAll()
	return result, err
}
