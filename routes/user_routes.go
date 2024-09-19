// routes/user_routes.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/handlers"

	// "github.com/pamateus-henrique/infinitepay-firewatchers-api/middlewares"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/services"
)

func SetupUserRoutes(app *fiber.App, services *services.Services) {
    userHandler := handlers.NewUserHandler(services.UserService)

    // Public routes
    app.Post("/register", userHandler.Register)
    app.Post("/login", userHandler.Login)

    // Protected routes
    // api := app.Group("/api", middlewares.JWTMiddleware())
    // api.Get("/users/:id", userHandler.GetUser)
}
