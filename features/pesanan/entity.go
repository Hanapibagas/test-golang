package pesanan

type OrderCore struct {
	ID         uint
	UserID     uint
	ProductsID uint
	Tanggal    string
	Status     string
}

type OrderDataInterface interface {
	SelectAll() ([]OrderCore, error)
	SearchOrderByQuery(query string) ([]OrderCore, error)
}

type OrderServiceInterface interface {
	SelectAll() ([]OrderCore, error)
	SearchOrderByQuery(query string) ([]OrderCore, error)
}
