package repositories

import (

	"github.com/jmoiron/sqlx"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
)


type UserRepository interface {
	CreateUser(user *models.Register) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.Register) error {
	query := `INSERT INTO users (name, email, password) values ($1, $2, $3)`

	if _, err := r.db.Exec(query, user.Name, user.Email, user.Password); err != nil {
		return err
	}

	return nil 
}


func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	
	user := models.User{}
	query := `SELECT name, email, password, team, role, avatar_url FROM users where email = $1`

	if err := r.db.Get(&user, query, email); err != nil {
		return nil, err
	}

	return &user, nil

}
