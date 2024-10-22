package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
)


type OptionsRepository interface {
	GetTypes() ([]*models.Type, error)
	GetStatuses() ([]*models.Status, error)
	GetSeverities() ([]*models.Severity, error)
	GetProducts() ([]*models.Product, error)
	GetAreas() ([]*models.Area, error)
	GetPerformanceIndicators() ([]*models.PerformanceIndicator, error)
	GetFaultySystems() ([]*models.FaultySystem, error)
	GetCauses() ([]*models.Cause, error)
}

type optionsRepository struct {
	db *sqlx.DB
}

func NewOptionsRepository(db *sqlx.DB) OptionsRepository {
	return &optionsRepository{db: db}
}


func (r *optionsRepository) GetTypes() ([]*models.Type, error) {
	query := "SELECT id, name FROM types WHERE active = true"

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []*models.Type

	for rows.Next() {
		var typeItem models.Type
		if err := rows.StructScan(&typeItem); err != nil {
			return nil, err
		}
		types = append(types, &typeItem)
	}

	if err = rows.Err(); err != nil {
		log.Printf("GetTypes: Error after scanning all rows: %v", err)
		return nil, err
	}

	return types, nil
}

func (r *optionsRepository) GetStatuses() ([]*models.Status, error) {
	query := "SELECT id, name FROM statuses WHERE active = true"

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []*models.Status

	for rows.Next() {
		var status models.Status
		if err := rows.StructScan(&status); err != nil {
			return nil, err
		}
		statuses = append(statuses, &status)
	}

	if err = rows.Err(); err != nil {
		log.Printf("GetStatuses: Error after scanning all rows: %v", err)
		return nil, err
	}

	return statuses, nil
}

func (r *optionsRepository) GetSeverities() ([]*models.Severity, error) {
	query := "SELECT id, name FROM severities WHERE active = true"

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var severities []*models.Severity

	for rows.Next() {
		var severity models.Severity
		if err := rows.StructScan(&severity); err != nil {
			return nil, err
		}
		severities = append(severities, &severity)
	}

	if err = rows.Err(); err != nil {
		log.Printf("GetSeverities: Error after scanning all rows: %v", err)
		return nil, err
	}

	return severities, nil
}

func (r *optionsRepository) GetProducts() ([]*models.Product, error) {
	query := "SELECT id, name FROM products WHERE active = true"

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product

	for rows.Next() {
		var product models.Product
		if err := rows.StructScan(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err = rows.Err(); err != nil {
		log.Printf("GetProducts: Error after scanning all rows: %v", err)
		return nil, err
	}

	return products, nil
}

func (r *optionsRepository) GetAreas() ([]*models.Area, error) {
	query := "SELECT id, name FROM areas WHERE active = true"

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var areas []*models.Area

	for rows.Next() {
		var area models.Area
		if err := rows.StructScan(&area); err != nil {
			return nil, err
		}
		areas = append(areas, &area)
	}

	if err = rows.Err(); err != nil {
		log.Printf("GetAreas: Error after scanning all rows: %v", err)
		return nil, err
	}

	return areas, nil
}

func (r *optionsRepository) GetPerformanceIndicators() ([]*models.PerformanceIndicator, error) {
	query := "SELECT id, name FROM performance_indicators WHERE active = true"

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var indicators []*models.PerformanceIndicator

	for rows.Next() {
		var indicator models.PerformanceIndicator
		if err := rows.StructScan(&indicator); err != nil {
			return nil, err
		}
		indicators = append(indicators, &indicator)
	}

	if err = rows.Err(); err != nil {
		log.Printf("GetPerformanceIndicators: Error after scanning all rows: %v", err)
		return nil, err
	}

	return indicators, nil
}

func (r *optionsRepository) GetFaultySystems() ([]*models.FaultySystem, error) {
	query := "SELECT id, name FROM faulty_systems WHERE active = true"

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var systems []*models.FaultySystem

	for rows.Next() {
		var system models.FaultySystem
		if err := rows.StructScan(&system); err != nil {
			return nil, err
		}
		systems = append(systems, &system)
	}

	if err = rows.Err(); err != nil {
		log.Printf("GetFaultySystems: Error after scanning all rows: %v", err)
		return nil, err
	}

	return systems, nil
}

func (r *optionsRepository) GetCauses() ([]*models.Cause, error) {
	query := "SELECT id, name FROM causes WHERE active = true"

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var causes []*models.Cause

	for rows.Next() {
		var cause models.Cause
		if err := rows.StructScan(&cause); err != nil {
			return nil, err
		}
		causes = append(causes, &cause)
	}

	if err = rows.Err(); err != nil {
		log.Printf("GetCauses: Error after scanning all rows: %v", err)
		return nil, err
	}

	return causes, nil
}
