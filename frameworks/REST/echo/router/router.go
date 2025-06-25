package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"example/service"
)

// --- Хендлер --- //
type Handler struct {
	svc service.UserService
}

func NewHandler(svc service.UserService) *Handler {
	return &Handler{svc}
}

func (h *Handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.svc.GetUser(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) CreateUser(c echo.Context) error {
	var user service.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid input")
	}
	newUser, err := h.svc.CreateUser(&user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, newUser)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user service.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid input")
	}
	updated, err := h.svc.UpdateUser(id, &user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, updated)
}
