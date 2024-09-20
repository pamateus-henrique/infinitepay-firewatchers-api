// handlers/user_handler.go
package handlers

import (
	// "strconv"

	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/services"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
    userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
    user := new(models.Register)
    if err := c.BodyParser(user); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid input format")
    }

    if err := h.userService.Register(user); err != nil {
        return err
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "User registered successfully",
    })
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
    loginData := new(models.Login)

    if err := c.BodyParser(loginData); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid input format")
    }

    user, err := h.userService.Login(loginData)
    if err != nil {
       return err
    }

    jwt, err := utils.GenerateJWT(user.Name)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	 // Create cookie
	 cookie := new(fiber.Cookie)
	 cookie.Name = "auth"
	 cookie.Value = jwt
	 cookie.HTTPOnly = true
	 c.Cookie(cookie)

	 return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg": "Login successful",
	})
}

// func (h *UserHandler) GetUser(c *fiber.Ctx) error {
//     idParam := c.Params("id")
//     id, err := strconv.Atoi(idParam)
//     if err != nil {
//         return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//             "error": "Invalid user ID",
//         })
//     }

//     user, err := h.userService.GetUser(id)
//     if err != nil {
//         return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
//             "error": "User not found",
//         })
//     }

//     return c.JSON(user)
// }
