package main

import (
	"fmt"
	"rep/interfaces"
	"rep/repositories"
)

func main() {
	var repository interfaces.UserRepository = repositories.NewInMemoryUserRepo()

	fmt.Println("repository created", repository)

	user := &interfaces.User{
		ID:    1,
		Name:  "Maksim",
		Email: "maksim@example.com",
	}

	if err := repository.Create(user); err != nil {
		fmt.Println("Create error:", err)
	}

	u, err := repository.GetByID(1)
	if err != nil {
		fmt.Println("Get error:", err)
	} else {
		fmt.Printf("Found user: %+v\n", u)
	}
}
