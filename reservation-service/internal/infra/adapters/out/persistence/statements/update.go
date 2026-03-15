package statements

func UpdateReservation() string {
	return `
		UPDATE reservations
		SET
			id = ?,
			event_id = ?,
			status = ?,
			updated_at = ?
		WHERE id = ?
	`
}
