package model

type User struct {
	Id       int    `json: "id" form: "id"`
	Age      int64  `json: "age" form: "age"`
	Email    string `json: "email" form: "email"`
	Name     string `json: "name" form: "name"`
	Address  string `json: "address" form: "address"`
	Password string `json: "password" form: "password"`
}
