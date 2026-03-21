package reservation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReservationTicketType(t *testing.T) {
	t.Run("should create a reservation ticket type with valid input", func(t *testing.T) {
		// Given
		name := "VIP"

		// When
		ticketType, err := NewReservationTicketType(name)

		// Then
		assert.NoError(t, err)
		assert.NotNil(t, ticketType)
		assert.Equal(t, name, ticketType.String())
	})

	t.Run("should fail with empty name", func(t *testing.T) {
		// Given
		name := ""

		// When
		ticketType, err := NewReservationTicketType(name)

		// Then
		assert.ErrorIs(t, err, ErrInvalidTicketType)
		assert.Equal(t, ReservationTicketType(""), ticketType)

	})

	t.Run("should fail with invalid name", func(t *testing.T) {
		// Given
		name := "INVALID"

		// When
		ticketType, err := NewReservationTicketType(name)

		// Then
		assert.ErrorIs(t, err, ErrInvalidTicketType)
		assert.Equal(t, ReservationTicketType(""), ticketType)
	})
}
