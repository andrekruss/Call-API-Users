package controllers

import (
	"call-ms-users/database"
	"call-ms-users/helpers"
	"call-ms-users/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func ParseUser(ctx *fiber.Ctx) error {

	collection := database.DBConn.Database("call-users").Collection("users")
	user := new(models.User)

	err := ctx.BodyParser(user)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": "Error while parsing user information.",
		})
	}

	ctx.Locals("user", user)
	ctx.Locals("collection", collection)
	return ctx.Next()
}

func ValidateUser(ctx *fiber.Ctx) error {
	err := helpers.Validate.Struct(ctx.Locals("user").(*models.User))
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "User validation error",
		})
	}
	return ctx.Next()
}

func PasswordHash(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.User)
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error: Error while processing request.",
		})
	}
	user.Password = hashedPassword
	ctx.Locals("user", user)
	return ctx.Next()
}

func SaveUser(ctx *fiber.Ctx) error {
	collection := ctx.Locals("collection").(*mongo.Collection)
	user := ctx.Locals("user").(*models.User)
	res, err := collection.InsertOne(ctx.Context(), user)

	log.Println(res.InsertedID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error: Could not save user information on database.",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}
