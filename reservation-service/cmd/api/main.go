package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lucasbrito3001/go-kit/observability/httpctx"
	"github.com/lucasbrito3001/go-kit/observability/logger"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/usecases"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/infra/adapters/in/rest"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/infra/adapters/out"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/infra/adapters/out/persistence"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/infra/clients"
)

func main() {
	logger.Init("ticketflow-reservation-service")
	router := gin.Default()
	router.Use(httpctx.GinContextMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	dsn := "app:app@tcp(localhost:3306)/app"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// adapters
	unitOfWork := out.NewUnitOfWork(db)
	eventRepository := persistence.NewEventRepository(db)
	ticketRepository := persistence.NewTicketRepository(db)

	// clients
	inventoryClient := clients.NewInventoryServiceClient("http://localhost:8081")

	// use cases
	reserveTicketUseCase := usecases.NewReserveTicket(unitOfWork, eventRepository, ticketRepository, inventoryClient)

	// controllers
	reserveTicketController := rest.NewEventController(reserveTicketUseCase)

	router.POST("/events/:event_id/reserve", reserveTicketController.ReserveTicket)

	router.Run(":8080")
}
