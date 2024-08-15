package config

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"ticket-booking-app/pkg/cresponse"
)

type ServeConfig struct {
	Host string
	Port string
}

var FiberConfig = fiber.Config{
	AppName:   "Ticket Booking App",
	BodyLimit: 1024 * 1024 * 50, // 50 MB

	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		var code int = fiber.StatusInternalServerError

		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		log.Error("Error occurred : ", err)

		return cresponse.ErrorResponse(ctx, code, "Unexpected error occurred")
	},
}
