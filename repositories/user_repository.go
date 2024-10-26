package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
)


type UserRepository interface {
	CreateUser(user *models.Register) error
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsersPublicData() ([]*models.UserPublicData, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.Register) error {
	log.Println("CreateUser: Starting user creation process")
	query := `INSERT INTO users (name, email, password, role, team) values ($1, $2, $3, $4, $5)`

	log.Printf("CreateUser: Executing query with name: %s, email: %s, role: Viewer, team: Cloudwalk", user.Name, user.Email)
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password, "Viewer", "Cloudwalk")
	if err != nil {
		log.Printf("CreateUser: Error executing query: %v", err)
		return err
	}

	log.Println("CreateUser: User created successfully")
	return nil 
}


func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	log.Printf("GetUserByEmail: Retrieving user with email: %s", email)
	
	user := models.User{}
	query := `SELECT id, name, email, password, team, role, avatar_url FROM users where email = $1`

	log.Println("GetUserByEmail: Executing query")
	err := r.db.Get(&user, query, email)
	if err != nil {
		log.Printf("GetUserByEmail: Error executing query: %v", err)
		return nil, err
	}

	log.Println("GetUserByEmail: User retrieved successfully")
	return &user, nil

}

func (r *userRepository) GetAllUsersPublicData() ([]*models.UserPublicData, error) {
	log.Println("GetAllUsersPublicData: Starting retrieval of all users' public data")

	query := `SELECT id, name, avatar_url FROM users`

	rows, err := r.db.Queryx(query)
	if err != nil {
		log.Printf("GetAllUsersPublicData: Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []*models.UserPublicData

	for rows.Next() {
		var user models.UserPublicData
		if err := rows.StructScan(&user); err != nil {
			log.Printf("GetAllUsersPublicData: Error scanning row: %v", err)
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		log.Printf("GetAllUsersPublicData: Error after scanning all rows: %v", err)
		return nil, err
	}

	log.Printf("GetAllUsersPublicData: Successfully retrieved public data for %d users", len(users))
	return users, nil
}
