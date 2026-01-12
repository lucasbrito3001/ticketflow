package domain

import (
	"time"

	address "github.com/lucasbrito3001/ticketflow-reservation-service/internal/address/domain"
	salewindow "github.com/lucasbrito3001/ticketflow-reservation-service/internal/salewindow/domain"
)

type (
	Event struct {
		id            int
		name          string
		startsAt      time.Time
		endsAt        time.Time
		venueName     string
		address       *address.Address
		saleWindow    *salewindow.SaleWindow
		ticketCatalog TicketCatalog
	}
)

func New(
	id int,
	name string,
	startsAt, endsAt time.Time,
	venueName string,
	address *address.Address,
	saleWindow *salewindow.SaleWindow,
	catalog TicketCatalog,
) (*Event, error) {
	if name == "" {
		return nil, ErrEventNameEmpty
	}

	if venueName == "" {
		return nil, ErrEventVenueNameEmpty
	}

	if endsAt.Before(startsAt) {
		return nil, ErrEventEndsBeforeStarts
	}

	return &Event{
		id:            id,
		name:          name,
		startsAt:      startsAt,
		endsAt:        endsAt,
		venueName:     venueName,
		address:       address,
		ticketCatalog: catalog,
		saleWindow:    saleWindow,
	}, nil
}

func Rehydrate(
	id int,
	name string,
	startsAt,
	endsAt time.Time,
	venueName string,
	address *address.Address,
	saleWindow *salewindow.SaleWindow,
) *Event {
	return &Event{
		id:         id,
		name:       name,
		startsAt:   startsAt,
		endsAt:     endsAt,
		venueName:  venueName,
		address:    address,
		saleWindow: saleWindow,
	}
}

func (e *Event) AttachTicketCatalog(catalog []*TicketCatalogItem) {
	e.ticketCatalog = make(TicketCatalog)
	for _, item := range catalog {
		e.ticketCatalog[item.Type] = item
	}
}
