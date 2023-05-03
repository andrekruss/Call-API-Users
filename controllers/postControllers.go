package controllers

import (
	"call-ms-users/database"
	"call-ms-users/helpers"
	"call-ms-users/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func PostUser(ctx *fiber.Ctx) error {

	collection := database.DBConn.Database("call-users").Collection("users")
	user := new(models.User)
	parseErr := ctx.BodyParser(user)

	if parseErr != nil {
		return parseErr
	}

	validateErr := helpers.Validate.Struct(user)

	if validateErr != nil {
		return validateErr
	}

	hashedPassword, hashErr := helpers.HashPassword(user.Password)

	if hashErr != nil {
		return hashErr
	}

	user.Password = hashedPassword

	res, err := collection.InsertOne(ctx.Context(), user)

	log.Println(res.InsertedID)

	if err != nil {
		log.Fatal(err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}
