package models

type User struct {
	Id       int    `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	UserType string `json:"userType"`
	IsActive bool   `json:"isActive"`
}
