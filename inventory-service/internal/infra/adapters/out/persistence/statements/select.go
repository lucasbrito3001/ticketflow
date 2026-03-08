package statements

func SelectEventByIdWithCatalogItem() string {
	return `
		SELECT
			e.id, 
			e.name, 
			e.starts_at,
			e.ends_at,
			e.venue,
			a.street,
			a.number,
			a.city,
			a.state,
			a.zip_code,
			e.open_sales_at,
			e.close_sales_at,
			t.type AS ticket_type,
			t.total_quantity,
			t.sold_quantity
		FROM events e
		JOIN addresses a ON e.address_id = a.id
		JOIN ticket_catalog_items t ON e.id = t.event_id
		WHERE e.id = ?
	`
}
