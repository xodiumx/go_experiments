package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// --- Модель и сервис --- //
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//go:generate mockery --name=UserService --output=mocks --case=underscore
type UserService interface {
	GetUser(id int) (*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(id int, user *User) (*User, error)
}

// --- Реализация сервиса --- //
type UserServiceImpl struct {
	data   map[int]*User
	nextID int
}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{data: make(map[int]*User), nextID: 1}
}

func (s *UserServiceImpl) GetUser(id int) (*User, error) {
	user, ok := s.data[id]
	if !ok {
		return nil, echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return user, nil
}

func (s *UserServiceImpl) CreateUser(user *User) (*User, error) {
	user.ID = s.nextID
	s.data[user.ID] = user
	s.nextID++
	return user, nil
}

func (s *UserServiceImpl) UpdateUser(id int, user *User) (*User, error) {
	existing, ok := s.data[id]
	if !ok {
		return nil, echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	existing.Name = user.Name
	existing.Email = user.Email
	return existing, nil
}
