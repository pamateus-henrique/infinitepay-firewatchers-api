package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/handlers"

	"github.com/pamateus-henrique/infinitepay-firewatchers-api/middlewares"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/services"
)

func SetupIncidentRoutes(app *fiber.App, services *services.Services) {
	incidentHandler := handlers.NewIncidentHandler(services.IncidentService)

    // Protected routes
    api := app.Group("/api/v1/incidents", middlewares.JWTMiddleware())
    api.Post("/create", incidentHandler.CreateIncident)
}
