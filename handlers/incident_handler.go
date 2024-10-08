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

	incidentID, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("GetSingleIncident: Invalid incident ID: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid incident ID")
	}

	incident, err := h.incidentService.GetSingleIncident(incidentID)

	if err != nil {
		log.Printf("GetSingleIncident: error while retrieving incident: %v", err)
		return err;
	}

	log.Printf("GetSingleIncident: retrived sucessfully: %v", incident.ID)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": "false",
		"msg":   "Fetched incident",
		"data": fiber.Map{
			"incidents": incident,
		},
	})
}

func (h *IncidentHandler) UpdateIncidentSummary(c *fiber.Ctx) error {
	log.Println("UpdateIncidentSummary: Started processing request")
	
	IncidentSummary := new(models.IncidentSummary)

	if err := c.BodyParser(IncidentSummary); err != nil {
		log.Printf("UpdateIncidentSummary: Error parsing request body: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input format")
	}

	if err := h.incidentService.UpdateIncidentSummary(IncidentSummary); err != nil {
		log.Printf("UpdateIncidentSummary: error while updating incident summary: %v", err)
		return err;
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": "false",
		"msg":   "Fetched incident",
		"data": "",
	})
}
