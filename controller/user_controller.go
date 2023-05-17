package controller

import (
	"Go-Echo/config"
	"Go-Echo/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
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
	ageStr := c.FormValue("age") // Assuming age is passed as a string in the form data

	age, errState := strconv.ParseInt(ageStr, 10, 64)
	if errState != nil {
		panic(errState.Error())
	}
	bytes, errBytes := bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), 14)
	if errBytes != nil {
		panic(errBytes.Error())
	}
	user := model.User{
		Email:    c.FormValue("email"),
		Age:      age,
		Name:     c.FormValue("name"),
		Address:  c.FormValue("address"),
		Password: string(bytes),
	}

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
