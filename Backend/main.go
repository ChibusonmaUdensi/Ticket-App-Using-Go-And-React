package main

import (
	"fmt"
	"github.com/ChibusonmaUdensi/Ticket-App-Using-Go-And-React/middlewares"
	"github.com/ChibusonmaUdensi/Ticket-App-Using-Go-And-React/config"
	"github.com/ChibusonmaUdensi/Ticket-App-Using-Go-And-React/db"
	"github.com/ChibusonmaUdensi/Ticket-App-Using-Go-And-React/services"
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
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	//servcie
	authService :=services.NewAuthService(authRepository)

	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

    app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
//This main function sets up a structured
// web application using the Fiber framework, integrating various
//components like configuration management, database access,service logic, 
//and route handling to create a functional ticket booking API. 
//Each part of the application is neatly separated into its own package, 
//promoting maintainability and scalability.