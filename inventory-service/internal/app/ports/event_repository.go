package ports

import (
	"context"

	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/event"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/ticket"
)

type EventRepository interface {
	FindById(ctx context.Context, eventID int64) (*event.Event, error)
	CreateEvent(ctx context.Context, event *event.Event) error
	CreateEventTicketCatalogItems(ctx context.Context, event *event.Event) error
	UpdateEvent(ctx context.Context, event *event.Event) error
	UpdateEventTicketCatalogItem(ctx context.Context, event *event.Event) error
	DecrementAvailableQuantity(ctx context.Context, ticketType ticket.TicketType, eventID int64, quantityDelta int) error
}
