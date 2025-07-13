package models

// Product if need skip - easyjson:skip
//
//go:generate easyjson -all product.go
type Product struct {
	ID    int    `json:"id"`
	Price int    `json:"price"`
	Title string `json:"title"`
}
