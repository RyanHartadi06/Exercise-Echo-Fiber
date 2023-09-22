package route

import (
	"Go-Echo/controller"
	"Go-Echo/middleware"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	"os"
)

func New() *echo.Echo {

	e := echo.New()
	middleware.LogMiddleware(e)
	e.GET("/uploads/:imageName", func(c echo.Context) error {
		filename := c.Param("imageName")
		return c.File("uploads/" + filename)
	})
	e.POST("/login", controller.LoginUserController)
	e.POST("/register", controller.RegisterController)

	eJwt := e.Group("/jwt")
	eJwt.Use(mid.JWT([]byte(os.Getenv("SECRET_KEY"))))
	eJwt.GET("/protected", controller.GetSession, middleware.ValidateToken)
	eJwt.GET("/user", controller.GetUserController)

	eJwt.GET("/product", controller.GetProductController)
	eJwt.POST("/product", controller.StoreProductController)
	eJwt.DELETE("/product/:name", controller.DeleteProductController)

	eJwt.GET("/news", controller.GetNewsController)
	eJwt.POST("/news", controller.PostNewsController)

	eJwt.GET("/category", controller.GetCategoryController)
	eJwt.POST("/category", controller.PostCategoryController)
	eJwt.DELETE("/category/:id", controller.DeleteCategoryController)
	return e
}
