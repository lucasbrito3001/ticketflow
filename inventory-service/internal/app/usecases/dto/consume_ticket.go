package dto

import (
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/ticket"
)

type (
	ConsumeTicketInput struct {
		EventID int64                     `uri:"event_id"`
		Tickets ConsumeTicketInputTickets `json:"tickets"`
	}

	ConsumeTicketInputTickets map[ticket.TicketType]int

	ConsumeTicketOutput struct {
		Message string `json:"message"`
	}
)
