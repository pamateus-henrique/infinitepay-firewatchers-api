package services

import (
	"log"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/repositories"
)

type OptionsService interface {
	GetTypes() ([]*models.Type, error)
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