package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// Setup CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, https://yourdomain.com",  // Specify your frontend origins
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
		MaxAge:           300,
	}))

	//initialize repos
	userRepo := repositories.NewUserRepository(db)
	incidentRepo := repositories.NewIncidentRepository(db)
	optionsRepo := repositories.NewOptionsRepository(db)

	//initialize services
	services := &services.Services{
		UserService: services.NewUserService(userRepo),
		IncidentService: services.NewIncidentService(incidentRepo),
		OptionsService: services.NewOptionsService(optionsRepo),
	}

	//setup routes
	routes.SetupRoutes(app, services)

	//start server
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Oops... Server is not running! Reason: %v", err)
	}
}