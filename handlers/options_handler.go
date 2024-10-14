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
			"type": types,
		},
	})
}

func (h *OptionsHandler) GetStatuses(c *fiber.Ctx) error {
	log.Println("GetStatuses: Started processing request")

	statuses, err := h.optionsService.GetStatuses()
	if err != nil {
		log.Printf("GetStatuses: Error fetching statuses: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Error fetching statuses")
	}

	log.Printf("GetStatuses: Successfully fetched %d statuses", len(statuses))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Fetched statuses",
		"data": fiber.Map{
			"status": statuses,
		},
	})
}

func (h *OptionsHandler) GetSeverities(c *fiber.Ctx) error {
	log.Println("GetSeverities: Started processing request")

	severities, err := h.optionsService.GetSeverities()
	if err != nil {
		log.Printf("GetSeverities: Error fetching severities: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Error fetching severities")
	}

	log.Printf("GetSeverities: Successfully fetched %d severities", len(severities))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Fetched severities",
		"data": fiber.Map{
			"severity": severities,
		},
	})
}
