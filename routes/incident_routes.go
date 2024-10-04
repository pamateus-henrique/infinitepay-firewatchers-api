package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/handlers"

	"github.com/pamateus-henrique/infinitepay-firewatchers-api/services"
)

func SetupIncidentRoutes(app *fiber.App, services *services.Services) {
	incidentHandler := handlers.NewIncidentHandler(services.IncidentService)

    // Protected routes
    api := app.Group("/api/v1/incidents")
    api.Post("/create", incidentHandler.CreateIncident)
	api.Get("/", incidentHandler.GetIncidents)
	api.Get("/:id", incidentHandler.GetSingleIncident)
}
