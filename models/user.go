package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	UserType string `json:"usertype"`
	IsActive bool   `json:"isactive"`
}
