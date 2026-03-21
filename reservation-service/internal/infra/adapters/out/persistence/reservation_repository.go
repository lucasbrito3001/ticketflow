package persistence

import (
	"context"
	"database/sql"
	"time"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/ports"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/reservation"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/infra/adapters/out/persistence/records"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/infra/adapters/out/persistence/statements"
)

type reservationRepository struct {
	db *sql.DB
}

func NewReservationRepository(db *sql.DB) ports.ReservationRepository {
	return &reservationRepository{
		db: db,
	}
}

func (r *reservationRepository) FindById(ctx context.Context, ticketID string) (*reservation.Reservation, error) {
	query := statements.SelectReservationWithTicktesById()

	rows, err := r.db.QueryContext(ctx, query, ticketID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservationRecords records.ReservationWithTicketsRecords
	for rows.Next() {
		var record records.ReservationWithTicketsRecord
		if err := rows.Scan(&record.Id, &record.EventId, &record.Status, &record.TicketType, &record.Quantity); err != nil {
			return nil, err
		}
		reservationRecords = append(reservationRecords, &record)
	}

	return reservationRecords.ToDomain(), nil
}

func (r *reservationRepository) Create(ctx context.Context, reservation *reservation.Reservation) (int64, error) {
	query := statements.InsertReservation()
	now := time.Now().UTC()

	result, err := r.db.ExecContext(ctx, query, reservation.Id(), reservation.EventId(), reservation.Status(), now, now)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *reservationRepository) Update(ctx context.Context, reservation *reservation.Reservation) error {
	query := statements.UpdateReservation()
	now := time.Now().UTC()

	_, err := r.db.ExecContext(ctx, query, reservation.Id(), reservation.EventId(), reservation.Status(), now, reservation.Id())

	return err
}
