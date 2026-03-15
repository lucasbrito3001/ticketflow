package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lucasbrito3001/go-kit/observability/httpctx"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/ports"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/domain/reservation"
)

type (
	ConsumeTicketRequest struct {
		EventID int64                       `json:"event_id"`
		Tickets ConsumeTicketRequestTickets `json:"tickets"`
	}

	ConsumeTicketRequestTickets map[reservation.ReservationTicketType]int

	ConsumeTicketResponse struct {
		Message string `json:"message"`
	}

	inventoryServiceClient struct {
		baseURL string
		client  *http.Client
	}
)

func NewInventoryServiceClient(baseURL string) ports.InventoryServiceClient {
	client := httpctx.NewClient(nil)
	return &inventoryServiceClient{
		baseURL: baseURL,
		client:  client,
	}
}

func (c *inventoryServiceClient) HoldTickets(ctx context.Context, eventID int64, tickets map[reservation.ReservationTicketType]int) error {
	url := fmt.Sprintf("%s/events/%d/consume", c.baseURL, eventID)

	requestBody := ConsumeTicketRequest{
		EventID: eventID,
		Tickets: tickets,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("inventory service returned status code %d", resp.StatusCode)
	}

	return nil
}
