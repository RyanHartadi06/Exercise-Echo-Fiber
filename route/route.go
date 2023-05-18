package route

import (
	"Go-Echo/controller"
	"Go-Echo/middleware"
	"fmt"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	"net/http"
)

func New() *echo.Echo {

	e := echo.New()
	middleware.LogMiddleware(e)
	e.GET("/user", controller.GetUserController)
	e.POST("/user", controller.RegisterController)

	e.GET("/product", controller.GetProductController)
	e.POST("/product", controller.StoreProductController)
	e.GET("/page2/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("Hello %s", name)

		return ctx.String(http.StatusOK, data)
	})
	e.GET("/uploads/:imageName", func(c echo.Context) error {
		filename := c.Param("imageName")
		return c.File("uploads/" + filename)
	})
	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(middleware.BasicAuthLogin))
	eAuthBasic.GET("/user", controller.GetUserController)
	return e
}
