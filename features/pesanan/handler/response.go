package handler

import "Test-Golang/features/pesanan"

type OrderResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	ProductID uint   `json:"product_id"`
	Tanggal   string `json:"tanggal"`
	Status    string `json:"status"`
}

func CoreResponOrder(data pesanan.OrderCore) OrderResponse {
	return OrderResponse{
		ID:        data.ID,
		UserID:    data.UserID,
		ProductID: data.ID,
		Tanggal:   data.Tanggal,
		Status:    data.Status,
	}
}

func CoreResoponListOrder(data []pesanan.OrderCore) []OrderResponse {
	var result []OrderResponse
	for _, value := range data {
		result = append(result, CoreResponOrder(value))
	}

	return result
}
