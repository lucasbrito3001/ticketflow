package event

import "errors"

var (
	ErrEventNameEmpty        = errors.New("The event name can't be empty")
	ErrEventEndsBeforeStarts = errors.New("The event cannot end before it starts")
	ErrEventVenueNameEmpty   = errors.New("The event venue name can't be empty")
	ErrTicketTypeNotFound    = errors.New("The ticket type was not found in the event ticket catalog")
	ErrInsufficientTickets   = errors.New("There are not enough tickets available for the requested ticket type")
	ErrTicketCatalogNotSet   = errors.New("The event ticket catalog is not set")
	ErrEventNotFound         = errors.New("The event was not found")
)
