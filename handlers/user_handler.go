// handlers/user_handler.go
package handlers

import (
	"fmt"
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

    // Create Set-Cookie header
    cookieValue := fmt.Sprintf("jwt=%s; HttpOnly; Path=/; Max-Age=%d; SameSite=Lax", jwt, 3600*24)
    c.Set("Set-Cookie", cookieValue)

    log.Println("Login: Set-Cookie header set successfully")

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

func (h *UserHandler) GetAllUsersPublicData(c *fiber.Ctx) error {
    log.Println("GetAllUsersPublicData: Started processing request")

    users, err := h.userService.GetAllUsersPublicData()
    if err != nil {
        log.Printf("GetAllUsersPublicData: Error retrieving users' public data: %v", err)
        return fiber.NewError(fiber.StatusInternalServerError, "Error retrieving users' data")
    }

    log.Printf("GetAllUsersPublicData: Successfully retrieved public data for %d users", len(users))

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Fetched severities",
		"data": fiber.Map{
			"users": users,
		},
	})
}
