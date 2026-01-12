package ports

import (
	"context"

	event "github.com/lucasbrito3001/ticketflow-reservation-service/internal/event/domain"
)

type EventRepository interface {
	FindById(ctx context.Context, eventID int) (*event.Event, error)
	FindByIDForUpdate(ctx context.Context, eventID int) (*event.Event, error)
	Create(ctx context.Context, event *event.Event) error
	Update(ctx context.Context, event *event.Event) error
}
