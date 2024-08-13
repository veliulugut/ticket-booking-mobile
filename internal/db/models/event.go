package models

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	ID                    uint      `json:"id" gorm:"primaryKey"`
	Name                  string    `json:"name"`
	Location              string    `json:"location"`
	TotalTicketsPurchased int64     `json:"totalTicketsPurchased" gorm:"-"`
	TotalTicketsEntered   int64     `json:"totalTicketsEntered" gorm:"-"`
	Date                  time.Time `json:"date"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}

func (Event) TableName() string {
	return "events"
}

func (e *Event) AfterFind(db *gorm.DB) (err error) {
	baseQuery := db.Model(&Ticket{}).Where(&Ticket{EventID: e.ID})

	if res := baseQuery.Count(&e.TotalTicketsPurchased); res.Error != nil {
		return res.Error
	}

	if res := baseQuery.Where("entered = ? ", true).Count(&e.TotalTicketsEntered); res.Error != nil {
		return res.Error
	}

	return nil
}
