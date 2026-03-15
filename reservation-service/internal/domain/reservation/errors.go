package reservation

import "errors"

var (
	ErrInvalidEventID             = errors.New("the event ID is invalid, must be an integer greater than 0")
	ErrEmptyTickets               = errors.New("the tickets list cannot be empty")
	ErrInvalidTicketQuantity      = errors.New("the ticket quantity is invalid, must be an integer greater than 0")
	ErrInvalidTicketType          = errors.New("the ticket type is invalid, must be one of: " + ValidTicketTypesString())
	ErrCannotSetAsCancelled       = errors.New("cannot cancel reservation with current status")
	ErrCannotSetAsAwaitingPayment = errors.New("cannot set reservation as awaiting payment with current status")
	ErrInvalidReservationStatus   = errors.New("the reservation status is invalid, must be one of: " + ValidReservationStatusesString())
)
