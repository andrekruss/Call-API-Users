package controllers

import (
	"call-ms-users/models"

	"github.com/gofiber/fiber/v2"
)

func PostUser(ctx *fiber.Ctx) error {
	// creating a dummy user
	user := models.User{
		Id:       1,
		Username: "Andre",
		Password: "123",
		UserType: "admin",
		IsActive: true,
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}
