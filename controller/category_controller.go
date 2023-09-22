package controller

import (
	"Go-Echo/config"
	"Go-Echo/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetCategoryController(e echo.Context) error {
	var category []model.Category

	err := config.DB.Find(&category).Error

	if err != nil {
		panic(err.Error())
	}

	return e.JSON(http.StatusOK, model.Response{
		Message: "Success",
		Data:    category,
	})
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
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create category",
		"category": category,
	})
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
	return e.JSON(http.StatusInternalServerError, map[string]interface{}{
		"message": "Berhasil Hapus",
		"status":  true,
	})
}
