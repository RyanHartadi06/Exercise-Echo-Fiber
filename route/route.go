package route

import (
	"Go-Echo/controller"
	"Go-Echo/middleware"
	"os"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
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
	eJwt.GET("/category/:id", controller.ShowCategoryController)
	eJwt.POST("/category", controller.PostCategoryController)
	eJwt.DELETE("/category/:id", controller.DeleteCategoryController)

	eJwt.GET("/orders", controller.GetOrderController, middleware.ValidateToken)
	eJwt.POST("/orders", controller.CreateOrderController, middleware.ValidateToken)
	eJwt.POST("/update-order/:id", controller.UpdateToPaidOrderController, middleware.ValidateToken)
	return e
}
