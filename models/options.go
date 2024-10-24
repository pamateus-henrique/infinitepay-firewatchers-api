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

type Product struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Area struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type PerformanceIndicator struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type FaultySystem struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Cause struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Source struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}





