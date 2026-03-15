package main

import (
	"context"
	"database/sql"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lucasbrito3001/go-kit/observability/httpctx"
	"github.com/lucasbrito3001/go-kit/observability/logger"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/app/usecases"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/infra/adapters/in/rest"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/infra/adapters/out"
	"github.com/lucasbrito3001/ticketflow-inventory-service/internal/infra/adapters/out/persistence"
)

func main() {
	logger.Init("ticketflow-inventory-service")
	ctx := context.Background()

	// Middlewares
	router := gin.Default()
	router.Use(httpctx.GinContextMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Database
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "app:app@tcp(localhost:3306)/ticketflow_inventory?parseTime=true"
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		slog.Error("Failed to open database connection", "error", err, "dsn", dsn)
		log.Fatal(err)
	}
	defer db.Close()

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		slog.Error("Failed to ping database", "error", err)
		log.Fatal(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Adapters
	// UnitOfWork is used for coordinating atomic operations across multiple repositories
	unitOfWork := out.NewMySQLUnitOfWork(db)
	slog.Debug("UnitOfWork initialized for atomic operations")

	eventRepository := persistence.NewEventRepository(db)
	slog.Debug("EventRepository initialized")

	// Use cases
	consumeTicketUseCase := usecases.NewConsumeTicket(unitOfWork, eventRepository)
	slog.Debug("ConsumeTicket use case initialized")

	createEventUseCase := usecases.NewCreateEvent(unitOfWork, eventRepository)
	slog.Debug("CreateEvent use case initialized")

	// Controllers
	eventController := rest.NewEventController(consumeTicketUseCase, createEventUseCase)
	slog.Debug("EventController initialized")

	// Routes
	router.POST("/events", eventController.Create)
	slog.Debug("Route registered", "method", "POST", "path", "/events")

	router.POST("/events/:event_id/consume", eventController.ConsumeTicket)
	slog.Debug("Route registered", "method", "POST", "path", "/events/:event_id/consume")

	// Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	if err := router.Run(":" + port); err != nil {
		slog.Error("Server failed to start", "error", err)
		log.Fatal(err)
	}
}
