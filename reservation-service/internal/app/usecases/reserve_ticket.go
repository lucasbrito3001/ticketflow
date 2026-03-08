package usecases

import (
	"context"
	"log/slog"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/ports"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/usecases/dto"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/event"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/ticket"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/infra/clients"
)

type (
	ReserveTicket interface {
		Execute(ctx context.Context, command dto.ReserveTicketInput) (*dto.ReserveTicketOutput, error)
	}

	reserveTicket struct {
		unitOfWork       ports.UnitOfWork
		eventRepository  ports.EventRepository
		ticketRepository ports.TicketRepository
		inventoryClient  clients.InventoryServiceClient
	}
)

func NewReserveTicket(unitOfWork ports.UnitOfWork, eventRepository ports.EventRepository, ticketRepository ports.TicketRepository, inventoryClient clients.InventoryServiceClient) ReserveTicket {
	return &reserveTicket{
		unitOfWork,
		eventRepository,
		ticketRepository,
		inventoryClient,
	}
}

func (r *reserveTicket) Execute(ctx context.Context, input dto.ReserveTicketInput) (*dto.ReserveTicketOutput, error) {
	event, err := r.eventRepository.FindById(ctx, input.EventID)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find event by id", "event_id", input.EventID, "error", err)
		return nil, err
	}

	if err := r.validatePurchaseIntent(ctx, event, input); err != nil {
		return nil, err
	}

	// Call inventory service to consume tickets
	if err := r.inventoryClient.ConsumeTicket(ctx, input.EventID, input.Tickets); err != nil {
		slog.ErrorContext(ctx, "failed to consume tickets from inventory service", "event_id", input.EventID, "error", err)
		return nil, err
	}

	tickets := make(ticket.Tickets, 0)
	if err = r.unitOfWork.Do(ctx, func(ctx context.Context) error {
		tickets, err = r.buildTickets(event, input)
		if err != nil {
			return err
		}

		if err := r.ticketRepository.CreateBatch(ctx, tickets); err != nil {
			return err
		}

		return nil
	}); err != nil {
		slog.ErrorContext(ctx, "failed to execute unit of work for reserving ticket", "event_id", input.EventID, "error", err)
		return nil, err
	}

	return &dto.ReserveTicketOutput{
		TicketIds: tickets.IDs(),
	}, nil
}

func (r *reserveTicket) validatePurchaseIntent(
	ctx context.Context,
	event *event.Event,
	input dto.ReserveTicketInput,
) error {
	for ticketType, quantity := range input.Tickets {
		if err := event.ValidatePurchaseIntent(ticketType, quantity); err != nil {
			slog.ErrorContext(ctx, "error on purchase intent",
				"event_id", input.EventID,
				"error", err,
			)
			return err
		}
	}
	return nil
}

func (r *reserveTicket) buildTickets(
	event *event.Event,
	input dto.ReserveTicketInput,
) (ticket.Tickets, error) {

	total := totalQuantity(input.Tickets)
	tickets := make(ticket.Tickets, 0, total)

	for ticketType, quantity := range input.Tickets {
		price, err := event.TicketCatalogItemPrice(ticketType)
		if err != nil {
			return nil, err
		}

		for i := 0; i < quantity; i++ {
			t, err := ticket.NewTicket(input.EventID, ticketType, price)
			if err != nil {
				return nil, err
			}
			tickets = append(tickets, t)
		}
	}

	return tickets, nil
}

func totalQuantity(m map[ticket.TicketType]int) int {
	total := 0
	for _, q := range m {
		total += q
	}
	return total
}
