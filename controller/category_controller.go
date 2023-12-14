package controller

import (
	"Go-Echo/config"
	"Go-Echo/constants"
	"Go-Echo/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCategoryController(e echo.Context) error {
	var category []model.Category

	err := config.DB.Find(&category).Error

	if err != nil {
		panic(err.Error())
	}
	response := constants.Response{
		Message: "Success",
		Data:    category,
	}

	return e.JSON(http.StatusOK, response)
}

func PostCategoryController(e echo.Context) error {
	category := model.Category{
		Name: e.FormValue("name"),
	}

	e.Bind(&category)
	errSave := config.DB.Save(&category).Error
	if errSave != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": errSave.Error(),
		})
	}

	response := constants.Response{
		Message: "Success",
		Data:    category,
	}


	return e.JSON(http.StatusOK, response)
}

func ShowCategoryController(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	var category model.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
	}

	response := constants.Response{
		Message: "Success",
		Data:    category,
	}


	return e.JSON(http.StatusOK, response)
}

func DeleteCategoryController(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Delete(&model.Category{Id: id}).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  false,
		})
	}
	
	response := constants.Response{
		Message: "Success",
	}

	return e.JSON(http.StatusOK, response)
}
