package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/handlers"

	"github.com/pamateus-henrique/infinitepay-firewatchers-api/services"
)

func SetupOptionsRoutes(app *fiber.App, services *services.Services) {
	optionsHandler := handlers.NewOptionsHandler(services.OptionsService)
	// Protected routes
	api := app.Group("/api/v1/options")
	api.Get("/types", optionsHandler.GetTypes)
	api.Get("/status", optionsHandler.GetStatuses)
	api.Get("/severity", optionsHandler.GetSeverities)
	api.Get("/products", optionsHandler.GetProducts)
	api.Get("/areas", optionsHandler.GetAreas)
	api.Get("/performance-indicators", optionsHandler.GetPerformanceIndicators)
	api.Get("/faulty-systems", optionsHandler.GetFaultySystems)
	api.Get("/causes", optionsHandler.GetCauses)
}
