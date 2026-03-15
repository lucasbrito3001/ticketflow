package records

import "github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/reservation"

type (
	ReservationWithTicketsRecords []*ReservationWithTicketsRecord

	ReservationWithTicketsRecord struct {
		// Reservation
		Id      int64
		EventId int64
		Status  string

		// Reservation Tickets
		TicketType string
		Quantity   int
	}
)

func (r ReservationWithTicketsRecords) ToDomain() *reservation.Reservation {
	tickets := make(map[reservation.ReservationTicketType]int)
	for _, row := range r {
		tickets[reservation.RehydrateReservationTicketType(row.TicketType)] = row.Quantity
	}

	return reservation.RehydrateReservation(r[0].Id, r[0].EventId, tickets, reservation.ReservationStatus(r[0].Status))
}
