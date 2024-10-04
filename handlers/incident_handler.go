package handlers

import (
	"log"

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

func (h *IncidentHandler) CreateIncident(c *fiber.Ctx) error {
	log.Println("CreateIncident: Started processing request")
	
	incidentInputModel := new(models.IncidentInput)

	if err := c.BodyParser(incidentInputModel); err != nil {
		log.Printf("CreateIncident: Error parsing request body: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input format")
	}

	log.Printf("CreateIncident: Parsed input model: %+v", incidentInputModel)

	incidentID, err := h.incidentService.CreateIncident(incidentInputModel)

	if err != nil {
		log.Printf("CreateIncident: Error creating incident: %v", err)
		return err
	}

	log.Printf("CreateIncident: Successfully created incident with ID: %v", incidentID)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": "false",
		"msg":   "Incident Created",
		"data": fiber.Map{
			"incidentID": incidentID,
		},
	})
}

func (h *IncidentHandler) GetIncidents(c *fiber.Ctx) error {
	log.Println("GetIncidents: Started processing request")
	
	params := new(models.IncidentQueryParams)

	if err := c.QueryParser(params); err != nil {
		log.Printf("GetIncidents: Error parsing query parameters: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "invalid input format")
	}

	log.Printf("GetIncidents: Parsed query parameters: %+v", params)

	incidents, err := h.incidentService.GetIncidents(params)
	
	if err != nil {
		log.Printf("GetIncidents: Error fetching incidents: %v", err)
		return err
	}

	log.Printf("GetIncidents: Successfully fetched %d incidents", len(incidents))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": "false",
		"msg":   "Fetched incidents",
		"data": fiber.Map{
			"incidents": incidents,
		},
	})
}



func (h *IncidentHandler) GetSingleIncident(c *fiber.Ctx) error {
	log.Println("GetSingleIncident: Started processing request")

	incident, err := h.incidentService.GetSingleIncident(90)

	if err != nil {
		return err;
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": "false",
		"msg":   "Fetched incident",
		"data": fiber.Map{
			"incidents": incident,
		},
	})
}
