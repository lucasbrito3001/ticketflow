package statements

var (
	SelectEventForUpdateByID = `
		SELECT
			e.id, 
			e.name, 
			e.starts_at,
			e.ends_at,
			e.venue_name,
			e.street,
			e.number,
			e.complement,
			e.city,
			e.state,
			e.zip_code,
			e.open_sales_at,
			e.close_sales_at,
			t.type AS ticket_type,
			t.total_quantity,
			t.sold_quantity
		FROM events e
		JOIN venues v ON e.venue_id = v.id
		JOIN addresses a ON v.address_id = a.id
		JOIN ticket_catalog_items t ON e.id = t.event_id
		WHERE e.id = ?
		FOR UPDATE
	`
)
