package rest

import (
	"errors"
	"net/http"

	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/address"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/event"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/money"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/ticket"
)

type ErrorMapperResultCode struct {
	StatusCode int
	ErrorCode  string
}

type ErrorMapperResult struct {
	StatusCode int
	Error      ErrorResponse
}

var errorStatusMap = map[error]ErrorMapperResultCode{
	// Event domain errors
	event.ErrEventNameEmpty:        {StatusCode: http.StatusBadRequest, ErrorCode: "EVENT_NAME_EMPTY"},
	event.ErrEventVenueNameEmpty:   {StatusCode: http.StatusBadRequest, ErrorCode: "EVENT_VENUE_NAME_EMPTY"},
	event.ErrEventEndsBeforeStarts: {StatusCode: http.StatusBadRequest, ErrorCode: "EVENT_ENDS_BEFORE_STARTS"},
	event.ErrTicketTypeNotFound:    {StatusCode: http.StatusBadRequest, ErrorCode: "TICKET_TYPE_NOT_FOUND"},
	event.ErrTicketCatalogNotSet:   {StatusCode: http.StatusBadRequest, ErrorCode: "TICKET_CATALOG_NOT_SET"},
	event.ErrInsufficientTickets:   {StatusCode: http.StatusConflict, ErrorCode: "INSUFFICIENT_TICKETS"},
	event.ErrEventNotFound:         {StatusCode: http.StatusNotFound, ErrorCode: "EVENT_NOT_FOUND"},

	// Address domain errors
	address.ErrStreetEmpty:  {StatusCode: http.StatusBadRequest, ErrorCode: "STREET_EMPTY"},
	address.ErrNumberEmpty:  {StatusCode: http.StatusBadRequest, ErrorCode: "NUMBER_EMPTY"},
	address.ErrCityEmpty:    {StatusCode: http.StatusBadRequest, ErrorCode: "CITY_EMPTY"},
	address.ErrStateEmpty:   {StatusCode: http.StatusBadRequest, ErrorCode: "STATE_EMPTY"},
	address.ErrZipCodeEmpty: {StatusCode: http.StatusBadRequest, ErrorCode: "ZIP_CODE_EMPTY"},

	// Money domain errors
	money.ErrInvalidMoneyAmount: {StatusCode: http.StatusBadRequest, ErrorCode: "INVALID_MONEY_AMOUNT"},

	// Ticket domain errors
	ticket.ErrInvalidTicketType:     {StatusCode: http.StatusBadRequest, ErrorCode: "INVALID_TICKET_TYPE"},
	ticket.ErrInvalidTicketQuantity: {StatusCode: http.StatusBadRequest, ErrorCode: "INVALID_TICKET_QUANTITY"},
}

// func MapError(err error) ErrorMapperResult {
// 	if result, exists := errorStatusMap[err]; exists {
// 		return ErrorMapperResult{StatusCode: result.StatusCode, Error: ErrorResponse{Error: err.Error(), ErrorCode: result.ErrorCode}}
// 	}

// 	return ErrorMapperResult{StatusCode: http.StatusInternalServerError, Error: ErrorResponse{Error: "An unexpected error occurred"}}
// }

func MapError(err error) ErrorMapperResult {
	for domainErr, result := range errorStatusMap {
		if errors.Is(err, domainErr) {
			return ErrorMapperResult{
				StatusCode: result.StatusCode,
				Error: ErrorResponse{
					Error:     err.Error(),
					ErrorCode: result.ErrorCode,
				},
			}
		}
	}

	return ErrorMapperResult{
		StatusCode: http.StatusInternalServerError,
		Error: ErrorResponse{
			Error: "an unexpected error occurred",
		},
	}
}
