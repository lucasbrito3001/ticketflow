package usecases

import (
	"context"
	"log/slog"

	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/app/ports"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/app/usecases/dto"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/event"
)

type ConsumeTicket interface {
	Execute(ctx context.Context, command dto.ConsumeTicketInput) (*dto.ConsumeTicketOutput, error)
}

type consumeTicket struct {
	unitOfWork      ports.UnitOfWork
	eventRepository ports.EventRepository
}

func NewConsumeTicket(unitOfWork ports.UnitOfWork, eventRepository ports.EventRepository) ConsumeTicket {
	return &consumeTicket{
		unitOfWork:      unitOfWork,
		eventRepository: eventRepository,
	}
}

func (c *consumeTicket) Execute(ctx context.Context, input dto.ConsumeTicketInput) (*dto.ConsumeTicketOutput, error) {
	e, err := c.eventRepository.FindById(ctx, input.EventID)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find event by id", "event_id", input.EventID, "error", err)
		return nil, err
	}

	if err := c.validateConsumption(ctx, e, input); err != nil {
		return nil, err
	}

	if err := c.unitOfWork.Do(ctx, func(txCtx context.Context) error {
		for ticketType, quantity := range input.Tickets {
			err := c.eventRepository.DecrementAvailableQuantity(txCtx, ticketType, input.EventID, quantity)
			if err != nil {
				slog.ErrorContext(ctx, "failed to decrement ticket availability", "event_id", input.EventID, "error", err)
				return err
			}
		}
		return nil
	}); err != nil {
		slog.ErrorContext(ctx, "transaction failed during ticket consumption", "event_id", input.EventID, "error", err)
		return nil, err
	}

	return &dto.ConsumeTicketOutput{
		Message: "Tickets consumed successfully",
	}, nil
}

func (c *consumeTicket) validateConsumption(
	ctx context.Context,
	event *event.Event,
	input dto.ConsumeTicketInput,
) error {
	for ticketType, quantity := range input.Tickets {
		if err := event.ValidateConsumption(ticketType, quantity); err != nil {
			slog.ErrorContext(ctx, "error on consumption validation",
				"event_id", input.EventID,
				"error", err,
			)
			return err
		}
	}
	return nil
}
