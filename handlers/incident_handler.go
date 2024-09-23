package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/services"
)

type IncidentHandler struct{
	incidentService services.IncidentService
}

func NewIncidentHandler(incidentService services.IncidentService) *IncidentHandler{
	return &IncidentHandler{incidentService: incidentService}
}

func (h *IncidentHandler) CreateIncident(c *fiber.Ctx) error{
	
	incidentInputModel := new(models.IncidentInput)

	if err := c.BodyParser(incidentInputModel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input format");
	}

	incidentID, err := h.incidentService.CreateIncident(incidentInputModel)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": "false",
		"msg": "Incident Created",
		"data": fiber.Map{
			"incidentID": incidentID,
		},
	})

}