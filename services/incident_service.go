package services

import (
	"log"

	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/repositories"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/validators"
)

type IncidentService interface {
	CreateIncident(incidentInput *models.IncidentInput) (int, error)
	GetIncidents(queryParams *models.IncidentQueryParams) ([]*models.IncidentOverviewOutput, error)
	GetSingleIncident(incidentID int) (*models.IncidentOutput, error)
	UpdateIncidentSummary(incidentSummary *models.IncidentSummary) error
	UpdateIncidentStatus(IncidentStatus *models.IncidentStatus) error
	UpdateIncidentSeverity(incidentSeverity *models.IncidentSeverity) error
	UpdateIncidentType(incidentType *models.IncidentType) error
	UpdateIncidentRoles(incidentRoles *models.IncidentRoles) error
	UpdateIncidentCustomFields(incident *models.IncidentCustomFieldsUpdate) error
}

type incidentService struct {
	incidentRepository repositories.IncidentRepository
}

func NewIncidentService(incidentRepository repositories.IncidentRepository) IncidentService {
	return &incidentService{incidentRepository: incidentRepository}
}

func (s *incidentService) CreateIncident(incidentInput *models.IncidentInput) (int, error) {
	log.Println("CreateIncident: Starting incident creation process")

	if err := validators.ValidateStruct(incidentInput); err != nil {
		log.Printf("CreateIncident: Validation error: %v", err)
		return 0, &validators.ValidationError{Err: err}
	}

	log.Println("CreateIncident: Validation passed, creating incident")
	incidentID, err := s.incidentRepository.CreateIncident(incidentInput)

	if err != nil {
		log.Printf("CreateIncident: Error creating incident: %v", err)
		return 0, err
	}

	log.Printf("CreateIncident: Incident created successfully with ID: %d", incidentID)
	return incidentID, nil
}

func (s *incidentService) GetIncidents(queryParams *models.IncidentQueryParams) ([]*models.IncidentOverviewOutput, error) {
	log.Println("GetIncidents: Starting incidents retrieval process")

	if err := validators.ValidateStruct(queryParams); err != nil {
		log.Printf("GetIncidents: Validation error: %v", err)
		return nil, &validators.ValidationError{Err: err}
	}
	
	log.Println("GetIncidents: Validation passed, retrieving incidents")
	incidents, err := s.incidentRepository.GetIncidents(queryParams)

	if err != nil {
		log.Printf("GetIncidents: Error retrieving incidents: %v", err)
		return nil, err
	}

	log.Printf("GetIncidents: Successfully retrieved %d incidents", len(incidents))
	return incidents, err
}


func (s *incidentService) GetSingleIncident(incidentID int) (*models.IncidentOutput, error) {
	log.Println("GetIncidents: Starting incidents retrieval process")

	incident, err := s.incidentRepository.GetIncidentByID(incidentID)

	if err != nil {
		log.Printf("Problem retrieving incident %v", incidentID)
		return nil, err
	}

	return incident, err
}


func (s *incidentService) UpdateIncidentSummary(incidentSummary *models.IncidentSummary) error {
	log.Printf("UpdateIncidentSummary: Starting update process for incident ID %d", incidentSummary.ID)

	if err := validators.ValidateStruct(incidentSummary); err != nil {
		log.Printf("UpdateIncidentSummary: Validation error: %v", err)
		return &validators.ValidationError{Err: err}
	}

	err := s.incidentRepository.UpdateIncidentSummary(incidentSummary)
	if err != nil {
		log.Printf("UpdateIncidentSummary: Error updating incident summary: %v", err)
		return err
	}

	log.Printf("UpdateIncidentSummary: Successfully updated summary for incident ID %d", incidentSummary.ID)
	return nil
}


func (s *incidentService) UpdateIncidentStatus(IncidentStatus *models.IncidentStatus) error {
	log.Printf("UpdateIncidentStatus: Starting update process for incident ID %d", IncidentStatus.ID)

	if err := validators.ValidateStruct(IncidentStatus); err != nil {
		log.Printf("UpdateIncidentStatus: Validation error: %v", err)
		return &validators.ValidationError{Err: err}
	}

	err := s.incidentRepository.UpdateIncidentStatus(IncidentStatus)
	if err != nil {
		log.Printf("UpdateIncidentSummary: Error updating incident status: %v", err)
		return err
	}

	log.Printf("UpdateIncidentSummary: Successfully updated status for incident ID %d", IncidentStatus.ID)
	return nil
}

func (s *incidentService) UpdateIncidentSeverity(incidentSeverity *models.IncidentSeverity) error {
	log.Printf("UpdateIncidentSeverity: Starting update process for incident ID %d", incidentSeverity.ID)

	if err := validators.ValidateStruct(incidentSeverity); err != nil {
		log.Printf("UpdateIncidentSeverity: Validation error: %v", err)
		return &validators.ValidationError{Err: err}
	}

	err := s.incidentRepository.UpdateIncidentSeverity(incidentSeverity)
	if err != nil {
		log.Printf("UpdateIncidentSeverity: Error updating incident severity: %v", err)
		return err
	}

	log.Printf("UpdateIncidentSeverity: Successfully updated severity for incident ID %d", incidentSeverity.ID)
	return nil
}

func (s *incidentService) UpdateIncidentType(incidentType *models.IncidentType) error {
	log.Printf("UpdateIncidentType: Updating type for incident ID %v", incidentType)

	if err := validators.ValidateStruct(incidentType); err != nil {
		log.Printf("UpdateIncidentType: Validation error: %v", err)
		return &validators.ValidationError{Err: err}
	}

	err := s.incidentRepository.UpdateIncidentType(incidentType)
	if err != nil {
		log.Printf("UpdateIncidentType: Error updating incident type: %v", err)
		return err
	}

	log.Printf("UpdateIncidentType: Successfully updated type for incident ID %d", incidentType.ID)
	return nil
}

func (s *incidentService) UpdateIncidentRoles(incidentRoles *models.IncidentRoles) error {
	log.Printf("UpdateIncidentRoles: Starting update process for incident ID %d", incidentRoles.ID)

	if err := validators.ValidateStruct(incidentRoles); err != nil {
		log.Printf("UpdateIncidentRoles: Validation error: %v", err)
		return &validators.ValidationError{Err: err}
	}

	err := s.incidentRepository.UpdateIncidentRoles(incidentRoles)
	if err != nil {
		log.Printf("UpdateIncidentRoles: Error updating incident roles: %v", err)
		return err
	}

	log.Printf("UpdateIncidentRoles: Successfully updated roles for incident ID %d", incidentRoles.ID)
	return nil
}

func (s *incidentService) UpdateIncidentCustomFields(incident *models.IncidentCustomFieldsUpdate) error {
	log.Printf("UpdateIncidentCustomFields: Starting update process for incident ID %d", incident.ID)

	if err := validators.ValidateStruct(incident); err != nil {
		log.Printf("UpdateIncidentCustomFields: Validation error: %v", err)
		return &validators.ValidationError{Err: err}
	}

	err := s.incidentRepository.UpdateIncidentCustomFields(incident)
	if err != nil {
		log.Printf("UpdateIncidentCustomFields: Error updating incident custom fields: %v", err)
		return err
	}

	log.Printf("UpdateIncidentCustomFields: Successfully updated custom fields for incident ID %d", incident.ID)
	return nil
}
