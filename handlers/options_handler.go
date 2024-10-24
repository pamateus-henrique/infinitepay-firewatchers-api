package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
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

func (h *OptionsHandler) GetProducts(c *fiber.Ctx) error {
	log.Println("GetProducts: Started processing request")

	products, err := h.optionsService.GetProducts()
	if err != nil {
		log.Printf("GetProducts: Error fetching products: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Error fetching products")
	}

	log.Printf("GetProducts: Successfully fetched %d products", len(products))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Fetched products",
		"data": fiber.Map{
			"products": products,
		},
	})
}

func (h *OptionsHandler) GetAreas(c *fiber.Ctx) error {
	log.Println("GetAreas: Started processing request")

	areas, err := h.optionsService.GetAreas()
	if err != nil {
		log.Printf("GetAreas: Error fetching areas: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Error fetching areas")
	}

	log.Printf("GetAreas: Successfully fetched %d areas", len(areas))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Fetched areas",
		"data": fiber.Map{
			"areas": areas,
		},
	})
}

func (h *OptionsHandler) GetPerformanceIndicators(c *fiber.Ctx) error {
	log.Println("GetPerformanceIndicators: Started processing request")

	indicators, err := h.optionsService.GetPerformanceIndicators()
	if err != nil {
		log.Printf("GetPerformanceIndicators: Error fetching performance indicators: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Error fetching performance indicators")
	}

	log.Printf("GetPerformanceIndicators: Successfully fetched %d performance indicators", len(indicators))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Fetched performance indicators",
		"data": fiber.Map{
			"performanceIndicators": indicators,
		},
	})
}

func (h *OptionsHandler) GetFaultySystems(c *fiber.Ctx) error {
	log.Println("GetFaultySystems: Started processing request")

	systems, err := h.optionsService.GetFaultySystems()
	if err != nil {
		log.Printf("GetFaultySystems: Error fetching faulty systems: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Error fetching faulty systems")
	}

	log.Printf("GetFaultySystems: Successfully fetched %d faulty systems", len(systems))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Fetched faulty systems",
		"data": fiber.Map{
			"faultySystems": systems,
		},
	})
}


func (h *OptionsHandler) GetCauses(c *fiber.Ctx) error {
    log.Println("GetCauses: Started processing request")

    causes, err := h.optionsService.GetCauses()
    if err != nil {
        log.Printf("GetCauses: Error fetching causes: %v", err)
        return fiber.NewError(fiber.StatusInternalServerError, "Error fetching causes")
    }

    log.Printf("GetCauses: Successfully fetched %d causes", len(causes))

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "error": false,
        "msg":   "Fetched causes",
        "data": fiber.Map{
            "causes": causes,
        },
    })
}


//remove this later on
func (h *OptionsHandler) GetSources(c *fiber.Ctx) error {
    sources := []models.Source{
        {ID: 1, Name: "Internal"},
        {ID: 2, Name: "External"},
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "error": false,
        "msg":   "Fetched sources",
        "data": fiber.Map{
            "source": sources,
        },
    })
}