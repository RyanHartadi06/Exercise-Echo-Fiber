package model

type News struct {
	Id          int    `json:"id" form:"id"`
	Title       string `json:"name" form:"name"`
	Description string `json:"name" form:"description"`
	Image       string `json:"image" form:"image"`
}
