package ports

import (
	"context"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/reservation"
)

type ReservationRepository interface {
	FindById(ctx context.Context, ticketID string) (*reservation.Reservation, error)
	Create(ctx context.Context, ticket *reservation.Reservation) (int64, error)
	Update(ctx context.Context, ticket *reservation.Reservation) error
}
