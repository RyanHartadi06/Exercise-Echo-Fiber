package main

import (
	"Go-Echo/config"
	"Go-Echo/route"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load()
	config.InitDB()
	e := route.New()
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Start(":3001")
}
