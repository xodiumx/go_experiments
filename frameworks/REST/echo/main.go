package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"example/router"
	"example/service"
)

// --- main --- //
func main() {
	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} ${remote_ip} ${method} ${uri} ${status} ${latency_human}` + "\n",
	}))

	svc := service.NewUserService()
	handler := router.NewHandler(svc)

	app.GET("/users/:id", handler.GetUser)
	app.POST("/users", handler.CreateUser)
	app.PUT("/users/:id", handler.UpdateUser)

	app.Logger.Fatal(app.Start(":8080"))
}
