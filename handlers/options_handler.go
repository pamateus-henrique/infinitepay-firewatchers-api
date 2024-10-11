package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/services"
)

type OptionsHandler struct {
	optionsService services.OptionsService
}

func NewOptionsHandler(optionsService services.OptionsService) *OptionsHandler {
	return &OptionsHandler{optionsService: optionsService}
}

func (h *OptionsHandler) GetTypes(c *fiber.Ctx) error {
	log.Println("GetTypes: Started processing request")

	types, err := h.optionsService.GetTypes()
	if err != nil {
		log.Printf("GetTypes: Error fetching types: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Error fetching types")
	}

	log.Printf("GetTypes: Successfully fetched %d types", len(types))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Fetched types",
		"data": fiber.Map{
			"types": types,
		},
	})
}
