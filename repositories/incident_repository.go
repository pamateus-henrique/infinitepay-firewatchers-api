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
	UpdateIncidentStatus(incident *models.IncidentStatus) error
	UpdateIncidentSeverity(incident *models.IncidentSeverity) error
	UpdateIncidentType(incident *models.IncidentType) error
	UpdateIncidentRoles(incident *models.IncidentRoles) error
	UpdateIncidentCustomFields(incident *models.IncidentCustomFieldsUpdate) error
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

	fields := []string{"title", "type", "severity", "summary", "reporter", "status", "reported_at"}
	placeholders := []string{":title", ":type", ":severity", ":summary", ":reporter", ":status", ":reported_at"}

	params := map[string]interface{}{
		"title": incident.Title,
		"type": incident.Type,
		"severity": incident.Severity,
		"summary": incident.Summary,
        "reporter": incident.Reporter,
        "status": incident.Status,
        "reported_at": incident.ReportedAt,
	}


	if incident.Impact != nil {
		fields = append(fields, "impact")
		placeholders = append(placeholders, ":impact")
		params["impact"] = *incident.Impact
	}

    if incident.Lead != nil {
        fields = append(fields, "lead")
        placeholders = append(placeholders, ":lead")
        params["lead"] = *incident.Lead
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
	
	query := `SELECT i.id, i.title, i.type, i.severity, i.summary, i.status, i.impact_started_at, u.name as lead, u.avatar_url FROM incidents as i LEFT JOIN users as u on i.lead = u.id WHERE 1=1 `

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

func (r *incidentRepository) UpdateIncidentStatus(incident *models.IncidentStatus) error {
    log.Printf("UpdateIncidentStatus: Updating summary for incident ID %d", incident.ID)

    query := `UPDATE incidents SET status = $1 WHERE id = $2`

    result, err := r.db.Exec(query, incident.Status, incident.ID)
    if err != nil {
        log.Printf("UpdateIncidentStatus: Error executing update query: %v", err)
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Printf("UpdateIncidentStatus: Error getting rows affected: %v", err)
        return err
    }

    if rowsAffected == 0 {
        log.Printf("UpdateIncidentStatus: No rows affected. Incident with ID %d not found", incident.ID)
        return fmt.Errorf("incident with ID %d not found", incident.ID)
    }

    log.Printf("UpdateIncidentStatus: Successfully updated status for incident ID %d", incident.ID)
    return nil
}


func (r *incidentRepository) UpdateIncidentSeverity(incident *models.IncidentSeverity) error {
    log.Printf("UpdateIncidentSeverity: Updating severity for incident ID %d", incident.ID)

    query := `UPDATE incidents SET severity = $1 WHERE id = $2`

    result, err := r.db.Exec(query, incident.Severity, incident.ID)
    if err != nil {
        log.Printf("UpdateIncidentSeverity: Error executing update query: %v", err)
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Printf("UpdateIncidentSeverity: Error getting rows affected: %v", err)
        return err
    }

    if rowsAffected == 0 {
        log.Printf("UpdateIncidentSeverity: No rows affected. Incident with ID %d not found", incident.ID)
        return fmt.Errorf("incident with ID %d not found", incident.ID)
    }

    log.Printf("UpdateIncidentSeverity: Successfully updated severity for incident ID %d", incident.ID)
    return nil
}

func (r *incidentRepository) UpdateIncidentType(incident *models.IncidentType) error {
    log.Printf("UpdateIncidentType: Updating type for incident ID %d", incident.ID)

    query := `UPDATE incidents SET type = $1 WHERE id = $2`

    result, err := r.db.Exec(query, incident.Type, incident.ID)
    if err != nil {
        log.Printf("UpdateIncidentType: Error executing update query: %v", err)
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Printf("UpdateIncidentType: Error getting rows affected: %v", err)
        return err
    }

    if rowsAffected == 0 {
        log.Printf("UpdateIncidentType: No rows affected. Incident with ID %d not found", incident.ID)
        return fmt.Errorf("incident with ID %d not found", incident.ID)
    }

    log.Printf("UpdateIncidentType: Successfully updated type for incident ID %d", incident.ID)
    return nil
}

func (r *incidentRepository) UpdateIncidentRoles(incident *models.IncidentRoles) error {
    if incident == nil {
        return fmt.Errorf("UpdateIncidentRoles: incident cannot be nil")
    }
    if incident.ID == 0 {
        return fmt.Errorf("UpdateIncidentRoles: incident ID cannot be 0")
    }

    log.Printf("UpdateIncidentRoles: Updating roles for incident ID %d", incident.ID)

    // Start building the query
    query := "UPDATE incidents SET "
    var args []interface{}
    var setFields []string

    // Check if Lead is provided
    if incident.Lead != nil {
        setFields = append(setFields, "lead = ?")
        args = append(args, *incident.Lead)
    }

    // Check if QE is provided
    if incident.QE != nil {
        setFields = append(setFields, "qe = ?")
        args = append(args, *incident.QE)
    }

    // If no fields to update, return early
    if len(setFields) == 0 {
        log.Printf("UpdateIncidentRoles: No fields to update for incident ID %d", incident.ID)
        return nil
    }

    // Complete the query
    query += strings.Join(setFields, ", ")
    query += " WHERE id = ?"
    args = append(args, incident.ID)

    // Use sqlx for easier query building
    query = r.db.Rebind(query)
    result, err := r.db.Exec(query, args...)
    if err != nil {
        log.Printf("UpdateIncidentRoles: Error executing update query: %v", err)
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Printf("UpdateIncidentRoles: Error getting rows affected: %v", err)
        return err
    }

    if rowsAffected == 0 {
        log.Printf("UpdateIncidentRoles: No rows affected. Incident with ID %d not found", incident.ID)
        return fmt.Errorf("incident with ID %d not found", incident.ID)
    }

    log.Printf("UpdateIncidentRoles: Successfully updated roles for incident ID %d", incident.ID)
    return nil
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

        var relatedData []models.RelatedItem
        for rows.Next() {
            var item models.RelatedItem
            if err := rows.Scan(&item.ID, &item.Name); err != nil {
                return fmt.Errorf("error scanning related data: %v", err)
            }
            relatedData = append(relatedData, item)
        }

        if err = rows.Err(); err != nil {
            return fmt.Errorf("error after scanning all related data: %v", err)
        }

        // Use reflection to set the slice field in the struct
        destValue := reflect.ValueOf(table.dest).Elem()
        destValue.Set(reflect.ValueOf(relatedData))
    }

    return nil
}

func (r *incidentRepository) UpdateIncidentCustomFields(incident *models.IncidentCustomFieldsUpdate) error {
    tx, err := r.db.Beginx()
    if err != nil {
        return fmt.Errorf("error starting transaction: %v", err)
    }
    defer tx.Rollback()

    // Update fields in the incidents table
    updateQuery := "UPDATE incidents SET"
    updateParams := []interface{}{}
    paramCount := 1

    if incident.Impact != nil {
        updateQuery += fmt.Sprintf(" impact = $%d,", paramCount)
        updateParams = append(updateParams, *incident.Impact)
        paramCount++
    }
    if incident.Treatment != nil {
        updateQuery += fmt.Sprintf(" treatment = $%d,", paramCount)
        updateParams = append(updateParams, *incident.Treatment)
        paramCount++
    }
    if incident.Mitigator != nil {
        updateQuery += fmt.Sprintf(" mitigator = $%d,", paramCount)
        updateParams = append(updateParams, *incident.Mitigator)
        paramCount++
    }

    // Remove trailing comma
    updateQuery = strings.TrimSuffix(updateQuery, ",")
    updateQuery += fmt.Sprintf(" WHERE id = $%d", paramCount)
    updateParams = append(updateParams, incident.ID)

    if len(updateParams) > 1 { // Only execute if there are fields to update
        _, err = tx.Exec(updateQuery, updateParams...)
        if err != nil {
            return fmt.Errorf("error updating incident fields: %v", err)
        }
    }

    // Update related items
    relatedItems := []struct {
        items    []int
        table    string
        idColumn string
    }{
        {incident.Products, "incident_products", "product_id"},
        {incident.Areas, "incident_areas", "area_id"},
        {incident.Causes, "incident_causes", "cause_id"},
        {incident.FaultySystems, "incident_faulty_systems", "faulty_system_id"},
        {incident.PerformanceIndicators, "incident_performance_indicators", "performance_indicator_id"},
    }

    for _, item := range relatedItems {
            // Delete existing relations
            _, err = tx.Exec(fmt.Sprintf("DELETE FROM %s WHERE incident_id = $1", item.table), incident.ID)
            if err != nil {
                return fmt.Errorf("error deleting existing %s: %v", item.table, err)
            }

            // Insert new relations
            insertQuery := fmt.Sprintf("INSERT INTO %s (incident_id, %s) VALUES ($1, $2)", item.table, item.idColumn)
            for _, id := range item.items {
                _, err = tx.Exec(insertQuery, incident.ID, id)
                if err != nil {
                    return fmt.Errorf("error inserting new %s: %v", item.table, err)
                }
            }
    }

    return tx.Commit()
}
