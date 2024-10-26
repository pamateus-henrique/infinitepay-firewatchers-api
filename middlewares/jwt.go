// middlewares/jwt_middleware.go
package middlewares

import (


	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/config"
)

func JWTMiddleware() fiber.Handler {
	cfg := config.GetConfig()

	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(cfg.JWTSecret),
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
		TokenLookup:  "cookie:jwt,header:Authorization",
		SuccessHandler: func(c *fiber.Ctx) error {
			// Get the user claims from the token
			token := c.Locals("jwt").(*jwt.Token)
			claims := token.Claims.(jwt.MapClaims)

			// Extract the user_id from the claims
			userId, ok := claims["user_id"].(float64)

			if !ok {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid token claims",
				})
			}

			// Attach the user_id to the request locals
			c.Locals("user_id", int(userId))

			return c.Next()
		},
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or malformed JWT",
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Invalid or expired JWT",
	})
}
