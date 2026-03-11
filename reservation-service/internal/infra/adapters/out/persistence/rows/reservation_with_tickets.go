package rows

type (
	ReservationWithTicketsRows []*ReservationWithTicketsRow

	ReservationWithTicketsRow struct {
		// Reservation
		ID      int64
		EventID int64
		Status  string

		// Reservation Tickets
		TicketType string
		Quantity   int
	}
)
