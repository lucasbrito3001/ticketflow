package ticket

import (
	"github.com/google/uuid"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/money"
)

type (
	Tickets []*Ticket
	Ticket  struct {
		id         uuid.UUID
		eventId    int64
		ticketType TicketType
		price      money.Money
	}
)

func NewTicket(
	eventID int64,
	ticketType TicketType,
	price money.Money,
) (*Ticket, error) {
	uuid := uuid.New()

	return &Ticket{
		id:         uuid,
		eventId:    eventID,
		ticketType: ticketType,
		price:      price,
	}, nil
}

func (t *Ticket) Id() uuid.UUID {
	return t.id
}

func (t *Ticket) EventId() int64 {
	return t.eventId
}

func (t *Ticket) Type() TicketType {
	return t.ticketType
}

func (t *Ticket) Price() money.Money {
	return t.price
}

func Rehydrate(id string, eventId int64, ticketType TicketType, price int64) *Ticket {
	return &Ticket{
		id:         uuid.MustParse(id),
		eventId:    eventId,
		ticketType: ticketType,
		price:      money.Money(price),
	}
}

func (ts Tickets) IDs() []uuid.UUID {
	ids := make([]uuid.UUID, 0, len(ts))

	for _, t := range ts {
		ids = append(ids, t.Id())
	}

	return ids
}
