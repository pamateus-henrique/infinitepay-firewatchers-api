package models

type Register struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=255"`
}

type Login struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=255"`
}