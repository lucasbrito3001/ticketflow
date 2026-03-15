package reservation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReservation(t *testing.T) {
	t.Run("should create a reservation with valid input", func(t *testing.T) {
		// Given
		eventId := int64(1)
		tickets := map[ReservationTicketType]int{
			ReservationTicketTypeVIP: 2,
		}

		// When
		reservation, err := NewReservation(eventId, tickets, ReservationStatusCreated)

		// Then
		assert.NoError(t, err)
		assert.NotNil(t, reservation)
		assert.Equal(t, eventId, reservation.EventId())
		assert.Equal(t, tickets, reservation.Tickets())
	})

	t.Run("should fail with empty event ID", func(t *testing.T) {
		// Given
		eventId := int64(0)
		tickets := map[ReservationTicketType]int{
			ReservationTicketTypeVIP: 2,
		}

		// When
		reservation, err := NewReservation(eventId, tickets, ReservationStatusCreated)

		// Then
		assert.ErrorIs(t, err, ErrInvalidEventID)
		assert.Nil(t, reservation)
	})

	t.Run("should fail with empty tickets", func(t *testing.T) {
		// Given
		eventId := int64(1)
		tickets := map[ReservationTicketType]int{}

		// When
		reservation, err := NewReservation(eventId, tickets, ReservationStatusCreated)

		// Then
		assert.ErrorIs(t, err, ErrEmptyTickets)
		assert.Nil(t, reservation)
	})

	t.Run("should fail with invalid ticket quantity", func(t *testing.T) {
		// Given
		eventId := int64(1)
		tickets := map[ReservationTicketType]int{
			ReservationTicketTypeVIP: 0,
		}

		// When
		reservation, err := NewReservation(eventId, tickets, ReservationStatusCreated)

		// Then
		assert.ErrorIs(t, err, ErrInvalidTicketQuantity)
		assert.Nil(t, reservation)
	})
}
