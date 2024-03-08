package handler

import (
	"Test-Golang/features/pesanan"
	"Test-Golang/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type orderHandler struct {
	orderService pesanan.OrderServiceInterface
}

func NewOrder(orderService pesanan.OrderServiceInterface) *orderHandler {
	return &orderHandler{
		orderService: orderService,
	}
}

func (handler *orderHandler) GetAll(c echo.Context) error {
	result, errorSelect := handler.orderService.SelectAll()

	if errorSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data."+errorSelect.Error(), nil))
	}

	var orderResult []OrderResponse
	for _, value := range result {
		orderResult = append(orderResult, OrderResponse{
			ID:        value.ID,
			UserID:    value.UserID,
			ProductID: value.ProductsID,
			Tanggal:   value.Tanggal,
			Status:    value.Status,
		})
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success", orderResult))
}

func (handler *orderHandler) SearchOrderByQuery(c echo.Context) error {
	// userId := middlewares.ExtractTokenUserId(c)
	query := c.QueryParam("search")
	result, errSearch := handler.orderService.SearchOrderByQuery(query)
	if errSearch != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Error searching data. "+errSearch.Error(), nil))
	}

	orderResult := CoreResoponListOrder(result)

	return c.JSON(http.StatusOK, responses.WebResponse("Success searching data.", orderResult))
}
