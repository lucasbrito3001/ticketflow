package event

import (
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/money"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/ticket"
)

type (
	TicketCatalog map[ticket.TicketType]*TicketCatalogItem

	TicketCatalogItem struct {
		ticketType    ticket.TicketType
		price         money.Money
		totalQuantity int
		soldQuantity  int
	}

	TicketCatalogItemQuantity int
)

func NewTicketCatalog(items []*TicketCatalogItem) *TicketCatalog {
	catalog := make(TicketCatalog)

	for _, item := range items {
		catalog[item.ticketType] = item
	}

	return &catalog
}

func NewTicketCatalogItem(
	ticketType ticket.TicketType,
	price money.Money,
	totalQuantity int,
	soldQuantity int,
) (*TicketCatalogItem, error) {
	return &TicketCatalogItem{
		ticketType,
		price,
		totalQuantity,
		soldQuantity,
	}, nil
}

func (tci *TicketCatalogItem) TicketType() ticket.TicketType {
	return tci.ticketType
}

func (tci *TicketCatalogItem) Price() money.Money {
	return tci.price
}

func (tci *TicketCatalogItem) TotalQuantity() int {
	return tci.totalQuantity
}

func (tci *TicketCatalogItem) SoldQuantity() int {
	return tci.soldQuantity
}

func (tci *TicketCatalogItem) AvailableQuantity() int {
	return tci.totalQuantity - tci.soldQuantity
}

func RehydrateTicketCatalogItem(
	ticketType string,
	price, total, sold int,
) *TicketCatalogItem {
	return &TicketCatalogItem{
		ticketType:    ticket.TicketType(ticketType),
		price:         money.Money(price),
		totalQuantity: total,
		soldQuantity:  sold,
	}
}

func NewTicketCatalogItemQuantity(quantity int) (TicketCatalogItemQuantity, error) {
	if quantity <= 0 {
		return 0, ticket.ErrInvalidTicketQuantity
	}

	return TicketCatalogItemQuantity(quantity), nil
}
