package repositories

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
)

type IncidentRepository interface {
	CreateIncident(incident *models.IncidentInput) (int, error)
	GetIncidents(queryParams *models.IncidentQueryParams) ([]*models.IncidentOverviewOutput, error)
	GetIncidentByID(id int) (*models.IncidentOutput, error)
	UpdateIncidentSummary(incident *models.IncidentSummary) error
}

type incidentRepository struct {
	db *sqlx.DB
}

func NewIncidentRepository(db *sqlx.DB) IncidentRepository {
	return &incidentRepository{db: db}
}

func (r *incidentRepository) CreateIncident(incident *models.IncidentInput) (int, error) {
	log.Println("CreateIncident: Starting transaction")
	tx, err := r.db.Beginx()
	if err != nil {
		log.Printf("CreateIncident: Error starting transaction: %v", err)
		return 0, err
	}

	defer func() {
		if err != nil {
			log.Printf("CreateIncident: Rolling back transaction due to error: %v", err)
			tx.Rollback()
		} else {
			log.Println("CreateIncident: Committing transaction")
			err = tx.Commit()
			if err != nil {
				log.Printf("CreateIncident: Error committing transaction: %v", err)
			}
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

	log.Printf("CreateIncident: Executing query: %s", query)
	log.Printf("CreateIncident: Query params: %+v", params)

	var incidentID int
	query = tx.Rebind(query)
	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Printf("CreateIncident: Error preparing named statement: %v", err)
		return 0, err
	}
	err = stmt.Get(&incidentID, params)
	if err != nil {
		log.Printf("CreateIncident: Error executing query: %v", err)
		return 0, err
	}

	log.Printf("CreateIncident: Incident created with ID: %d", incidentID)

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

	log.Println("CreateIncident: Successfully created incident and related data")
	return incidentID, nil
}

func (r *incidentRepository) GetIncidents(queryParams *models.IncidentQueryParams) ([]*models.IncidentOverviewOutput, error) {
	log.Println("GetIncidents: Starting query construction")
	
	query := `SELECT i.id, i.title, t.name as type, i.severity, i.summary, i.status, i.impact_started_at, u.name as lead, u.avatar_url FROM incidents as i  LEFT JOIN types as t on t.id = i.type LEFT JOIN users as u on i.lead = u.id WHERE 1=1 `

	params := make(map[string]interface{})

	if queryParams.Category != nil {
		query += "AND i.category = :category"
		params["category"] = *queryParams.Category
	}

	if queryParams.Severity != nil {
		query += "AND i.severity = :severity"
		params["severity"] = *queryParams.Severity
	}

	if queryParams.Status != nil {
		query += "AND i.status = :status"
		params["status"] = *queryParams.Status
	}
	
    rows, err := r.db.NamedQuery(query, params)
    if err != nil {
        log.Printf("GetIncidents: Error executing query: %v", err)
        return nil, err
    }
    defer rows.Close()

    var incidents []*models.IncidentOverviewOutput
    for rows.Next() {
        var incident models.IncidentOverviewOutput
        if err := rows.StructScan(&incident); err != nil {
            log.Printf("GetIncidents: Error scanning row: %v", err)
            return nil, err
        }
        incidents = append(incidents, &incident)
    }

    if err = rows.Err(); err != nil {
        log.Printf("GetIncidents: Error after scanning all rows: %v", err)
        return nil, err
    }

	log.Printf("GetIncidents: Successfully retrieved %d incidents", len(incidents))
	return incidents, nil;
}


func (r *incidentRepository) GetIncidentByID(id int) (*models.IncidentOutput, error) {
    log.Printf("GetIncidentByID: Starting query for incident ID %d", id)

    query := 
	`
	SELECT 
    	i.*,
    	lead_user.name AS lead_name,
		lead_user.avatar_url AS lead_avatar,
    	reporter_user.name AS reporter_name,
		reporter_user.avatar_url AS reporter_avatar,
    	qe_user.name AS qe_name,
		qe_user.avatar_url AS qe_avatar
	FROM 
    	incidents i
	LEFT JOIN 
    	users lead_user ON i.lead = lead_user.id
	LEFT JOIN 
    	users reporter_user ON i.reporter = reporter_user.id
	LEFT JOIN 
    	users qe_user ON i.qe = qe_user.id	
	WHERE 
    	i.id = $1
	`

    incidentOutput := new(models.IncidentOutput)

    if err := r.db.Get(incidentOutput, query, id); err != nil {
        log.Printf("Error while retrieving incident %v: %s", id, err)
        return nil, err
    }

    if err := r.getRelatedData(incidentOutput); err != nil {
        log.Printf("Error while retrieving related data for incident %v: %s", id, err)
        return nil, err
    }

    log.Printf("GetIncidentByID: Successfully retrieved incident with ID %d", id)
    return incidentOutput, nil
}

func (r *incidentRepository) getRelatedData(incident *models.IncidentOutput) error {
    relatedTables := []struct {
        query    string
        dest     interface{}
    }{
        {
            query: `SELECT p.id, p.name 
                    FROM incident_products ip 
                    JOIN products p ON ip.product_id = p.id 
                    WHERE ip.incident_id = $1`,
            dest: &incident.Products,
        },
        {
            query: `SELECT a.id, a.name 
                    FROM incident_areas ia 
                    JOIN areas a ON ia.area_id = a.id 
                    WHERE ia.incident_id = $1`,
            dest: &incident.Areas,
        },
        {
            query: `SELECT c.id, c.name 
                    FROM incident_causes ic 
                    JOIN causes c ON ic.cause_id = c.id 
                    WHERE ic.incident_id = $1`,
            dest: &incident.Causes,
        },
        {
            query: `SELECT fs.id, fs.name 
                    FROM incident_faulty_systems ifs 
                    JOIN faulty_systems fs ON ifs.faulty_system_id = fs.id 
                    WHERE ifs.incident_id = $1`,
            dest: &incident.FaultySystems,
        },
        {
            query: `SELECT pi.id, pi.name 
                    FROM incident_performance_indicators ipi 
                    JOIN performance_indicators pi ON ipi.performance_indicator_id = pi.id 
                    WHERE ipi.incident_id = $1`,
            dest: &incident.PerformanceIndicators,
        },
    }

    for _, table := range relatedTables {
        rows, err := r.db.Query(table.query, incident.ID)
        if err != nil {
            return fmt.Errorf("error querying related data: %v", err)
        }
        defer rows.Close()

        relatedMap := make(map[int]string)
        for rows.Next() {
            var id int
            var name string
            if err := rows.Scan(&id, &name); err != nil {
                return fmt.Errorf("error scanning related data: %v", err)
            }
            relatedMap[id] = name
        }

        if err = rows.Err(); err != nil {
            return fmt.Errorf("error after scanning all related data: %v", err)
        }

        // Use reflection to set the map field in the struct
        reflect.ValueOf(table.dest).Elem().Set(reflect.ValueOf(relatedMap))
    }

    // Handle events separately as they don't have IDs
    // eventsQuery := `SELECT event FROM incident_events WHERE incident_id = $1`
    // rows, err := r.db.Query(eventsQuery, incident.ID)
    // if err != nil {
    //     return fmt.Errorf("error querying events: %v", err)
    // }
    // defer rows.Close()

    // var events []string
    // for rows.Next() {
    //     var event string
    //     if err := rows.Scan(&event); err != nil {
    //         return fmt.Errorf("error scanning event: %v", err)
    //     }
    //     events = append(events, event)
    // }

    // if err = rows.Err(); err != nil {
    //     return fmt.Errorf("error after scanning all events: %v", err)
    // }

    // incident.Events = events

    return nil
}

func (r *incidentRepository) UpdateIncidentSummary(incident *models.IncidentSummary) error {
    log.Printf("UpdateIncidentSummary: Updating summary for incident ID %d", incident.ID)

    query := `UPDATE incidents SET summary = $1 WHERE id = $2`

    result, err := r.db.Exec(query, incident.Summary, incident.ID)
    if err != nil {
        log.Printf("UpdateIncidentSummary: Error executing update query: %v", err)
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Printf("UpdateIncidentSummary: Error getting rows affected: %v", err)
        return err
    }

    if rowsAffected == 0 {
        log.Printf("UpdateIncidentSummary: No rows affected. Incident with ID %d not found", incident.ID)
        return fmt.Errorf("incident with ID %d not found", incident.ID)
    }

    log.Printf("UpdateIncidentSummary: Successfully updated summary for incident ID %d", incident.ID)
    return nil
}