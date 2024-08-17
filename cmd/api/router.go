package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ticket-booking-app/cmd/api/handlers/v1/auth"
	"ticket-booking-app/cmd/api/handlers/v1/event"
	"ticket-booking-app/cmd/api/handlers/v1/ticket"
	"ticket-booking-app/internal/db/repository"
	"ticket-booking-app/internal/service"
)

func InitializeRouter(app *fiber.App, connection *gorm.DB) {
	//repository
	authRepository := repository.NewAuthRepository(connection)
	eventRepository := repository.NewEventRepository(connection)
	ticketRepository := repository.NewTicketRepository(connection)

	//Services
	authService := service.NewAuthService(authRepository)

	//Handlers
	authHandler := auth.NewAuthHandler(*authService)
	eventHandler := event.NewEventHandler(*eventRepository)
	ticketHandler := ticket.NewTicketHandler(*ticketRepository)

	v1 := app.Group("/v1")

	authRouter := v1.Group("/auth")

	authRouter.Post("/login", authHandler.Login)
	authRouter.Post("/register", authHandler.Register)

	eventRouter := v1.Group("/event")
	eventRouter.Get("/", eventHandler.GetMany)
	eventRouter.Post("/", eventHandler.CreateEvent)
	eventRouter.Get("/:eventId", eventHandler.GetByID)
	eventRouter.Put("/:eventId", eventHandler.UpdateEvent)
	eventRouter.Delete("/:eventId", eventHandler.DeleteEvent)

	ticketRouter := v1.Group("/ticket")
	ticketRouter.Get("/", ticketHandler.GetMany)
	ticketRouter.Post("/", ticketHandler.CreateOne)
	ticketRouter.Get("/:ticketId", ticketHandler.GetOne)
	ticketRouter.Put("/:ticketId", ticketHandler.UpdateOne)

}
