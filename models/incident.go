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

type IncidentQueryParams struct {
		Status   *string `query:"status"`
		Category *string `query:"category"`
		Severity *string `query:"severity"`
}

type IncidentOverviewOutput struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Type string `json:"type"`
	Severity string `json:"severity"`
	Summary string `json:"summary"`
	ImpactStartedAt time.Time `json:"impactStartedAt" db:"impact_started_at"`
	Status string `json:"status"`
	Reporter int `json:"reporter"`
	LeadAvatar string `json:"leadAvatar" db:"avatar_url"`
}