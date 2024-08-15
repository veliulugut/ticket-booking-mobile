package repository

import (
	"context"
	"gorm.io/gorm"
	"ticket-booking-app/internal/db/models"
)

var _ TicketRepository = (*Ticket)(nil)

func NewTicketRepository(db *gorm.DB) *Ticket {
	return &Ticket{
		dbClient: db,
	}
}

type Ticket struct {
	dbClient *gorm.DB
}

func (t *Ticket) GetMany(ctx context.Context, userId uint) ([]*models.Ticket, error) {
	tickets := []*models.Ticket{}

	res := t.dbClient.Model(tickets).WithContext(ctx).Where("user_id = ?", userId).Preload("Event").Order("created_at desc").Find(&tickets)
	if res.Error != nil {
		return nil, res.Error
	}

	return tickets, nil
}

func (t *Ticket) GetOne(ctx context.Context, userId uint, ticketId uint) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	res := t.dbClient.Model(ticket).WithContext(ctx).Where("id = ?", ticketId).Where("user_id = ?", userId).First(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (t *Ticket) CreateOne(ctx context.Context, userId uint, ticket *models.Ticket) (*models.Ticket, error) {
	ticket.UserID = userId

	res := t.dbClient.Model(ticket).WithContext(ctx).Create(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return t.GetOne(ctx, userId, ticket.ID)
}

func (t *Ticket) UpdateOne(ctx context.Context, userId, ticketId uint, updateData map[string]interface{}) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	updateRes := t.dbClient.Model(ticket).WithContext(ctx).Where("id = ?", ticketId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	return t.GetOne(ctx, userId, ticketId)

}
