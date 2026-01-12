package mysql

import (
	"context"
	"database/sql"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/event/adapters/out/mysql/rows"
	event "github.com/lucasbrito3001/ticketflow-reservation-service/internal/event/domain"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/event/usecases/ports"
)

type eventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) ports.EventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) FindById(ctx context.Context, eventID int) (*event.Event, error) {
	// Implementation goes here
	return nil, nil
}

func (r *eventRepository) FindByIDForUpdate(
	ctx context.Context,
	id int,
) (*event.Event, error) {
	query := "SELECT ... FOR UPDATE"

	sqlRows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer sqlRows.Close()

	var (
		e       *event.Event
		catalog = make([]*event.TicketCatalogItem, 0)
	)

	for sqlRows.Next() {
		row, err := r.scanEventWithTicketCatalogItemRow(sqlRows)
		if err != nil {
			return nil, err
		}

		if e == nil {
			e = row.ToDomain()
		}

		item := event.RehydrateTicketCatalogItem(event.TicketType(row.TicketType), row.Total, row.Sold)

		catalog = append(catalog, item)
	}

	if e == nil {
		return nil, ErrEventNotFound
	}

	e.AttachTicketCatalog(catalog)

	return e, nil
}

func (r *eventRepository) Create(ctx context.Context, event *event.Event) error {
	// Implementation goes here
	return nil
}

func (r *eventRepository) Update(ctx context.Context, event *event.Event) error {
	// Implementation goes here
	return nil
}

func (r *eventRepository) scanEventWithTicketCatalogItemRow(scanner interface{ Scan(...interface{}) error }) (rows.EventWithTicketCatalogItemRow, error) {
	var row rows.EventWithTicketCatalogItemRow
	err := scanner.Scan(
		&row.EventID, &row.Name, &row.StartsAt, &row.EndsAt,
		&row.VenueName, &row.TicketType, &row.Total, &row.Sold,
		&row.Street, &row.Number, &row.City, &row.State, &row.ZipCode,
	)
	return row, err
}
