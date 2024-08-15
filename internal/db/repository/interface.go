package repository

import (
	"context"
	"ticket-booking-app/internal/db/models"
)

type EventRepository interface {
	GetMany(ctx context.Context) ([]*models.Event, error)
	GetByID(ctx context.Context, id string) (*models.Event, error)
	CreateEvent(ctx context.Context, event *models.Event) (*models.Event, error)
	UpdateEvent(ctx context.Context, eventID uint, updateData map[string]interface{}) (*models.Event, error)
	DeleteEvent(ctx context.Context, eventID uint) error
}

type AuthRepository interface {
	RegisterUser(ctx context.Context, registerData *models.AuthCredentials) (*models.User, error)
	GetUser(ctx context.Context, query interface{}, args ...interface{}) (*models.User, error)
}

type TicketRepository interface {
	GetMany(ctx context.Context, userId uint) ([]*models.Ticket, error)
	GetOne(ctx context.Context, userId uint, ticketId uint) (*models.Ticket, error)
	CreateOne(ctx context.Context, userId uint, ticket *models.Ticket) (*models.Ticket, error)
	UpdateOne(ctx context.Context, userId, ticket uint, updateData map[string]interface{}) (*models.Ticket, error)
}
