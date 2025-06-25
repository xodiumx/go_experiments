package tests

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"

	"example/router"
	"example/service"
	"example/service/mocks"
)

func TestGetUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockSvc := new(mocks.UserService)
	mockSvc.On("GetUser", 1).Return(&service.User{ID: 1, Name: "Test", Email: "test@example.com"}, nil)

	h := router.NewHandler(mockSvc)

	err := h.GetUser(c)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
	require.Contains(t, rec.Body.String(), "Test")

	mockSvc.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	e := echo.New()
	body := `{"name": "Alice", "email": "alice@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockSvc := new(mocks.UserService)
	mockSvc.On("CreateUser", mock.MatchedBy(func(u *service.User) bool {
		return u.Name == "Alice" && u.Email == "alice@example.com"
	})).Return(&service.User{ID: 1, Name: "Alice", Email: "alice@example.com"}, nil)

	h := router.NewHandler(mockSvc)

	err := h.CreateUser(c)

	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, rec.Code)
	require.Contains(t, rec.Body.String(), "Alice")

	mockSvc.AssertExpectations(t)
}

func TestGetUser_NotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/98", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("99")

	mockSvc := new(mocks.UserService)
	mockSvc.On("GetUser", 99).Return(nil, echo.NewHTTPError(http.StatusNotFound, "user not found"))

	h := router.NewHandler(mockSvc)

	err := h.GetUser(c)

	require.Error(t, err)
	var httpError *echo.HTTPError
	ok := errors.As(err, &httpError)
	require.True(t, ok)
	require.Equal(t, http.StatusNotFound, httpError.Code)

	mockSvc.AssertExpectations(t)
}
