package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/usecases/dto"
	"github.com/lucasbrito3001/ticketflow-reservation-service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReservationController_ReserveTickets(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := &mocks.CreateReservationUseCase{}
	controller := NewReservationController(mockUseCase)
	router := gin.Default()
	router.POST("/reserve", controller.ReserveTickets)

	t.Run("should reserve tickets successfully", func(t *testing.T) {
		mockUseCase.On("Execute", mock.Anything, mock.AnythingOfType("*dto.CreateReservationInput")).Return(&dto.CreateReservationOutput{
			ReservationID: 123,
		}, nil)

		input := &dto.CreateReservationInput{
			EventID: 1,
			Tickets: map[string]int{
				"VIP": 2,
			},
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/reserve", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		var responseBody map[string]interface{}
		json.Unmarshal(resp.Body.Bytes(), &responseBody)
		assert.Equal(t, "Tickets reserved successfully", responseBody["message"])
		data := responseBody["data"].(map[string]interface{})
		assert.Equal(t, float64(123), data["reservation_id"])
		mockUseCase.AssertExpectations(t)
	})

	t.Run("should return error when use case fails", func(t *testing.T) {
		mockUseCase.On("Execute", mock.Anything, mock.AnythingOfType("*dto.CreateReservationInput")).Return(nil, assert.AnError)

		input := &dto.CreateReservationInput{
			EventID: 1,
			Tickets: map[string]int{
				"VIP": 2,
			},
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/reserve", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		var responseBody map[string]interface{}
		json.Unmarshal(resp.Body.Bytes(), &responseBody)
		assert.Equal(t, assert.AnError.Error(), responseBody["error"])
		mockUseCase.AssertExpectations(t)
	})

	t.Run("should return error when JSON is invalid", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/reserve", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		var responseBody map[string]interface{}
		json.Unmarshal(resp.Body.Bytes(), &responseBody)
		assert.Contains(t, responseBody["error"], "invalid character")
	})
}
