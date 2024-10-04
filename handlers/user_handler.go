// handlers/user_handler.go
package handlers

import (
	"log"

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
    log.Println("Register: Started processing request")

    user := new(models.Register)
    if err := c.BodyParser(user); err != nil {
        log.Printf("Register: Error parsing request body: %v", err)
        return fiber.NewError(fiber.StatusBadRequest, "Invalid input format")
    }

    log.Printf("Register: Parsed user data: %+v", user)

    if err := h.userService.Register(user); err != nil {
        log.Printf("Register: Error registering user: %v", err)
        return err
    }

    log.Println("Register: User registered successfully")

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "User registered successfully",
    })
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
    log.Println("Login: Started processing request")

    loginData := new(models.Login)

    if err := c.BodyParser(loginData); err != nil {
        log.Printf("Login: Error parsing request body: %v", err)
        return fiber.NewError(fiber.StatusBadRequest, "Invalid input format")
    }

    log.Printf("Login: Parsed login data: %+v", loginData)

    user, err := h.userService.Login(loginData)
    if err != nil {
       log.Printf("Login: Error during login: %v", err)
       return err
    }

    log.Printf("Login: User logged in successfully: %s", user.Name)

    jwt, err := utils.GenerateJWT(user.Name)
    if err != nil {
        log.Printf("Login: Error generating JWT: %v", err)
        return fiber.NewError(fiber.StatusInternalServerError)
    }

    log.Println("Login: JWT generated successfully")

    // Create cookie
    cookie := new(fiber.Cookie)
    cookie.Name = "jwt"
    cookie.Value = jwt
    cookie.HTTPOnly = true
    cookie.Secure = false  // Set to true if using HTTPS
    cookie.SameSite = "Lax"
    cookie.MaxAge = 3600 * 24  // 24 hours
    cookie.Path = "/"
    c.Cookie(cookie)

    log.Println("Login: Cookie set successfully")

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
