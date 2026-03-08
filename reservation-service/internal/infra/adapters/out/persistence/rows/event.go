package rows

import (
	"time"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/address"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/event"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/salewindow"
)

type EventRow struct {
	// Event
	EventID  int64
	Name     string
	StartsAt time.Time
	EndsAt   time.Time

	// SaleWindow
	StartOfSales time.Time
	EndOfSales   time.Time

	// Address
	VenueName string
	Street    string
	Number    int
	City      string
	State     string
	ZipCode   string
}

func (r *EventRow) ToDomain() *event.Event {
	return event.Rehydrate(
		r.EventID,
		event.RehydrateEventName(r.Name),
		event.RehydrateEventVenue(r.VenueName),
		r.StartsAt,
		r.EndsAt,
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
