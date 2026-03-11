package dto

import "github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/reservation"

type CreateReservationInput struct {
	EventID int64          `json:"event_id"`
	Tickets map[string]int `json:"tickets"`
}

type CreateReservationOutput struct {
	ReservationID int64 `json:"reservation_id"`
}

func (i *CreateReservationInput) ToDomain() (*reservation.Reservation, error) {
	tickets := make(map[reservation.ReservationTicketType]int)

	for ticketType, quantity := range i.Tickets {
		ticketType, err := reservation.NewReservationTicketType(ticketType)
		if err != nil {
			return nil, err
		}

		tickets[ticketType] = quantity
	}

	reservation, err := reservation.NewCreatedReservation(i.EventID, tickets)

	if err != nil {
		return nil, err
	}

	return reservation, nil
}
