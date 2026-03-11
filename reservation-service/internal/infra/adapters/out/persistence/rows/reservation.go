package rows

import "github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/reservation"

type ReservationRow struct {
	ID      int64
	EventID int64
	Status  string
}

func (r *ReservationRow) ToDomain() *Reservation {
	status := reservation.RehydrateReservationStatus(r.Status)
	return reservation.RehydrateReservation(
		r.ID,
		r.EventID,
		map[reservation.ReservationTicketType]int{},
		status,
	)
}
