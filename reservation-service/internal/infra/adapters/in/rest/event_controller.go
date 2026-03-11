package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/usecases"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/usecases/dto"
)

type EventController struct {
	createReservationUseCase usecases.CreateReservationUseCase
}

func NewEventController(createReservationUseCase usecases.CreateReservationUseCase) *EventController {
	return &EventController{
		createReservationUseCase: createReservationUseCase,
	}
}

func (c *EventController) ReserveTicket(ctx *gin.Context) {
	input := &dto.CreateReservationInput{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	output, err := c.createReservationUseCase.Execute(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Tickets reserved successfully", "data": output})
}
