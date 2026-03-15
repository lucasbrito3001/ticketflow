package statements

func InsertReservation() string {
	return `
		INSERT INTO reservations (
			id,
			event_id,
			status,
			created_at,
			updated_at
		) VALUES (?, ?, ?, ?, ?)
	`
}

func InsertReservationTickets() string {
	return `
		INSERT INTO reservation_tickets (
			reservation_id,
			ticket_type,
			quantity
		) VALUES (?, ?, ?)
	`
}
