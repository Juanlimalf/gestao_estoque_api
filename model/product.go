package model

type Product struct {
	ID          int     `json:"id_product"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
