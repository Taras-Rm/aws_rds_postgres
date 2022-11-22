package models

type Product struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Price uint64 `json:"price"`
}
