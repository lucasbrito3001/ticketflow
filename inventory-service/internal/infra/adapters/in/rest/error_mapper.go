package rest

import (
	"net/http"

	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/address"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/event"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/money"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/ticket"
)

type ErrorResponse struct {
	StatusCode int
	Message    string
}

var errorStatusMap = map[error]int{
	// Event domain errors
	event.ErrEventNameEmpty:        http.StatusBadRequest,
	event.ErrEventVenueNameEmpty:   http.StatusBadRequest,
	event.ErrEventEndsBeforeStarts: http.StatusBadRequest,
	event.ErrTicketTypeNotFound:    http.StatusBadRequest,
	event.ErrTicketCatalogNotSet:   http.StatusBadRequest,
	event.ErrInsufficientTickets:   http.StatusConflict,
	event.ErrEventNotFound:         http.StatusNotFound,

	// Address domain errors
	address.ErrStreetEmpty:  http.StatusBadRequest,
	address.ErrNumberEmpty:  http.StatusBadRequest,
	address.ErrCityEmpty:    http.StatusBadRequest,
	address.ErrStateEmpty:   http.StatusBadRequest,
	address.ErrZipCodeEmpty: http.StatusBadRequest,

	// Money domain errors
	money.ErrInvalidMoneyAmount: http.StatusBadRequest,

	// Ticket domain errors
	ticket.ErrInvalidTicketType:     http.StatusBadRequest,
	ticket.ErrInvalidTicketQuantity: http.StatusBadRequest,
}

func MapError(err error) ErrorResponse {
	if statusCode, exists := errorStatusMap[err]; exists {
		return ErrorResponse{statusCode, err.Error()}
	}

	return ErrorResponse{http.StatusInternalServerError, "An unexpected error occurred"}
}
