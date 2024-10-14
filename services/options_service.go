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
