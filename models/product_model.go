package models

type Product struct {
	ProductId  int64  `json:"productId"`
	CategoryId string `json:"categoryId"`
	Name       string `json:"name"`
	Price      int64  `json:"price"`
}
