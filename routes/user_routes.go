// routes/user_routes.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/handlers"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/middlewares"

	"github.com/pamateus-henrique/infinitepay-firewatchers-api/services"
)

func SetupUserRoutes(app *fiber.App, services *services.Services) {
    userHandler := handlers.NewUserHandler(services.UserService)

    // Public routes
    app.Post("/api/v1/auth/register", userHandler.Register)
    app.Post("/api/v1/auth/login", userHandler.Login)

    // Protected routes
    api := app.Group("/api/v1/users")
    
    api.Use(middlewares.JWTMiddleware())
    api.Get("/", userHandler.GetAllUsersPublicData)
}
