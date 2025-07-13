package models

//easyjson:json
//go:generate easyjson -all user.go
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    int    `json:"id"`
}
