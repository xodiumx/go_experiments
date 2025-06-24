package main

import (
	"fmt"
	"rep/memory"
	"rep/repo"
)

func main() {
	var repository repo.UserRepository = memory.NewInMemoryUserRepo()

	fmt.Println("repository created", repository)

	user := &repo.User{
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
