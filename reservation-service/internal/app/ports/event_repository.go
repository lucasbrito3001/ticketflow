package ports

import (
	"context"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/event"
)

type EventRepository interface {
	FindById(ctx context.Context, eventID int64) (*event.Event, error)
	Create(ctx context.Context, event *event.Event) error
	Update(ctx context.Context, event *event.Event) error
}
