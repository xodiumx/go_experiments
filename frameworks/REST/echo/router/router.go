package router

import (
	"example/schemas"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var users = map[int]*schemas.User{}
var idCounter = 1

// GET /users/:id
func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, ok := users[id]
	if !ok {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
	}
	return c.JSON(http.StatusOK, user)
}

// POST /users
func CreateUser(c echo.Context) error {
	u := new(schemas.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}
	u.ID = idCounter
	idCounter++ // just example
	users[u.ID] = u
	return c.JSON(http.StatusCreated, u)
}

// PUT /users/:id
func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	existing, ok := users[id]
	if !ok {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
	}

	u := new(schemas.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	// Обновляем только поля Name и Email
	existing.Name = u.Name
	existing.Email = u.Email
	return c.JSON(http.StatusOK, existing)
}
