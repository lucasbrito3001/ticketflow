package domain

import "errors"

var (
	ErrInvalidSaleWindowPeriod = errors.New("The closing time must be after the opening time")
)
