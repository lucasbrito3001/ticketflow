package statements

func InsertEvent() string {
	return `
		INSERT INTO events (
			name,
			starts_at,
			ends_at,
			venue,
			street,
			number,
			city,
			state,
			zip_code,
			open_sales_at,
			close_sales_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
}

func InsertTicketCatalogItem() string {
	return `
		INSERT INTO ticket_catalog_items (
			event_id,
			type,
			price,
			total_quantity,
			sold_quantity
		) VALUES (?, ?, ?, ?, ?)
	`
}
