package usecases

import (
	"context"
	"testing"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/usecases/dto"
	"github.com/lucasbrito3001/ticketflow-reservation-service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func validCreateReservationInput() *dto.CreateReservationInput {
	return &dto.CreateReservationInput{
		EventID: 1,
		Tickets: map[string]int{
			"GENERAL": 2,
			"VIP":     1,
		},
	}
}

func TestCreateReservation_Execute(t *testing.T) {
	mockInventoryClient := new(mocks.InventoryServiceClient)
	mockReservationRepo := new(mocks.ReservationRepository)

	t.Run("should create a reservation successfully", func(t *testing.T) {
		// Given
		input := validCreateReservationInput()
		useCase := NewCreateReservationUseCase(mockReservationRepo, mockInventoryClient)

		mockInventoryClient.On(
			"HoldTickets",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(nil).Once()
		mockReservationRepo.On(
			"Create",
			mock.Anything,
			mock.AnythingOfType("*reservation.Reservation"),
		).Return(nil).Once()

		// When
		output, err := useCase.Execute(context.Background(), input)

		// Then
		assert.NoError(t, err)
		assert.NotNil(t, output)
	})

	t.Run("should return error when fails to convert input into domain", func(t *testing.T) {
		// Given
		input := validCreateReservationInput()
		input.EventID = 0
		useCase := NewCreateReservationUseCase(mockReservationRepo, mockInventoryClient)

		// When
		output, err := useCase.Execute(context.Background(), input)

		// Then
		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("should return error when fails to create the reservation on database", func(t *testing.T) {
		// Given
		input := validCreateReservationInput()
		useCase := NewCreateReservationUseCase(mockReservationRepo, mockInventoryClient)

		mockReservationRepo.On(
			"Create",
			mock.Anything,
			mock.AnythingOfType("*reservation.Reservation"),
		).Return(assert.AnError).Once()

		// When
		output, err := useCase.Execute(context.Background(), input)

		// Then
		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		assert.Nil(t, output)
	})

	t.Run("should return error when fails to hold tickets on inventory", func(t *testing.T) {
		// Given
		input := validCreateReservationInput()
		useCase := NewCreateReservationUseCase(mockReservationRepo, mockInventoryClient)

		mockReservationRepo.On(
			"Create",
			mock.Anything,
			mock.AnythingOfType("*reservation.Reservation"),
		).Return(nil).Once()
		mockInventoryClient.On(
			"HoldTickets",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(assert.AnError).Once()

		// When
		output, err := useCase.Execute(context.Background(), input)

		// Then
		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		assert.Nil(t, output)
	})

	t.Run("should return error when fails to update the reservation status after hold tickets", func(t *testing.T) {
		// Given
		input := validCreateReservationInput()
		useCase := NewCreateReservationUseCase(mockReservationRepo, mockInventoryClient)

		mockReservationRepo.On(
			"Create",
			mock.Anything,
			mock.AnythingOfType("*reservation.Reservation"),
		).Return(nil).Once()
		mockInventoryClient.On(
			"HoldTickets",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(nil).Once()
		mockReservationRepo.On(
			"Update",
			mock.Anything,
			mock.AnythingOfType("*reservation.Reservation"),
		).Return(assert.AnError).Once()

		// When
		output, err := useCase.Execute(context.Background(), input)

		// Then
		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		assert.Nil(t, output)
	})
}
