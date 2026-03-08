package ticket

type TicketType string

const (
	TicketTypeGeneral TicketType = "GENERAL"
	TicketTypeVIP     TicketType = "VIP"
	TicketTypeBoxSeat TicketType = "BOX_SEAT"
)

func NewTicketType(name string) (TicketType, error) {
	if !IsValidTicketTypeName(name) {
		return "", ErrInvalidTicketType
	}

	return TicketType(name), nil
}

func IsValidTicketTypeName(name string) bool {
	if name == "" {
		return false
	}

	switch TicketType(name) {
	case TicketTypeGeneral, TicketTypeVIP, TicketTypeBoxSeat:
		return true
	default:
		return false
	}
}

func (t TicketType) String() string {
	return string(t)
}
