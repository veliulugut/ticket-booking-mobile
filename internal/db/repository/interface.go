package repository

import "context"

type EventRepository interface {
	GetMany(ctx context.Context)
}
