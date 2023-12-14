package model

type Order struct {
	Id         int    `json:"id" form:"id"`
	UserId     int    `json:"user_id" form:"user_id"`
	ProductId  int    `json:"product_id" form:"product_id"`
	TotalPrice int    `json:"total_price" form:"total_price"`
	Quantity   int    `json:"quantity" form:"quantity"`
	Status     string `json:"status" form:"status"`

	Product Product `json:"product" gorm:"foreignKey:ProductId"`
	User    User    `json:"user" gorm:"foreignKey:UserId"`
}