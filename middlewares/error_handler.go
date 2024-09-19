package middlewares

import (
    "errors"

    "github.com/gofiber/fiber/v2"
    "github.com/pamateus-henrique/infinitepay-firewatchers-api/validators"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
    // Default to 500 Internal Server Error
    code := fiber.StatusInternalServerError
    message := "Internal Server Error"

    // Check if it's a *fiber.Error
    var fiberErr *fiber.Error
    if errors.As(err, &fiberErr) {
        code = fiberErr.Code
        message = fiberErr.Message
    }

    // Check for ValidationError
    var ve *validators.ValidationError
    if errors.As(err, &ve) {
        code = fiber.StatusBadRequest
        return c.Status(code).JSON(fiber.Map{
            "error":   true,
            "message": "Validation failed",
            "details": ve.ErrorMessages(),
        })
    }

    // Handle other custom errors here if necessary

    // Default error response
    return c.Status(code).JSON(fiber.Map{
        "error":   true,
        "message": message,
    })
}