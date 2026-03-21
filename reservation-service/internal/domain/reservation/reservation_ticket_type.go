package reservation

import "strings"

const (
	ReservationTicketTypeGeneral ReservationTicketType = "GENERAL"
	ReservationTicketTypeVIP     ReservationTicketType = "VIP"
	ReservationTicketTypeBoxSeat ReservationTicketType = "BOX_SEAT"
)

type (
	ReservationTicketType string

	AllowedReservationTicketTypes map[ReservationTicketType]struct{}
)

var validTicketTypes = AllowedReservationTicketTypes{
	ReservationTicketTypeGeneral: {},
	ReservationTicketTypeVIP:     {},
	ReservationTicketTypeBoxSeat: {},
}

func RehydrateReservationTicketType(name string) ReservationTicketType {
	return ReservationTicketType(name)
}

func NewReservationTicketType(name string) (ReservationTicketType, error) {
	ticketType := ReservationTicketType(name)

	if _, ok := validTicketTypes[ticketType]; !ok {
		return "", ErrInvalidTicketType
	}

	return ticketType, nil
}

func (tt ReservationTicketType) String() string {
	return string(tt)
}

func (a AllowedReservationTicketTypes) IsAllowed(ticketType ReservationTicketType) bool {
	_, ok := a[ticketType]
	return ok
}

func ValidTicketTypesString() string {
	var types strings.Builder
	for ticketType := range validTicketTypes {
		types.WriteString(ticketType.String() + ", ")
	}

	return types.String()
}
