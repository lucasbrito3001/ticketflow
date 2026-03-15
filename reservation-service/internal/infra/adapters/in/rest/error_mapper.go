package rest

import (
	"net/http"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/reservation"
)

type ErrorMapperResult struct {
	StatusCode int
	Error      ErrorResponse
}

var errorStatusMap = map[error]int{
	// Reservation domain errors
	reservation.ErrCannotSetAsAwaitingPayment: http.StatusBadRequest,
	reservation.ErrCannotSetAsCancelled:       http.StatusBadRequest,
	reservation.ErrEmptyTickets:               http.StatusBadRequest,
	reservation.ErrInvalidEventID:             http.StatusBadRequest,
	reservation.ErrInvalidReservationStatus:   http.StatusBadRequest,
	reservation.ErrInvalidTicketQuantity:      http.StatusBadRequest,
	reservation.ErrInvalidTicketType:          http.StatusBadRequest,
}

func MapError(err error) ErrorMapperResult {
	if statusCode, exists := errorStatusMap[err]; exists {
		return ErrorMapperResult{statusCode, ErrorResponse{Error: err.Error()}}
	}

	return ErrorMapperResult{http.StatusInternalServerError, ErrorResponse{Error: "An unexpected error occurred"}}
}
