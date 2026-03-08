package dto

import (
	"github.com/google/uuid"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/ticket"
)

type (
	ReserveTicketInput struct {
		EventID int64                     `json:"event_id"`
		Tickets ReserveTicketInputTickets `json:"tickets"`
	}

	ReserveTicketInputTickets map[ticket.TicketType]int

	ReserveTicketOutput struct {
		TicketIds []uuid.UUID `json:"ticket_ids"`
	}
)
