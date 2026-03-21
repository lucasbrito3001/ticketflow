package usecases

import (
	"context"
	"log/slog"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/ports"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/usecases/dto"
)

type (
	CreateReservationUseCase interface {
		Execute(ctx context.Context, input *dto.CreateReservationInput) (*dto.CreateReservationOutput, error)
	}

	createReservationUseCase struct {
		reservationRepository  ports.ReservationRepository
		inventoryServiceClient ports.InventoryServiceClient
	}
)

func NewCreateReservationUseCase(reservationRepository ports.ReservationRepository, inventoryServiceClient ports.InventoryServiceClient) CreateReservationUseCase {
	return &createReservationUseCase{
		reservationRepository:  reservationRepository,
		inventoryServiceClient: inventoryServiceClient,
	}
}

func (c *createReservationUseCase) Execute(ctx context.Context, input *dto.CreateReservationInput) (*dto.CreateReservationOutput, error) {
	reservation, err := input.ToDomain()
	if err != nil {
		slog.ErrorContext(ctx, "failed to validate create reservation input", "error", err, "input", input)
		return nil, err
	}

	id, err := c.reservationRepository.Create(ctx, reservation)
	if err != nil {
		slog.ErrorContext(ctx, "failed to persist reservation", "error", err, "event_id", reservation.EventId(), "input", input)
		return nil, err
	}

	reservation.SetId(id)
	reservation.SetAsAwaitingPayment()

	if err := c.inventoryServiceClient.HoldTickets(ctx, reservation.EventId(), reservation.Tickets()); err != nil {
		slog.ErrorContext(ctx, "failed to hold tickets in inventory, cancelling reservation", "error", err, "reservation_id", reservation.Id(), "event_id", reservation.EventId(), "tickets", reservation.Tickets())
		reservation.SetAsCancelled()
		if err := c.reservationRepository.Update(ctx, reservation); err != nil {
			slog.ErrorContext(ctx, "failed to update reservation status to cancelled", "error", err, "reservation_id", reservation.Id())
			return nil, err
		}
		return nil, err
	}

	if err := c.reservationRepository.Update(ctx, reservation); err != nil {
		slog.ErrorContext(ctx, "failed to update reservation status to awaiting payment", "error", err, "reservation_id", reservation.Id())
		return nil, err
	}

	return &dto.CreateReservationOutput{
		ReservationID: id,
	}, nil
}
