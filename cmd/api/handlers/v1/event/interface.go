package event

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetMany(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	CreateEvent(c *fiber.Ctx) error
	UpdateEvent(c *fiber.Ctx) error
	DeleteEvent(c *fiber.Ctx) error
}
