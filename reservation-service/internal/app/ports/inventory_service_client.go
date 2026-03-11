package ports

import (
	"context"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/reservation"
)

type InventoryServiceClient interface {
	HoldTickets(ctx context.Context, eventID int64, tickets map[reservation.ReservationTicketType]int) error
}
