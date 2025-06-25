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
	h := router.NewHandler(svc)

	app.GET("/users/:id", h.GetUser)
	app.POST("/users", h.CreateUser)
	app.PUT("/users/:id", h.UpdateUser)

	app.Logger.Fatal(app.Start(":8080"))
}
