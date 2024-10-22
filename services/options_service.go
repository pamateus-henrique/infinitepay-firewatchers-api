package services

import (
	"log"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/repositories"
)

type OptionsService interface {
	GetTypes() ([]*models.Type, error)
	GetStatuses() ([]*models.Status, error)
	GetSeverities() ([]*models.Severity, error)
	GetProducts() ([]*models.Product, error)
	GetAreas() ([]*models.Area, error)
	GetPerformanceIndicators() ([]*models.PerformanceIndicator, error)
	GetFaultySystems() ([]*models.FaultySystem, error)
	GetCauses() ([]*models.Cause, error)
}

type optionsService struct {
	optionsRepository repositories.OptionsRepository
}

func NewOptionsService(optionsRepository repositories.OptionsRepository) OptionsService {
	return &optionsService{optionsRepository: optionsRepository}
}

func (s *optionsService) GetTypes() ([]*models.Type, error) {
	log.Println("GetTypes: Starting types retrieval process")

	types, err := s.optionsRepository.GetTypes()
	if err != nil {
		log.Printf("GetTypes: Error retrieving types: %v", err)
		return nil, err
	}

	log.Printf("GetTypes: Successfully retrieved %d types", len(types))
	return types, nil
}

func (s *optionsService) GetStatuses() ([]*models.Status, error) {
	log.Println("GetStatuses: Starting statuses retrieval process")

	statuses, err := s.optionsRepository.GetStatuses()
	if err != nil {
		log.Printf("GetStatuses: Error retrieving statuses: %v", err)
		return nil, err
	}

	log.Printf("GetStatuses: Successfully retrieved %d statuses", len(statuses))
	return statuses, nil
}

func (s *optionsService) GetSeverities() ([]*models.Severity, error) {
	log.Println("GetSeverities: Starting severities retrieval process")

	severities, err := s.optionsRepository.GetSeverities()
	if err != nil {
		log.Printf("GetSeverities: Error retrieving severities: %v", err)
		return nil, err
	}

	log.Printf("GetSeverities: Successfully retrieved %d severities", len(severities))
	return severities, nil
}

func (s *optionsService) GetProducts() ([]*models.Product, error) {
	log.Println("GetProducts: Starting products retrieval process")

	products, err := s.optionsRepository.GetProducts()
	if err != nil {
		log.Printf("GetProducts: Error retrieving products: %v", err)
		return nil, err
	}

	log.Printf("GetProducts: Successfully retrieved %d products", len(products))
	return products, nil
}

func (s *optionsService) GetAreas() ([]*models.Area, error) {
	log.Println("GetAreas: Starting areas retrieval process")

	areas, err := s.optionsRepository.GetAreas()
	if err != nil {
		log.Printf("GetAreas: Error retrieving areas: %v", err)
		return nil, err
	}

	log.Printf("GetAreas: Successfully retrieved %d areas", len(areas))
	return areas, nil
}

func (s *optionsService) GetPerformanceIndicators() ([]*models.PerformanceIndicator, error) {
	log.Println("GetPerformanceIndicators: Starting performance indicators retrieval process")

	indicators, err := s.optionsRepository.GetPerformanceIndicators()
	if err != nil {
		log.Printf("GetPerformanceIndicators: Error retrieving performance indicators: %v", err)
		return nil, err
	}

	log.Printf("GetPerformanceIndicators: Successfully retrieved %d performance indicators", len(indicators))
	return indicators, nil
}

func (s *optionsService) GetFaultySystems() ([]*models.FaultySystem, error) {
	log.Println("GetFaultySystems: Starting faulty systems retrieval process")

	systems, err := s.optionsRepository.GetFaultySystems()
	if err != nil {
		log.Printf("GetFaultySystems: Error retrieving faulty systems: %v", err)
		return nil, err
	}

	log.Printf("GetFaultySystems: Successfully retrieved %d faulty systems", len(systems))
	return systems, nil
}

func (s *optionsService) GetCauses() ([]*models.Cause, error) {
	log.Println("GetCauses: Starting causes retrieval process")

	causes, err := s.optionsRepository.GetCauses()
	if err != nil {
		log.Printf("GetCauses: Error retrieving causes: %v", err)
		return nil, err
	}

	log.Printf("GetCauses: Successfully retrieved %d causes", len(causes))
	return causes, nil
}
