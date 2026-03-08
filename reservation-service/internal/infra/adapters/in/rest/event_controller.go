package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/usecases"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/usecases/dto"
)

type EventController struct {
	reserveTicketUseCase usecases.ReserveTicket
}

func NewEventController(reserveTicketUseCase usecases.ReserveTicket) *EventController {
	return &EventController{
		reserveTicketUseCase: reserveTicketUseCase,
	}
}

func (c *EventController) ReserveTicket(ctx *gin.Context) {
	input := &dto.ReserveTicketInput{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	output, err := c.reserveTicketUseCase.Execute(ctx.Request.Context(), *input)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Tickets reserved successfully", "data": output})
}
