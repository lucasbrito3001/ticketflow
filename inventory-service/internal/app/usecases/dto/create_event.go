package dto

import (
	"time"

	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/address"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/event"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/money"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/salewindow"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/ticket"
)

type (
	CreateEventInput struct {
		Name         string                    `json:"name"`
		Venue        string                    `json:"venue"`
		StartsAt     time.Time                 `json:"starts_at"`
		EndsAt       time.Time                 `json:"ends_at"`
		OpenSalesAt  time.Time                 `json:"open_sales_at"`
		CloseSalesAt time.Time                 `json:"close_sales_at"`
		Street       string                    `json:"street"`
		Number       int                       `json:"number"`
		City         string                    `json:"city"`
		State        string                    `json:"state"`
		ZipCode      string                    `json:"zip_code"`
		CatalogItems []CreateTicketCatalogItem `json:"catalog_items"`
	}

	CreateTicketCatalogItem struct {
		TicketType    string `json:"ticket_type"`
		Price         int64  `json:"price"`
		TotalQuantity int    `json:"total_quantity"`
	}

	CreateEventOutput struct {
		ID      int64  `json:"id"`
		Message string `json:"message"`
	}
)

func (input *CreateEventInput) ToDomain() (*event.Event, error) {
	eventName, err := event.NewEventName(input.Name)
	if err != nil {
		return nil, err
	}

	eventVenue, err := event.NewEventVenue(input.Venue)
	if err != nil {
		return nil, err
	}

	addr, err := address.New(input.Street, input.Number, input.City, input.State, input.ZipCode)
	if err != nil {
		return nil, err
	}

	saleWindow, err := salewindow.New(input.OpenSalesAt, input.CloseSalesAt)
	if err != nil {
		return nil, err
	}

	catalogItems := make([]*event.TicketCatalogItem, 0, len(input.CatalogItems))
	for _, item := range input.CatalogItems {
		ticketType, err := ticket.NewTicketType(item.TicketType)
		if err != nil {
			return nil, err
		}

		price, err := money.FromCents(item.Price)
		if err != nil {
			return nil, err
		}

		quantity, err := event.NewTicketCatalogItemQuantity(item.TotalQuantity)
		if err != nil {
			return nil, err
		}

		catalogItem, err := event.NewTicketCatalogItem(ticketType, price, int(quantity), 0)
		if err != nil {
			return nil, err
		}

		catalogItems = append(catalogItems, catalogItem)
	}

	catalog := event.NewTicketCatalog(catalogItems)

	evt, err := event.New(0, eventName, eventVenue, input.StartsAt, input.EndsAt, catalog, addr, saleWindow)
	if err != nil {
		return nil, err
	}

	return evt, nil
}
