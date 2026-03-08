package ports

import (
	"context"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/ticket"
)

type TicketRepository interface {
	FindById(ctx context.Context, ticketID string) (*ticket.Ticket, error)
	Create(ctx context.Context, ticket *ticket.Ticket) error
	CreateBatch(ctx context.Context, tickets []*ticket.Ticket) error
}
