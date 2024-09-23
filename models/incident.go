package models

import "time"

type IncidentInput struct {
	Title string `json:"title" db:"title"`
	Type int `json:"type" db:"type"`
	Severity string `json:"severity" db:"severity"`
	Summary string `json:"summary" db:"summary"`
	Impact *string `json:"impact,omitempty" db:"impact"`
	Source *string `json:"source,omitempty" db:"incident_source"`
	Products []int `json:"products,omitempty" db:"product_id"`
	Areas []int `json:"areas,omitempty" db:"area_id"`
	Indicators []int `json:"indicators,omitempty" db:"indicator_id"`
	ImpactStartedAt *time.Time `json:"started_at" db:"impact_started_at"`
	SlackThread *string `json:"slack_thread,omitempty" db:"slack_thread"`
}
