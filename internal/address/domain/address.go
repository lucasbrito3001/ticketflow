package domain

import (
	"strconv"
)

type Address struct {
	street  string
	number  int
	city    string
	state   string
	zipCode string
}

func New(street string, number int, city string, state string, zipCode string) (*Address, error) {
	if street == "" {
		return nil, ErrStreetEmpty
	}

	if number < 1 {
		return nil, ErrNumberEmpty
	}

	if city == "" {
		return nil, ErrCityEmpty
	}

	if state == "" {
		return nil, ErrStateEmpty
	}

	if !isZipCodeValid(zipCode) {
		return nil, ErrZipCodeEmpty
	}

	return &Address{
		street:  street,
		number:  number,
		city:    city,
		state:   state,
		zipCode: zipCode,
	}, nil
}

func Rehydrate(street string, number int, city string, state string, zipCode string) *Address {
	return &Address{
		street:  street,
		number:  number,
		city:    city,
		state:   state,
		zipCode: zipCode,
	}
}

func isZipCodeValid(zipCode string) bool {
	if len(zipCode) != 8 {
		return false
	}

	return true
}

func (a *Address) Street() string {
	return a.street
}

func (a *Address) Number() int {
	return a.number
}

func (a *Address) City() string {
	return a.city
}

func (a *Address) State() string {
	return a.state
}

func (a *Address) ZipCode() string {
	return a.zipCode
}

func (a *Address) FullAddress() string {
	return a.street + ", " + strconv.Itoa(a.number) + ", " + a.city + ", " + a.state + " " + a.zipCode
}
