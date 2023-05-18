package route

import (
	"Go-Echo/controller"
	"Go-Echo/middleware"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	middleware.LogMiddleware(e)
	e.GET("/user", controller.GetUserController)
	e.POST("/user", controller.RegisterController)

	e.GET("/product", controller.GetProductController)
	e.POST("/product", controller.StoreProductController)

	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(middleware.BasicAuthLogin))
	eAuthBasic.GET("/user", controller.GetUserController)
	return e
}
