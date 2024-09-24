package main

import (
	"fmt"

	"github.com/ChibusonmaUdensi/Ticket-App-Using-Go-And-React/config"
	"github.com/ChibusonmaUdensi/Ticket-App-Using-Go-And-React/db"
	"github.com/ChibusonmaUdensi/Ticket-App-Using-Go-And-React/handlers"
	"github.com/ChibusonmaUdensi/Ticket-App-Using-Go-And-React/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig :=config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)
	

    app := fiber.New(fiber.Config{
	AppName: "TicketBooking",
	ServerHeader: "Fiber",
	})
	

	eventRepository := repositories.NewEventRepository(db)

	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

    app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}