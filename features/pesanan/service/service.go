package service

import "Test-Golang/features/pesanan"

type orderService struct {
	orderData pesanan.OrderDataInterface
}

// SearchOrderByQuery implements pesanan.OrderServiceInterface.
func (service *orderService) SearchOrderByQuery(query string) ([]pesanan.OrderCore, error) {
	result, err := service.orderData.SearchOrderByQuery(query)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return []pesanan.OrderCore{}, nil
	}

	return result, nil
}

// SelectAll implements pesanan.OrderServiceInterface.
func (service *orderService) SelectAll() ([]pesanan.OrderCore, error) {
	result, err := service.orderData.SelectAll()
	return result, err
}

func NewOrder(repo pesanan.OrderDataInterface) pesanan.OrderServiceInterface {
	return &orderService{
		orderData: repo,
	}
}
