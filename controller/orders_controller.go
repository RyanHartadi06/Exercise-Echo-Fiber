package controller

import (
	"Go-Echo/config"
	"Go-Echo/constants"
	"Go-Echo/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)


func GetOrderController(e echo.Context) error {
	claims :=  e.Get("claims").(jwt.MapClaims)
	var order []model.Order

	err := config.DB.Preload("Product").Preload("User").Where("user_id = ?", claims["userId"]).Find(&order).Error
	

	if err != nil {
		panic(err.Error())
	}

	response := constants.Response{
		Message: "Success",
		Data:    order,
	}

	return e.JSON(http.StatusOK, response)
}

func CreateOrderController(e echo.Context) error {
	claims :=  e.Get("claims").(jwt.MapClaims)
	
	product_id, _ := strconv.Atoi(e.FormValue("product_id"))
	total_price, _ := strconv.Atoi(e.FormValue("total_price"))
	quantity, _ := strconv.Atoi(e.FormValue("quantity"))
	var product []model.Product

	checkProduct := config.DB.Find(&product).Where("id = ?", product_id).Error

	if checkProduct != nil {
		response := constants.Response{
			Message: "Product not found",
			Data:    product,
		}
		return e.JSON(http.StatusNotFound, response)
	}

	order := model.Order{
		UserId:     int(claims["userId"].(float64)),
		ProductId:  product_id,
		TotalPrice: total_price,
		Quantity:   quantity,
		Status:     "PENDING",
	}
	
	e.Bind(&order)

	err := config.DB.Create(&order).Error

	if err != nil {
		panic(err.Error())
	}

	response := constants.Response{
		Message: "Created Order Success",
	}

	return e.JSON(http.StatusOK, response)
}

func UpdateToPaidOrderController(e echo.Context) error {
	claims :=  e.Get("claims").(jwt.MapClaims)
	id, _ := strconv.Atoi(e.Param("id"))
	var order model.Order

	err := config.DB.Model(&order).Where("id = ?", id).Where("user_id = ?" , claims["userId"]).Update("status", "PAID").Error

	if err != nil {
		panic(err.Error())
	}

	response := constants.Response{
		Message: "Update Order Success",
	}

	return e.JSON(http.StatusOK, response)
}