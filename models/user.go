package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	UserType string `json:"usertype" validate:"required"`
	IsActive bool   `json:"isactive"`
}
