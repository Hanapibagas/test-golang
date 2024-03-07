package handler

import (
	"Test-Golang/app/middlewares"
	"Test-Golang/features/produk"
	"Test-Golang/utils/responses"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productService produk.ProductServiceInterface
}

func NewProduct(productService produk.ProductServiceInterface) *productHandler {
	return &productHandler{
		productService: productService,
	}
}

func (handler *productHandler) Create(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	newProduct := ProductRequest{}

	errBind := c.Bind(&newProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	newProductCore := produk.ProductCore{
		Name:   newProduct.Name,
		Harga:  newProduct.Harga,
		UserID: uint(userId),
	}

	errCreate := handler.productService.Create(newProductCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error bind data. data not valid", nil))
	}

	return c.JSON(http.StatusOK, "success create product")
}

func (handler *productHandler) Update(c echo.Context) error {
	productID := c.Param("product_id")

	productId, errConv := strconv.Atoi(productID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error convert id param", nil))
	}

	newPorduct := ProductUpdateRequest{}
	errBind := c.Bind(&newPorduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	newProductRespon := RequestProductToCore(newPorduct)

	errUpdate := handler.productService.Edit(uint(productId), newProductRespon)
	if errUpdate != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	return c.JSON(http.StatusOK, "success update product")
}

func (handler *productHandler) GetAll(c echo.Context) error {
	result, errorSelect := handler.productService.SelectAll()
	log.Println("resul :", result)

	if errorSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data."+errorSelect.Error(), nil))
	}

	var productResult []ProductResponse
	for _, value := range result {
		productResult = append(productResult, ProductResponse{
			ID:     value.Id,
			UserID: value.UserID,
			Name:   value.Name,
			Harga:  value.Harga,
		})
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success", productResult))
}

func (handler *productHandler) Delete(c echo.Context) error {
	productID := c.Param("product_id")

	productId, errConv := strconv.Atoi(productID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error convert id param", nil))
	}

	err := handler.productService.Delete(uint(productId))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data."+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}
