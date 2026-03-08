package event

import (
	"time"

	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/address"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/money"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/salewindow"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/ticket"
)

type (
	Event struct {
		id         int64
		name       EventName
		venue      EventVenue
		startsAt   time.Time
		endsAt     time.Time
		catalog    *TicketCatalog
		address    *address.Address
		saleWindow *salewindow.SaleWindow
	}
)

func New(
	id int64,
	name EventName,
	venue EventVenue,
	startsAt, endsAt time.Time,
	catalog *TicketCatalog,
	address *address.Address,
	saleWindow *salewindow.SaleWindow,
) (*Event, error) {
	if name == "" {

		return nil, ErrEventNameEmpty
	}

	if venue == "" {
		return nil, ErrEventVenueNameEmpty
	}

	if endsAt.Before(startsAt) {
		return nil, ErrEventEndsBeforeStarts
	}

	return &Event{
		id:         id,
		name:       name,
		startsAt:   startsAt,
		endsAt:     endsAt,
		venue:      venue,
		address:    address,
		saleWindow: saleWindow,
		catalog:    catalog,
	}, nil
}

func (e *Event) ID() int64 {
	return e.id
}

func (e *Event) Name() string {
	return e.name.String()
}

func (e *Event) StartsAt() time.Time {
	return e.startsAt
}

func (e *Event) EndsAt() time.Time {
	return e.endsAt
}

func (e *Event) Venue() string {
	return e.venue.String()
}

func (e *Event) TicketCatalog() *TicketCatalog {
	return e.catalog
}

func (e *Event) Address() *address.Address {
	return e.address
}

func (e *Event) SaleWindow() *salewindow.SaleWindow {
	return e.saleWindow
}

func (e *Event) TicketCatalogItemPrice(ticketType ticket.TicketType) (money.Money, error) {
	item, exists := (*e.catalog)[ticketType]
	if !exists {
		return 0, ErrTicketTypeNotFound
	}

	return item.Price(), nil
}

func Rehydrate(
	id int64,
	name EventName,
	venue EventVenue,
	startsAt time.Time,
	endsAt time.Time,
	address *address.Address,
	saleWindow *salewindow.SaleWindow,
) *Event {
	return &Event{
		id:         id,
		name:       name,
		startsAt:   startsAt,
		endsAt:     endsAt,
		venue:      venue,
		address:    address,
		saleWindow: saleWindow,
	}
}

func RehydrateWithCatalog(
	id int64,
	name EventName,
	venue EventVenue,
	startsAt,
	endsAt time.Time,
	catalog *TicketCatalog,
	address *address.Address,
	saleWindow *salewindow.SaleWindow,
) *Event {
	return &Event{
		id:         id,
		name:       name,
		startsAt:   startsAt,
		endsAt:     endsAt,
		venue:      venue,
		catalog:    catalog,
		address:    address,
		saleWindow: saleWindow,
	}
}

func (e *Event) ValidateConsumption(ticketType ticket.TicketType, quantity int) error {
	if e.catalog == nil {
		return ErrTicketCatalogNotSet
	}

	item, exists := (*e.catalog)[ticketType]
	if !exists {
		return ErrTicketTypeNotFound
	}

	if quantity > item.AvailableQuantity() {
		return ErrInsufficientTickets
	}

	return nil
}
