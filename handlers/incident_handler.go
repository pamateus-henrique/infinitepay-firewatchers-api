package handlers

import (
	"fmt"
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
	
	incidentInputModel := new(models.IncidentInput)

	if err := c.BodyParser(incidentInputModel); err != nil {
		log.Printf("CreateIncident: Error parsing request body: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input format")
	}

	incidentID, err := h.incidentService.CreateIncident(c.Context(), incidentInputModel)

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
		"msg":   "Updated incident summary",
		"data": "",
	})
}


func (h *IncidentHandler) UpdateIncidentStatus(c *fiber.Ctx) error {
	log.Println("UpdateIncidentStatus: Started processing request")
	
	IncidentStatus := new(models.IncidentStatus)

	if err := c.BodyParser(IncidentStatus); err != nil {
		log.Printf("UpdateIncidentStatus: Error parsing request body: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input format")
	}

	if err := h.incidentService.UpdateIncidentStatus(IncidentStatus); err != nil {
		log.Printf("UpdateIncidentStatus: error while updating incident summary: %v", err)
		return err;
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": "false",
		"msg":   "Updated incident status",
		"data": "",
	})
}

func (h *IncidentHandler) UpdateIncidentSeverity(c *fiber.Ctx) error {
	log.Println("UpdateIncidentSeverity: Started processing request")
	
	incidentSeverity := new(models.IncidentSeverity)

	if err := c.BodyParser(incidentSeverity); err != nil {
		log.Printf("UpdateIncidentSeverity: Error parsing request body: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input format")
	}

	if err := h.incidentService.UpdateIncidentSeverity(incidentSeverity); err != nil {
		log.Printf("UpdateIncidentSeverity: error while updating incident severity: %v", err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": "false",
		"msg":   "Updated incident severity",
		"data":  "",
	})
}

func (h *IncidentHandler) UpdateIncidentType(c *fiber.Ctx) error {
	log.Println("UpdateIncidentType: Started processing request")
	
	incidentType := new(models.IncidentType)

	if err := c.BodyParser(incidentType); err != nil {
		log.Printf("UpdateIncidentSeverity: Error parsing request body: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input format")
	}

	if err := h.incidentService.UpdateIncidentType(incidentType); err != nil {
		log.Printf("UpdateIncidentType: error while updating incident type: %v", err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": "false",
		"msg":   "Updated incident type",
		"data":  "",
	})
}

func (h *IncidentHandler) UpdateIncidentRoles(c *fiber.Ctx) error {
	log.Println("UpdateIncidentRoles: Started processing request")
	
	incidentRoles := new(models.IncidentRoles)

	if err := c.BodyParser(incidentRoles); err != nil {
		log.Printf("UpdateIncidentRoles: Error parsing request body: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input format")
	}

	fmt.Println(incidentRoles)

	if err := h.incidentService.UpdateIncidentRoles(incidentRoles); err != nil {
		log.Printf("UpdateIncidentRoles: error while updating incident roles: %v", err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": "false",
		"msg":   "Updated incident roles",
		"data":  "",
	})
}

func (h *IncidentHandler) UpdateIncidentCustomFields(c *fiber.Ctx) error {
	log.Println("UpdateIncidentCustomFields: Started processing request")

	incidentCustomFields := new(models.IncidentCustomFieldsUpdate)

	if err := c.BodyParser(incidentCustomFields); err != nil {
		log.Printf("UpdateIncidentCustomFields: Error parsing request body: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input format")
	}

	if err := h.incidentService.UpdateIncidentCustomFields(incidentCustomFields); err != nil {
		log.Printf("UpdateIncidentCustomFields: error while updating incident custom fields: %v", err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Updated incident custom fields",
		"data":  "",
	})
}
