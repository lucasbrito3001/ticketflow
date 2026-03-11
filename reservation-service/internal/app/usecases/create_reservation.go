package usecases

import (
	"context"

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
		return nil, err
	}

	id, err := c.reservationRepository.Create(ctx, reservation)
	if err != nil {
		return nil, err
	}

	reservation.SetAsAwaitingPayment()

	if err := c.inventoryServiceClient.HoldTickets(ctx, reservation.EventID(), reservation.Tickets()); err != nil {
		reservation.SetAsCancelled()
	}

	if err := c.reservationRepository.Update(ctx, reservation); err != nil {
		return nil, err
	}

	return &dto.CreateReservationOutput{
		ReservationID: id,
	}, nil
}
