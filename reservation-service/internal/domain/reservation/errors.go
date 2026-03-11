package reservation

import "errors"

var (
	ErrInvalidEventID             = errors.New("The event ID is invalid")
	ErrEmptyTickets               = errors.New("The tickets are empty")
	ErrInvalidTicketQuantity      = errors.New("The ticket quantity is invalid")
	ErrInvalidTicketType          = errors.New("The ticket type is invalid")
	ErrCannotSetAsCancelled       = errors.New("Cannot cancel reservation with current status")
	ErrCannotSetAsAwaitingPayment = errors.New("Cannot set reservation as awaiting payment with current status")
	ErrInvalidReservationStatus   = errors.New("The ticket status is invalid")
)
