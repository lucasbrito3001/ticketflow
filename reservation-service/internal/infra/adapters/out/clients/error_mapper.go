package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/reservation"
)

type InventoryErrorResponse struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

func convertClientErrorToDomainError(resp *http.Response) error {
	defer resp.Body.Close()

	var errResp InventoryErrorResponse
	json.NewDecoder(resp.Body).Decode(&errResp)

	fmt.Println(errResp)

	switch errResp.ErrorCode {
	case "EVENT_NOT_FOUND":
		return reservation.ErrEventNotFound
	case "INSUFFICIENT_TICKETS":
		return reservation.ErrInsufficientTickets
	}

	return reservation.ErrUnexpected
}
