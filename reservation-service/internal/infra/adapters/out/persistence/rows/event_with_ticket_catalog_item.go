package rows

import (
	"time"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/address"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/event"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/salewindow"
)

type (
	EventWithTicketCatalogItemRows []*EventWithTicketCatalogItemRow

	EventWithTicketCatalogItemRow struct {
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

		// TicketCatalogItem
		TicketType string
		Price      int
		Total      int
		Sold       int
	}
)

func (rows EventWithTicketCatalogItemRows) ToDomain() *event.Event {
	if len(rows) == 0 {
		return nil
	}

	items := make([]*event.TicketCatalogItem, len(rows))
	for i, row := range rows {
		items[i] = event.RehydrateTicketCatalogItem(
			row.TicketType,
			row.Price,
			row.Total,
			row.Sold,
		)
	}

	first := rows[0]
	evt := event.RehydrateWithCatalog(
		first.EventID,
		event.RehydrateEventName(first.Name),
		event.RehydrateEventVenue(first.VenueName),
		first.StartsAt,
		first.EndsAt,
		event.NewTicketCatalog(items),
		address.Rehydrate(
			first.Street,
			first.Number,
			first.City,
			first.State,
			first.ZipCode,
		),
		salewindow.Rehydrate(
			first.StartOfSales,
			first.EndOfSales,
		),
	)

	return evt
}
