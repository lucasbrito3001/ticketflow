package statements

func SelectReservationWithTicktesById() string {
	return `
		SELECT
			r.id,
			r.event_id,
			r.status,
			rt.ticket_type,
			rt.quantity
		FROM reservations r
		INNER JOIN reservation_tickets rt ON r.id = rt.reservation_id
		WHERE r.id = ?
	`
}
