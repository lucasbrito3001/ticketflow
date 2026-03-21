package reservation

import "strings"

type ReservationStatus string

const (
	ReservationStatusCreated         ReservationStatus = "CREATED"
	ReservationStatusCancelled       ReservationStatus = "CANCELLED"
	ReservationStatusAwaitingPayment ReservationStatus = "AWAITING_PAYMENT"
	ReservationStatusReserved        ReservationStatus = "RESERVED"
	ReservationStatusExpired         ReservationStatus = "EXPIRED"
)

var validReservationStatuses map[ReservationStatus]struct{} = map[ReservationStatus]struct{}{
	ReservationStatusCreated:         {},
	ReservationStatusCancelled:       {},
	ReservationStatusAwaitingPayment: {},
	ReservationStatusReserved:        {},
	ReservationStatusExpired:         {},
}

func RehydrateReservationStatus(name string) ReservationStatus {
	return ReservationStatus(name)
}

func NewReservationStatus(name string) (ReservationStatus, error) {
	status := ReservationStatus(name)

	if _, ok := validReservationStatuses[status]; !ok {
		return "", ErrInvalidReservationStatus
	}

	return status, nil
}

func (rs ReservationStatus) String() string {
	return string(rs)
}

func ValidReservationStatusesString() string {
	var statuses strings.Builder
	for status := range validReservationStatuses {
		statuses.WriteString(status.String() + ", ")
	}

	return statuses.String()
}
