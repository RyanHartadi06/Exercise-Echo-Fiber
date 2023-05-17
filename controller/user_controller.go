package controller

import (
	"Go-Echo/config"
	"Go-Echo/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUserController(e echo.Context) error {
	var users []model.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func RegisterController(c echo.Context) error {
	//email := e.FormValue("email")
	//name := e.FormValue("name")

	user := model.User{}

	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}
