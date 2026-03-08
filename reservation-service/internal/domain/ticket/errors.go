package ticket

import "errors"

var (
	ErrInvalidTicketType     = errors.New("The ticket type is invalid")
	ErrInvalidTicketQuantity = errors.New("The ticket quantity is invalid")
)
