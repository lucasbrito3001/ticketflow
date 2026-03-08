package dto

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/event"
)

func validCreateEventInput() *CreateEventInput {
	return &CreateEventInput{
		Name:         "Test Event",
		Venue:        "Test Venue",
		StartsAt:     time.Date(2024, 9, 15, 20, 0, 0, 0, time.UTC),
		EndsAt:       time.Date(2024, 9, 15, 23, 0, 0, 0, time.UTC),
		OpenSalesAt:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		CloseSalesAt: time.Date(2024, 9, 14, 23, 59, 59, 0, time.UTC),
		Street:       "Test Street",
		Number:       123,
		City:         "Test City",
		State:        "TS",
		ZipCode:      "12345678",
		CatalogItems: []CreateTicketCatalogItem{
			{
				TicketType:    "GENERAL",
				Price:         50000,
				TotalQuantity: 100,
			},
		},
	}
}

func TestCreateEventInput_ToDomain(t *testing.T) {
	t.Run("should successfully convert valid input to domain with multiple catalog items", func(t *testing.T) {
		// Given
		input := &CreateEventInput{
			Name:         "Test Event",
			Venue:        "Test Venue",
			StartsAt:     time.Date(2024, 9, 15, 20, 0, 0, 0, time.UTC),
			EndsAt:       time.Date(2024, 9, 15, 23, 0, 0, 0, time.UTC),
			OpenSalesAt:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			CloseSalesAt: time.Date(2024, 9, 14, 23, 59, 59, 0, time.UTC),
			Street:       "Test Street",
			Number:       123,
			City:         "Test City",
			State:        "TS",
			ZipCode:      "12345678",
			CatalogItems: []CreateTicketCatalogItem{
				{TicketType: "GENERAL", Price: 50000, TotalQuantity: 100},
				{TicketType: "VIP", Price: 100000, TotalQuantity: 50},
				{TicketType: "BOX_SEAT", Price: 150000, TotalQuantity: 20},
			},
		}

		// When
		result, err := input.ToDomain()

		// Then
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "Test Event", result.Name())
		assert.NotNil(t, result.TicketCatalog())
	})

	t.Run("should fail with empty name", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		input.Name = ""

		// When
		result, err := input.ToDomain()

		// Then
		assert.Error(t, err)
		assert.Equal(t, event.ErrEventNameEmpty, err)
		assert.Nil(t, result)
	})

	t.Run("should fail with empty venue", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		input.Venue = ""

		// When
		result, err := input.ToDomain()

		// Then
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should fail when event ends before starts", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		input.EndsAt = input.StartsAt.Add(-1 * time.Hour)

		// When
		result, err := input.ToDomain()

		// Then
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should fail with empty street", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		input.Street = ""

		// When
		result, err := input.ToDomain()

		// Then
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should fail with invalid address number", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		input.Number = 0

		// When
		result, err := input.ToDomain()

		// Then
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should fail with empty city", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		input.City = ""

		// When
		result, err := input.ToDomain()

		// Then
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should fail with empty state", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		input.State = ""

		// When
		result, err := input.ToDomain()

		// Then
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should fail with invalid zip code", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		input.ZipCode = "123"

		// When
		result, err := input.ToDomain()

		// Then
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should fail with invalid sale window period", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		input.CloseSalesAt = input.OpenSalesAt.Add(-1 * time.Hour)

		// When
		result, err := input.ToDomain()

		// Then
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should fail with invalid ticket type", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		input.CatalogItems = []CreateTicketCatalogItem{
			{TicketType: "INVALID_TYPE_XYZ", Price: 50000, TotalQuantity: 100},
		}

		// When
		result, err := input.ToDomain()

		// Then
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should fail with invalid money amount", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		input.CatalogItems = []CreateTicketCatalogItem{
			{TicketType: "GENERAL", Price: -5000, TotalQuantity: 100},
		}

		// When
		result, err := input.ToDomain()

		// Then
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should fail with invalid ticket quantity", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		input.CatalogItems = []CreateTicketCatalogItem{
			{TicketType: "GENERAL", Price: 50000, TotalQuantity: 0},
		}

		// When
		result, err := input.ToDomain()

		// Then
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
