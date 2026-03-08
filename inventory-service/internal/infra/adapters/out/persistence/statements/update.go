package statements

func UpdateEvent() string {
	return `
		UPDATE events
		SET
			name = ?,
			starts_at = ?,
			ends_at = ?,
			venue = ?,
			street = ?,
			number = ?,
			city = ?,
			state = ?,
			zip_code = ?,
			open_sales_at = ?,
			close_sales_at = ?
		WHERE id = ?
	`
}

func UpdateTicketCatalogItem() string {
	return `
		UPDATE ticket_catalog_items
		SET
			type = ?,
			price = ?,
			total_quantity = ?,
			sold_quantity = ?
		WHERE event_id = ? AND type = ?
	`
}

func DecrementAvailableQuantity() string {
	return `
		UPDATE ticket_catalog_items
		SET sold_quantity = sold_quantity + ?
		WHERE 
			event_id = ? AND type = ?
			AND (sold_quantity + ?) <= total_quantity
	`
}
