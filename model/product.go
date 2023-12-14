package model

type Product struct {
	Id    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Stock string `json:"stock" form:"stock"`
	Image string `json:"image" form:"image"`
}
