package domain

import "time"

type SaleWindow struct {
	opensAt  time.Time
	closesAt time.Time
}

func New(opensAt, closesAt time.Time) (*SaleWindow, error) {
	if closesAt.Before(opensAt) {
		return nil, ErrInvalidSaleWindowPeriod
	}

	return &SaleWindow{
		opensAt:  opensAt,
		closesAt: closesAt,
	}, nil
}

func Rehydrate(opensAt, closesAt time.Time) *SaleWindow {
	return &SaleWindow{
		opensAt:  opensAt,
		closesAt: closesAt,
	}
}

func (sw SaleWindow) OpensAt() time.Time {
	return sw.opensAt
}

func (sw SaleWindow) ClosesAt() time.Time {
	return sw.closesAt
}

func (sw SaleWindow) IsOpen() bool {
	return !time.Now().Before(sw.opensAt) && !time.Now().After(sw.closesAt)
}
