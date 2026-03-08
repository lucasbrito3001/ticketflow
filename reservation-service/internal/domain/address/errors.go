package address

import "errors"

var (
	ErrStreetEmpty  = errors.New("The street can't be empty")
	ErrNumberEmpty  = errors.New("The number can't be empty")
	ErrCityEmpty    = errors.New("The city can't be empty")
	ErrStateEmpty   = errors.New("The state can't be empty")
	ErrZipCodeEmpty = errors.New("The zip code can't be empty")
)
