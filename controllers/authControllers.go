package controllers

import (
	"call-ms-users/database"
	"call-ms-users/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type LoginInfo struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func ParseUserLoginInfo(ctx *fiber.Ctx) error {
	loginInfo := new(LoginInfo)
	err := ctx.BodyParser(loginInfo)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": "Error while parsing user login information.",
		})
	}
	ctx.Locals("loginInfo", loginInfo)
	return ctx.Next()
}

func GetUserByUsername(ctx *fiber.Ctx) error {
	var user models.User
	collection := database.DBConn.Database("call-users").Collection("users")
	loginInfo := ctx.Locals("loginInfo").(*LoginInfo)
	filter := bson.M{"username": loginInfo.Username}
	err := collection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Invalid username or password",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Username found",
	})
}
