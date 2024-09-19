package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/database"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/middlewares"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/repositories"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/routes"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/services"
)

func main(){
	db, err := database.OpenDBConnection();

	if err != nil {
		log.Fatal("Error connecting to the database, shutting down server.")
	}

	defer db.Close()

	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})


	//initialize repos
	userRepo := repositories.NewUserRepository(db)

	//initialize services
	services := &services.Services{
		UserService: services.NewUserService(userRepo),
	}

	//setup routes
	routes.SetupRoutes(app, services)

	//start server
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Oops... Server is not running! Reason: %v", err)
	}
}