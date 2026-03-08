package persistence

import (
	"context"
	"database/sql"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/ports"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/event"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/infra/adapters/out"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/infra/adapters/out/persistence/rows"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/infra/adapters/out/persistence/statements"
)

type mySqlEventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) ports.EventRepository {
	return &mySqlEventRepository{db: db}
}

func (r *mySqlEventRepository) FindById(
	ctx context.Context,
	id int64,
) (*event.Event, error) {
	query := statements.SelectEventByIdWithCatalogItem()

	resultRows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer resultRows.Close()

	catalogRows := make(rows.EventWithTicketCatalogItemRows, 0)
	for resultRows.Next() {
		row, err := r.scanEventWithTicketCatalogItemRow(resultRows)
		if err != nil {
			return nil, err
		}

		catalogRows = append(catalogRows, &row)
	}

	if err := resultRows.Err(); err != nil {
		return nil, err
	}

	domainEvent := catalogRows.ToDomain()
	if domainEvent == nil {
		return nil, ErrEventNotFound
	}

	return domainEvent, nil
}

func (r *mySqlEventRepository) Create(ctx context.Context, event *event.Event) error {
	exec := out.GetQueryExecutor(ctx, r.db)

	if err := r.insertEvent(ctx, exec, event); err != nil {
		return err
	}

	return r.insertTicketCatalog(ctx, exec, event)
}

func (r *mySqlEventRepository) Update(ctx context.Context, event *event.Event) error {
	exec := out.GetQueryExecutor(ctx, r.db)

	if err := r.updateEvent(ctx, exec, event); err != nil {
		return err
	}

	return r.updateTicketCatalog(ctx, exec, event)
}

func (r *mySqlEventRepository) scanEventWithTicketCatalogItemRow(scanner interface{ Scan(...interface{}) error }) (rows.EventWithTicketCatalogItemRow, error) {
	var row rows.EventWithTicketCatalogItemRow
	err := scanner.Scan(
		&row.EventID, &row.Name, &row.StartsAt, &row.EndsAt,
		&row.VenueName, &row.TicketType, &row.Total, &row.Sold,
		&row.Street, &row.Number, &row.City, &row.State, &row.ZipCode,
	)
	return row, err
}

func (r *mySqlEventRepository) insertEvent(
	ctx context.Context,
	exec out.QueryExecutor,
	event *event.Event,
) error {
	_, err := exec.ExecContext(ctx, statements.InsertEvent(),
		event.Name(),
		event.StartsAt(),
		event.EndsAt(),
		event.Venue(),
		event.Address().Street(),
		event.Address().Number(),
		event.Address().City(),
		event.Address().State(),
		event.Address().ZipCode(),
		event.SaleWindow().OpensAt(),
		event.SaleWindow().ClosesAt(),
	)
	return err
}

func (r *mySqlEventRepository) insertTicketCatalog(
	ctx context.Context,
	exec out.QueryExecutor,
	event *event.Event,
) error {
	for _, item := range *event.TicketCatalog() {
		_, err := exec.ExecContext(ctx, statements.InsertTicketCatalogItem(),
			event.ID(),
			item.TicketType(),
			item.Price(),
			item.TotalQuantity(),
			item.SoldQuantity(),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *mySqlEventRepository) updateEvent(
	ctx context.Context,
	exec out.QueryExecutor,
	event *event.Event,
) error {
	_, err := exec.ExecContext(ctx, statements.UpdateEvent(),
		event.Name(),
		event.StartsAt(),
		event.EndsAt(),
		event.Venue(),
		event.Address().Street(),
		event.Address().Number(),
		event.Address().City(),
		event.Address().State(),
		event.Address().ZipCode(),
		event.SaleWindow().OpensAt(),
		event.SaleWindow().ClosesAt(),
		event.ID(),
	)
	return err
}

func (r *mySqlEventRepository) updateTicketCatalog(
	ctx context.Context,
	exec out.QueryExecutor,
	event *event.Event,
) error {
	for _, item := range *event.TicketCatalog() {
		_, err := exec.ExecContext(ctx, statements.UpdateTicketCatalogItem(),
			item.TicketType(),
			item.Price(),
			item.TotalQuantity(),
			item.SoldQuantity(),
			event.ID(),
			item.TicketType(),
		)
		if err != nil {
			return err
		}
	}
	return nil
}
