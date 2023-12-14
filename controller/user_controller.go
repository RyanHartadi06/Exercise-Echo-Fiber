package controller

import (
	"Go-Echo/config"
	"Go-Echo/constants"
	"Go-Echo/middleware"
	"Go-Echo/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUserController(e echo.Context) error {
	var users []model.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, model.Response{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := constants.Response{
		Message: "Success",
		Data:    users,
	}


	return e.JSON(http.StatusOK, response)
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
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: err.Error(),
			Data:    nil,
		})
	}
	response := constants.Response{
		Message: "Success",
		Data:    user,
	}
	return c.JSON(http.StatusOK, response)
}

func LoginUserController(c echo.Context) error {
	user := model.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	c.Bind(&user)

	err := config.DB.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Login",
			"error":   err.Error(),
		})
	}

	errCheckPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.FormValue("password")))
	if errCheckPassword != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Login Password Not Match",
			"error":   errCheckPassword.Error(),
		})
	}
	token, err := middleware.CreateToken(user.Id, user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Login",
			"error":   err.Error(),
		})
	}

	userResponse := model.UsersResponse{
		ID:      user.Id,
		Age:     user.Age,
		Email:   user.Email,
		Name:    user.Name,
		Address: user.Address,
		Token:   token,
	}

	
	response := constants.Response{
		Message: "success login",
		Data:    userResponse,
	}


	return c.JSON(http.StatusOK, response)

}

func GetSession(c echo.Context) error {
	claims := c.Get("claims").(jwt.MapClaims)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": claims,
	})
}
