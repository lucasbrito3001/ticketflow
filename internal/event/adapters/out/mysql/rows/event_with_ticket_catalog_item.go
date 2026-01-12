package rows

import (
	"time"

	address "github.com/lucasbrito3001/ticketflow-reservation-service/internal/address/domain"
	event "github.com/lucasbrito3001/ticketflow-reservation-service/internal/event/domain"
	salewindow "github.com/lucasbrito3001/ticketflow-reservation-service/internal/salewindow/domain"
)

type EventWithTicketCatalogItemRow struct {
	EventID      int
	Name         string
	StartsAt     time.Time
	EndsAt       time.Time
	StartOfSales time.Time
	EndOfSales   time.Time
	VenueName    string

	TicketType string
	Total      int
	Sold       int

	Street  string
	Number  int
	City    string
	State   string
	ZipCode string
}

func (r *EventWithTicketCatalogItemRow) ToDomain() *event.Event {
	return event.Rehydrate(
		r.EventID,
		r.Name,
		r.StartsAt,
		r.EndsAt,
		r.VenueName,
		address.Rehydrate(
			r.Street,
			r.Number,
			r.City,
			r.State,
			r.ZipCode,
		),
		salewindow.Rehydrate(
			r.StartOfSales,
			r.EndOfSales,
		),
	)
}

func (r *EventWithTicketCatalogItemRow) ToTicketCatalogItemDomain() *event.TicketCatalogItem {
	return event.RehydrateTicketCatalogItem(
		event.TicketType(r.TicketType),
		r.Total,
		r.Sold,
	)
}
