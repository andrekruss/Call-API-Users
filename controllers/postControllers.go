package controllers

import (
	"call-ms-users/database"
	"call-ms-users/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func PostUser(ctx *fiber.Ctx) error {

	collection := database.DBConn.Database("call-users").Collection("users")

	// creating a dummy user
	user := models.User{
		Id:       1,
		Username: "Andre",
		Password: "123",
		UserType: "admin",
		IsActive: true,
	}

	res, err := collection.InsertOne(ctx.Context(), user)

	fmt.Printf("Inserted id: %v", res.InsertedID)

	if err != nil {
		log.Fatal(err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}
