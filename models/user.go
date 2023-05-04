package models

type User struct {
	ExternalId int    `json:"externalId" bson:"externalId"`
	Username   string `json:"username" validate:"required" bson:"username"`
	Password   string `json:"password" bson:"password" validate:"required,min=8,max=32"`
	UserType   string `json:"userType" bson:"userType" validate:"required"`
	IsActive   bool   `json:"isActive" bson:"isActive"`
}
