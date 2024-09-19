// routes/routes.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/services"
)

func SetupRoutes(app *fiber.App, services *services.Services) {
    // Setup user routes
    SetupUserRoutes(app, services)
    // Setup more routes here (e.g., product routes)
}
