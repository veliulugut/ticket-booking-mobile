package repository

import (
	"context"

	"gorm.io/gorm"

	"ticket-booking-app/internal/db/models"
)

var _ EventRepository = (*Event)(nil)

func NewEventRepository(dbClient *gorm.DB) *Event {
	return &Event{
		dbClient: dbClient,
	}
}

type Event struct {
	dbClient  *gorm.DB
	tableName string
}

func (e *Event) UpdateEvent(ctx context.Context, eventID uint, updateData map[string]interface{}) (*models.Event, error) {
	event := &models.Event{}

	updateRes := e.dbClient.Model(event).Where("id = ?", eventID).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := e.dbClient.Model(event).Where("id = ?", eventID).First(event)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return event, nil
}
func (e *Event) DeleteEvent(ctx context.Context, eventID uint) error {
	res := e.dbClient.WithContext(ctx).Delete(&models.Event{}, eventID)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (e *Event) CreateEvent(ctx context.Context, event *models.Event) (*models.Event, error) {
	res := e.dbClient.Model(event).WithContext(ctx).Create(event)

	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}

// GetByID implements EventRepository.
func (e *Event) GetByID(ctx context.Context, id string) (*models.Event, error) {
	var ticket models.Event

	result := e.dbClient.Table(e.tableName).WithContext(ctx).Where("id = ? ", id).First(&ticket)

	if result.Error != nil {
		return nil, result.Error
	}

	return &ticket, nil
}

// GetMany implements EventRepository.
func (e *Event) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	res := e.dbClient.Model(&models.Event{}).Order("updated_at desc").WithContext(ctx).Find(&events)

	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil
}
