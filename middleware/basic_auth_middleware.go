package middleware

import (
	"Go-Echo/config"
	"Go-Echo/model"
	"github.com/labstack/echo/v4"
)

func BasicAuthLogin(email, password string, c echo.Context) (bool, error) {
	var user model.User

	err := config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
