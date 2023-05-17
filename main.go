package main

import (
	"Go-Echo/config"
	"Go-Echo/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":3001")
}
