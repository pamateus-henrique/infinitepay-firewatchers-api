package repositories

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
)

type IncidentRepository interface {
	CreateIncident(incident *models.IncidentInput) (int, error)
	GetIncidents(queryParams *models.IncidentQueryParams) ([]*models.IncidentOverviewOutput, error)
}

type incidentRepository struct {
	db *sqlx.DB
}

func NewIncidentRepository(db *sqlx.DB) IncidentRepository {
	return &incidentRepository{db: db}
}

func (r *incidentRepository) CreateIncident(incident *models.IncidentInput) (int, error) {
	//starts transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	//rollback in case anything goes wrong
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	fields := []string{"title", "type", "severity", "summary"}
	placeholders := []string{":title", ":type", ":severity", ":summary"}

	params := map[string]interface{}{
		"title": incident.Title,
		"type": incident.Type,
		"severity": incident.Severity,
		"summary": incident.Summary,
	}


	if incident.Impact != nil {
		fields = append(fields, "impact")
		placeholders = append(placeholders, ":impact")
		params["impact"] = *incident.Impact
	}

	if incident.Source != nil {
		fields = append(fields, "source")
		placeholders = append(placeholders, ":source")
		params["source"] = *incident.Source
	}

	if incident.ImpactStartedAt != nil {
		fields = append(fields, "impact_started_at")
		placeholders = append(placeholders, ":impact_started_at")
		params["impact_started_at"] = *incident.ImpactStartedAt
	}

	if incident.SlackThread != nil {
		fields = append(fields, "slack_thread")
		placeholders = append(placeholders, ":slack_thread")
		params["slack_thread"] = *incident.SlackThread
	}

	query := fmt.Sprintf(`INSERT INTO incidents (%s) VALUES (%s) returning id`, strings.Join(fields, ", "), strings.Join(placeholders, ", "))

	 // Execute the query
	 var incidentID int
	 query = tx.Rebind(query)
	 stmt, err := tx.PrepareNamed(query)
	 if err != nil {
		 return 0, err
	 }
	 err = stmt.Get(&incidentID, params)
	 if err != nil {
		 return 0, err
	 }

	     // Insert related products
		 if len(incident.Products) > 0 {
			productQuery := `INSERT INTO incident_products (incident_id, product_id) VALUES (:incident_id, :product_id)`
			for _, productID := range incident.Products {
				_, err := tx.NamedExec(productQuery, map[string]interface{}{
					"incident_id": incidentID,
					"product_id":  productID,
				})
				if err != nil {
					return 0, err
				}
			}
		}
	
		// Insert related areas
		if len(incident.Areas) > 0 {
			areaQuery := `INSERT INTO incident_areas (incident_id, area_id) VALUES (:incident_id, :area_id)`
			for _, areaID := range incident.Areas {
				_, err := tx.NamedExec(areaQuery, map[string]interface{}{
					"incident_id": incidentID,
					"area_id":     areaID,
				})
				if err != nil {
					return 0, err
				}
			}
		}
	
		// Insert related indicators
		if len(incident.Indicators) > 0 {
			indicatorQuery := `INSERT INTO incident_indicators (incident_id, indicator_id) VALUES (:incident_id, :indicator_id)`
			for _, indicatorID := range incident.Indicators {
				_, err := tx.NamedExec(indicatorQuery, map[string]interface{}{
					"incident_id":  incidentID,
					"indicator_id": indicatorID,
				})
				if err != nil {
					return 0, err
				}
			}
		}

		return incidentID, nil
}

func (r *incidentRepository) GetIncidents(queryParams *models.IncidentQueryParams) ([]*models.IncidentOverviewOutput, error){
	var incidents []*models.IncidentOverviewOutput
	
	query := `SELECT i.id, i.title, t.name as type, i.severity, i.summary, i.status, i.impact_started_at, i.reporter, u.avatar_url FROM incidents as i  LEFT JOIN types as t on t.id = i.type LEFT JOIN users as u on i.lead = u.id`;

	err := r.db.Select(&incidents, query)

	if err != nil {
		return nil, err
	};

	return incidents, nil;
}
