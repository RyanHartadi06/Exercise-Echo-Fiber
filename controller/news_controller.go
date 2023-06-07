package controller

import (
	"Go-Echo/config"
	"Go-Echo/helper"
	"Go-Echo/model"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

func GetNewsController(e echo.Context) error {
	var news []model.News

	err := config.DB.Find(&news).Error

	if err != nil {
		panic(err)
	}

	return e.JSON(http.StatusOK, model.Response{
		Message: "Success",
		Data:    news,
	})
}

func PostNewsController(e echo.Context) error {
	file, err := e.FormFile("file")
	if err != nil {
		return err
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	generateRandom := helper.GenerateRandomFileName()
	// Destination file path
	dstPath := "uploads/" + generateRandom

	// Create the destination file
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the file content to the destination
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	news := model.News{
		Title:       e.FormValue("title"),
		Description: e.FormValue("description"),
		Image:       dstPath,
	}

	e.Bind(&news)

	errSave := config.DB.Save(&news).Error
	if errSave != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create news",
		"news":    news,
	})
}
