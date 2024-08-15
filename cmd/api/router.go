package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ticket-booking-app/cmd/api/handlers/v1/auth"
	"ticket-booking-app/internal/db/repository"
	"ticket-booking-app/internal/service"
)

func InitializeRouter(app *fiber.App, connection *gorm.DB) {
	//repository
	authRepository := repository.NewAuthRepository(connection)

	//Services
	authService := service.NewAuthService(authRepository)

	//Handlers
	authHandler := auth.NewAuthHandler(*authService)

	v1 := app.Group("/v1")

	authRouter := v1.Group("/auth")

	authRouter.Post("/login", authHandler.Login)
	authRouter.Post("/register", authHandler.Register)

}
