package domain

import (
	money "github.com/lucasbrito3001/ticketflow-reservation-service/internal/money/domain"
)

type (
	TicketCatalog map[TicketType]*TicketCatalogItem

	TicketCatalogItem struct {
		Type          TicketType
		Price         money.Money
		TotalQuantity int
		SoldQuantity  int
	}
)

func NewTicketCatalogItem(
	ticketType TicketType,
	totalQuantity int,
	soldQuantity int,
) (TicketCatalogItem, error) {
	return TicketCatalogItem{
		Type:          ticketType,
		TotalQuantity: totalQuantity,
		SoldQuantity:  soldQuantity,
	}, nil
}

func RehydrateTicketCatalogItem(ticketType TicketType, total, sold int) *TicketCatalogItem {
	return &TicketCatalogItem{
		Type:          ticketType,
		TotalQuantity: total,
		SoldQuantity:  sold,
	}
}
