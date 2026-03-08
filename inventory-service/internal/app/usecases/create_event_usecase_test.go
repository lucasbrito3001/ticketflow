package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/app/usecases/dto"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/domain/event"
	"github.com/lucasbrito3001/ticketflow-inventory-service/mocks"
)

type mockUnitOfWork struct {
	doFunc func(context.Context, func(context.Context) error) error
}

func (m *mockUnitOfWork) Do(ctx context.Context, fn func(context.Context) error) error {
	if m.doFunc != nil {
		return m.doFunc(ctx, fn)
	}
	// Default: just execute the function without transaction
	return fn(ctx)
}

func validCreateEventInput() dto.CreateEventInput {
	return dto.CreateEventInput{
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
		CatalogItems: []dto.CreateTicketCatalogItem{
			{
				TicketType:    "GENERAL",
				Price:         50000,
				TotalQuantity: 100,
			},
			{
				TicketType:    "VIP",
				Price:         100000,
				TotalQuantity: 50,
			},
		},
	}
}

func TestCreateEvent_Execute(t *testing.T) {
	mockUoW := &mocks.UnitOfWork{}
	mockEventRepo := &mocks.EventRepository{}

	t.Run("should return error when fails to convert the input to domain", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		input.Name = ""

		useCase := NewCreateEvent(mockUoW, mockEventRepo)

		// When
		output, err := useCase.Execute(context.Background(), input)

		// Then
		assert.Equal(t, event.ErrEventNameEmpty, err)
		assert.Nil(t, output)
	})

	t.Run("should return error when fails to create the unit of work", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		useCase := NewCreateEvent(mockUoW, mockEventRepo)

		mockUoW.On(
			"Do",
			mock.Anything,
			mock.AnythingOfType("func(context.Context) error"),
		).Return(assert.AnError).Once()

		// When
		output, err := useCase.Execute(context.Background(), input)

		// Then
		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		assert.Nil(t, output)
	})

	t.Run("should return error when fails to save the event", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		useCase := NewCreateEvent(mockUoW, mockEventRepo)

		mockUoW.On(
			"Do",
			mock.Anything,
			mock.Anything,
		).Return(func(ctx context.Context, fn func(context.Context) error) error {
			return fn(ctx)
		}).Once()
		mockEventRepo.On(
			"CreateEvent",
			mock.Anything,
			mock.AnythingOfType("*event.Event"),
		).Return(assert.AnError).Once()

		// When
		output, err := useCase.Execute(context.Background(), input)

		// Then
		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		assert.Nil(t, output)
	})

	t.Run("should return error when fails to save the ticket catalog items", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		useCase := NewCreateEvent(mockUoW, mockEventRepo)

		mockUoW.On(
			"Do",
			mock.Anything,
			mock.Anything,
		).Return(func(ctx context.Context, fn func(context.Context) error) error {
			return fn(ctx)
		}).Once()
		mockEventRepo.On(
			"CreateEvent",
			mock.Anything,
			mock.AnythingOfType("*event.Event"),
		).Return(nil).Once()
		mockEventRepo.On(
			"CreateEventTicketCatalogItems",
			mock.Anything,
			mock.AnythingOfType("*event.Event"),
		).Return(assert.AnError).Once()

		// When
		output, err := useCase.Execute(context.Background(), input)

		// Then
		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		assert.Nil(t, output)
	})

	t.Run("should create the event successfully", func(t *testing.T) {
		// Given
		input := validCreateEventInput()
		mockUow := &mockUnitOfWork{
			doFunc: func(ctx context.Context, fn func(context.Context) error) error {
				return fn(ctx)
			},
		}
		useCase := NewCreateEvent(mockUow, mockEventRepo)

		mockUoW.On(
			"Do",
			mock.Anything,
			mock.Anything,
		).Return(func(ctx context.Context, fn func(context.Context) error) error {
			return fn(ctx)
		}).Once()
		mockEventRepo.On(
			"CreateEvent",
			mock.AnythingOfType("context.backgroundCtx"),
			mock.AnythingOfType("*event.Event"),
		).Return(nil).Once()
		mockEventRepo.On(
			"CreateEventTicketCatalogItems",
			mock.AnythingOfType("context.backgroundCtx"),
			mock.AnythingOfType("*event.Event"),
		).Return(nil).Once()

		// When
		output, err := useCase.Execute(context.Background(), input)

		// Then
		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, "Event created successfully", output.Message)
	})
}
