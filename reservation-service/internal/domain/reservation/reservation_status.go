package reservation

type ReservationStatus string

const (
	ReservationStatusCreated         ReservationStatus = "CREATED"
	ReservationStatusCancelled       ReservationStatus = "CANCELLED"
	ReservationStatusAwaitingPayment ReservationStatus = "AWAITING_PAYMENT"
	ReservationStatusReserved        ReservationStatus = "RESERVED"
	ReservationStatusExpired         ReservationStatus = "EXPIRED"
)

var AllowedStatuses map[ReservationStatus]struct{} = map[ReservationStatus]struct{}{
	ReservationStatusCreated:         {},
	ReservationStatusCancelled:       {},
	ReservationStatusAwaitingPayment: {},
	ReservationStatusReserved:        {},
	ReservationStatusExpired:         {},
}

func RehydrateReservationStatus(name string) {
	return ReservationStatus(name)
}

func NewReservationStatus(name string) (ReservationStatus, error) {
	status := ReservationStatus(name)

	if _, ok := AllowedStatuses[status]; !ok {
		return "", ErrInvalidReservationStatus
	}

	return status, nil
}

func (rs ReservationStatus) String() string {
	return string(rs)
}
