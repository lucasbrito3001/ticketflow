package persistence

import (
	"context"
	"database/sql"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/ports"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/ticket"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/infra/adapters/out"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/infra/adapters/out/persistence/statements"
)

type mySqlTicketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) ports.TicketRepository {
	return &mySqlTicketRepository{db: db}
}

func (r *mySqlTicketRepository) FindById(
	ctx context.Context,
	ticketID string,
) (*ticket.Ticket, error) {
	// Implementation goes here
	return nil, nil
}

func (r *mySqlTicketRepository) Create(
	ctx context.Context,
	ticket *ticket.Ticket,
) error {
	// Implementation goes here
	return nil
}

func (r *mySqlTicketRepository) CreateBatch(
	ctx context.Context,
	tickets []*ticket.Ticket,
) error {
	exec := out.GetQueryExecutor(ctx, r.db)

	if len(tickets) == 0 {
		return nil
	}

	cols := 3
	values := make([]string, 0, len(tickets))
	args := make([]any, 0, len(tickets)*cols)

	for _, t := range tickets {
		values = append(values, "(?, ?, ?)")
		args = append(args,
			t.EventId(),
			t.Type(),
			t.Price(),
		)
	}

	query := statements.InsertTicketBatch(values)

	_, err := exec.ExecContext(ctx, query, args...)
	return err
}
