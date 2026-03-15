package usecases

import (
	"context"
	"errors"
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

		mockReservationRepo.On(
			"Create",
			mock.Anything,
			mock.AnythingOfType("*reservation.Reservation"),
		).Return(int64(1), nil).Once()
		mockInventoryClient.On(
			"HoldTickets",
			mock.Anything,
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("map[reservation.ReservationTicketType]int"),
		).Return(nil).Once()
		mockReservationRepo.On(
			"Update",
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
		).Return(int64(0), assert.AnError).Once()

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
		).Return(int64(1), nil).Once()
		mockInventoryClient.On(
			"HoldTickets",
			mock.Anything,
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("map[reservation.ReservationTicketType]int"),
		).Return(assert.AnError).Once()
		mockReservationRepo.On(
			"Update",
			mock.Anything,
			mock.AnythingOfType("*reservation.Reservation"),
		).Return(nil).Once()

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
		).Return(int64(1), nil).Once()
		mockInventoryClient.On(
			"HoldTickets",
			mock.Anything,
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("map[reservation.ReservationTicketType]int"),
		).Return(errors.New("error hold")).Once()
		mockReservationRepo.On(
			"Update",
			mock.Anything,
			mock.AnythingOfType("*reservation.Reservation"),
		).Return(errors.New("error update")).Once()

		// When
		output, err := useCase.Execute(context.Background(), input)

		// Then
		assert.Error(t, err)
		assert.Equal(t, "error update", err.Error())
		assert.Nil(t, output)
	})
}
