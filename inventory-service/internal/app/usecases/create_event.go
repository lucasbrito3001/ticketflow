package usecases

import (
	"context"
	"log/slog"

	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/app/ports"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/app/usecases/dto"
)

type CreateEvent interface {
	Execute(ctx context.Context, input dto.CreateEventInput) (*dto.CreateEventOutput, error)
}

type createEvent struct {
	unitOfWork      ports.UnitOfWork
	eventRepository ports.EventRepository
}

func NewCreateEvent(unitOfWork ports.UnitOfWork, eventRepository ports.EventRepository) CreateEvent {
	return &createEvent{
		unitOfWork:      unitOfWork,
		eventRepository: eventRepository,
	}
}

func (c *createEvent) Execute(ctx context.Context, input dto.CreateEventInput) (*dto.CreateEventOutput, error) {

	event, err := input.ToDomain()
	if err != nil {
		slog.ErrorContext(ctx, "failed to convert input to domain", "error", err, "event_name", input.Name)
		return nil, err
	}

	if err := c.unitOfWork.Do(ctx, func(txCtx context.Context) error {
		if err := c.eventRepository.CreateEvent(txCtx, event); err != nil {
			return err
		}
		return c.eventRepository.CreateEventTicketCatalogItems(txCtx, event)
	}); err != nil {
		slog.ErrorContext(ctx, "failed to create event", "error", err, "event_name", input.Name)
		return nil, err
	}

	return &dto.CreateEventOutput{
		ID:      event.ID(),
		Message: "Event created successfully",
	}, nil
}
