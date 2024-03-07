package handler

type ProductResponse struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Harga  string `json:"harga"`
}
