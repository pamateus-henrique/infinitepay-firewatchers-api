package services

import (
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/repositories"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/validators"
)

type IncidentService interface {
	CreateIncident(incidentInput *models.IncidentInput) (int, error)
}

type incidentService struct {
	incidentRepository repositories.IncidentRepository
}

func NewIncidentService(incidentRepository repositories.IncidentRepository) IncidentService {
	return &incidentService{incidentRepository: incidentRepository}
}


func (s *incidentService) CreateIncident(incidentInput *models.IncidentInput) (int, error) {

	if err := validators.ValidateStruct(incidentInput); err != nil {
		return 0, &validators.ValidationError{Err: err}
	}

	incidentID, err := s.incidentRepository.CreateIncident(incidentInput)

	if err != nil {
		return 0, err
	}

	return incidentID, nil
}
