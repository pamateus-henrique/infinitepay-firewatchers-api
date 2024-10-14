package models

type Type struct {
	ID int `json:"id" db:"id"`
	Name string  `json:"name" db:"name"`
}

type Status struct {
	ID int `json:"id" db:"id"`
	Name string  `json:"name" db:"name"`
}

type Severity struct {
	ID int `json:"id" db:"id"`
	Name string  `json:"name" db:"name"`
}


