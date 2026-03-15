package rest

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/usecases"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/usecases/dto"
)

type ReservationController struct {
	createReservationUseCase usecases.CreateReservationUseCase
}

func NewReservationController(createReservationUseCase usecases.CreateReservationUseCase) *ReservationController {
	return &ReservationController{
		createReservationUseCase: createReservationUseCase,
	}
}

func (c *ReservationController) ReserveTickets(ctx *gin.Context) {
	input := &dto.CreateReservationInput{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "error binding JSON", "error", err)
		ctx.JSON(400, ErrorResponse{Error: "invalid request body", Details: err.Error()})
		return
	}

	output, err := c.createReservationUseCase.Execute(ctx.Request.Context(), input)
	if err != nil {
		mapperResult := MapError(err)
		ctx.JSON(mapperResult.StatusCode, mapperResult.Error)
		return
	}

	ctx.JSON(200, gin.H{"message": "Tickets reserved successfully", "data": output})
}
