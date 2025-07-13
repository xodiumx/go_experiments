package main

import (
	"ej/models"
	"fmt"
)

func main() {

	// Marshal user
	user := models.User{ID: 1, Name: "Alice", Email: "alice@example.com"}

	data, err := user.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// Unmarshal user
	jsonInput := []byte(`{"id":2,"name":"Bob","email":"bob@example.com"}`)
	var user2 models.User
	err = user2.UnmarshalJSON(jsonInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(user2)

	// Marshal product
	product := models.Product{ID: 1, Title: "Name", Price: 100}

	productData, err := product.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(productData))
}
