package models

type User struct {
	Name string `db:"name" json:"name" validate:"required"`
	Email string `db:"email" json:"email" validate:"required,email"`
	Password string `db:"password" json:"password" validate:"required,gte=8,lte=255"`
	Team string `db:"team" json:"team" validate:"required"`
	Role string `db:"role" json:"role" validate:"required"`
	Avatar_url string `db:"avatar_url" json:"avatar_url" validate:"required"`
}
