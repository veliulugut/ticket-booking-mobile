package ticket

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetMany(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	CreateOne(c *fiber.Ctx) error
	UpdateOne(c *fiber.Ctx) error
}
