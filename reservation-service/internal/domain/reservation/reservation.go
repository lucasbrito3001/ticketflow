package reservation

type Reservation struct {
	id      int64
	eventID int64
	tickets map[ReservationTicketType]int
	status  ReservationStatus
}

func RehydrateReservation(id int64, eventID int64, tickets map[ReservationTicketType]int, status ReservationStatus) *Reservation {
	return &Reservation{
		id:      id,
		eventID: eventID,
		tickets: tickets,
		status:  status,
	}
}

func NewCreatedReservation(eventID int64, tickets map[ReservationTicketType]int) (*Reservation, error) {
	return NewReservation(eventID, tickets, ReservationStatusCreated)
}

func NewReservation(eventID int64, tickets map[ReservationTicketType]int, status ReservationStatus) (*Reservation, error) {
	if eventID <= 0 {
		return nil, ErrInvalidEventID
	}

	if len(tickets) == 0 {
		return nil, ErrEmptyTickets
	}

	for _, quantity := range tickets {
		if quantity <= 0 {
			return nil, ErrInvalidTicketQuantity
		}
	}

	return &Reservation{
		eventID: eventID,
		tickets: tickets,
		status:  status,
	}, nil
}

func (r *Reservation) Id() int64 {
	return r.id
}

func (r *Reservation) EventID() int64 {
	return r.eventID
}

func (r *Reservation) Tickets() map[ReservationTicketType]int {
	return r.tickets
}

func (r *Reservation) Status() ReservationStatus {
	return r.status
}

func (r *Reservation) SetAsCancelled() error {
	if r.status == ReservationStatusReserved || r.status == ReservationStatusExpired {
		return ErrCannotSetAsCancelled
	}

	r.status = ReservationStatusCancelled

	return nil
}

func (r *Reservation) SetAsAwaitingPayment() error {
	if r.status != ReservationStatusCreated {
		return ErrCannotSetAsAwaitingPayment
	}

	r.status = ReservationStatusAwaitingPayment

	return nil
}
