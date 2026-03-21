package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/app/usecases"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/app/usecases/dto"
)

type EventController struct {
	consumeTicketUseCase usecases.ConsumeTicket
	createEventUseCase   usecases.CreateEvent
}

func NewEventController(consumeTicketUseCase usecases.ConsumeTicket, createEventUseCase usecases.CreateEvent) *EventController {
	return &EventController{
		consumeTicketUseCase: consumeTicketUseCase,
		createEventUseCase:   createEventUseCase,
	}
}

func (c *EventController) ConsumeTicket(ctx *gin.Context) {
	input := &dto.ConsumeTicketInput{}

	if err := ctx.ShouldBindUri(input); err != nil {
		errResponse := MapError(err)
		ctx.JSON(errResponse.StatusCode, errResponse.Error)
		return
	}

	if err := ctx.ShouldBindJSON(input); err != nil {
		errResponse := MapError(err)
		ctx.JSON(errResponse.StatusCode, errResponse.Error)
		return
	}

	output, err := c.consumeTicketUseCase.Execute(ctx.Request.Context(), *input)
	if err != nil {
		errResponse := MapError(err)
		ctx.JSON(errResponse.StatusCode, errResponse.Error)
		return
	}

	ctx.JSON(200, gin.H{"message": output.Message})
}

// Admin functions

func (c *EventController) Create(ctx *gin.Context) {
	input := &dto.CreateEventInput{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		errResponse := MapError(err)
		ctx.JSON(errResponse.StatusCode, errResponse.Error)
		return
	}

	output, err := c.createEventUseCase.Execute(ctx.Request.Context(), *input)
	if err != nil {
		errResponse := MapError(err)
		ctx.JSON(errResponse.StatusCode, errResponse.Error)
		return
	}

	ctx.JSON(201, output)
}
