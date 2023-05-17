package route

import (
	"Go-Echo/controller"
	"Go-Echo/middleware"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()
	middleware.LogMiddleware(e)
	e.GET("/user", controller.GetUserController)
	e.POST("/user", controller.RegisterController)
	return e
}
