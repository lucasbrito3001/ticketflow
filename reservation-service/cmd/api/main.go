package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lucasbrito3001/go-kit/observability/httpctx"
	"github.com/lucasbrito3001/go-kit/observability/logger"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/app/usecases"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/infra/adapters/in/rest"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/infra/adapters/out/clients"
	"github.com/lucasbrito3001/ticketflow-reservation-service/internal/infra/adapters/out/persistence"
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

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "app:app@tcp(localhost:3306)/ticketflow_reservation"
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// adapters
	reservationRepository := persistence.NewReservationRepository(db)

	// clients
	inventoryURL := os.Getenv("INVENTORY_URL")
	if inventoryURL == "" {
		inventoryURL = "http://localhost:8081"
	}
	inventoryClient := clients.NewInventoryServiceClient(inventoryURL)

	// use cases
	createReservationUseCase := usecases.NewCreateReservationUseCase(reservationRepository, inventoryClient)

	// controllers
	createReservationController := rest.NewReservationController(createReservationUseCase)

	router.POST("/events/reserve", createReservationController.ReserveTickets)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Match container EXPOSE / Healthcheck
	}
	router.Run(":" + port)
}
