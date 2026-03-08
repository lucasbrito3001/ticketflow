package salewindow

import "errors"

var (
	ErrInvalidSaleWindowPeriod = errors.New("The closing time must be after the opening time")
)
