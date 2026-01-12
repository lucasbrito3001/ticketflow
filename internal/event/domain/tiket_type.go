package domain

type TicketType string

const (
	TicketTypeGeneral TicketType = "GENERAL"
	TicketTypeVIP     TicketType = "VIP"
	TicketTypeBoxSeat TicketType = "BOX_SEAT"
)
