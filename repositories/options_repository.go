package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
)


type OptionsRepository interface {
	GetTypes() ([]*models.Type, error)
}

type optionsRepository struct {
	db *sqlx.DB
}

func NewOptionsRepository(db *sqlx.DB) OptionsRepository {
    return &optionsRepository{db: db}
}


func (r *optionsRepository) GetTypes() ([]*models.Type, error){

	query := "SELECT id, name FROM types where active = true"

	rows, err := r.db.Queryx(query);

	if err != nil {
		return nil, err
	}

	var types []*models.Type

	for rows.Next() {
		var incidenttypes models.Type
		if err := rows.StructScan(&incidenttypes); err != nil {
			return nil, err
		}
		types = append(types, &incidenttypes)
	}

	if err = rows.Err(); err != nil {
        log.Printf("getTypes: Error after scanning all rows: %v", err)
        return nil, err
    }

	return types, nil
}