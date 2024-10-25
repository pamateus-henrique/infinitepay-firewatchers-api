// middlewares/jwt_middleware.go
package middlewares

import (
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func JWTMiddleware() fiber.Handler {
    cfg := config.GetConfig()

    return jwtware.New(jwtware.Config{
        SigningKey:   []byte(cfg.JWTSecret),
        ContextKey:   "jwt",
        ErrorHandler: jwtError,
        TokenLookup:  "cookie:jwt,header:Authorization",
    })
}

func jwtError(c *fiber.Ctx, err error) error {
    if err.Error() == "Missing or malformed JWT" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Missing or malformed JWT",
        })
    }
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
        "error": "Invalid or expired JWT",
    })
}
