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
