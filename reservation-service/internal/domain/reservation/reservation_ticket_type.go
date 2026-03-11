package reservation

const (
	ReservationTicketTypeGeneral ReservationTicketType = "GENERAL"
	ReservationTicketTypeVIP     ReservationTicketType = "VIP"
	ReservationTicketTypeBoxSeat ReservationTicketType = "BOX_SEAT"
)

type ReservationTicketType string

var validTicketTypes = map[ReservationTicketType]struct{}{
	ReservationTicketTypeGeneral: {},
	ReservationTicketTypeVIP:     {},
	ReservationTicketTypeBoxSeat: {},
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
