package produk

type ProductCore struct {
	Id     uint
	Name   string `valid:"required"`
	Harga  string `valid:"numeric"`
	UserID uint
}

type ProductDataInterface interface {
	Create(input ProductCore) error
	SelectAll() ([]ProductCore, error)
	Edit(id uint, input ProductCore) error
	Delete(id uint) error
}

type ProductServiceInterface interface {
	Create(input ProductCore) error
	SelectAll() ([]ProductCore, error)
	Edit(id uint, input ProductCore) error
	Delete(id uint) error
}
