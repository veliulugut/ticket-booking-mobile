package auth

import "github.com/gofiber/fiber/v2"

type AutoHandler interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}
