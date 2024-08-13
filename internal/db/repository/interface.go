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
