package controllers

import "github.com/gofiber/fiber/v2"

func GetRoot(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello World!")
}
