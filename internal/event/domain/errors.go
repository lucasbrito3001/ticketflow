package domain

import "errors"

var (
	ErrEventNameEmpty        = errors.New("The event name can't be empty")
	ErrEventEndsBeforeStarts = errors.New("The event cannot end before it starts")
	ErrEventVenueNameEmpty   = errors.New("The event venue name can't be empty")
)
