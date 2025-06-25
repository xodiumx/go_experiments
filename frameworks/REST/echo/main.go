package main

import (
	"example/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} ${remote_ip} ${method} ${uri} ${status} ${latency_human}` + "\n",
	}))

	// Роуты
	e.GET("/users/:id", router.GetUser)
	e.POST("/users", router.CreateUser)
	e.PUT("/users/:id", router.UpdateUser)

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8080"))
}
