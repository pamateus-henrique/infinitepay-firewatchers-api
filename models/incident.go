package models

import "time"

type RelatedItem struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type IncidentInput struct {
	Title           string     	`json:"title" db:"title"`
	Type            string     	`json:"type" db:"type"`
	Severity        string     	`json:"severity" db:"severity"`
	Summary         string     	`json:"summary" db:"summary"`
	Status 			string		`json:"status" db:"status"`
	Reporter		int			`db:"reporter"`
	Impact          *string    	`json:"impact,omitempty" db:"impact"`
	Source          *string    	`json:"source,omitempty" db:"incident_source"`
	Lead            *int        `json:"lead" db:"lead"`
	Products        []int      	`json:"products,omitempty" db:"product_id"`
	Areas           []int      	`json:"areas,omitempty" db:"area_id"`
	Indicators      []int      	`json:"indicators,omitempty" db:"indicator_id"`
	ImpactStartedAt *CustomTime 	`json:"impactStartedAt" db:"impact_started_at"`
	SlackThread     *string    	`json:"slack_thread,omitempty" db:"slack_thread"`
	ReportedAt		*CustomTime `db:"reported_at"`
}

type IncidentQueryParams struct {
	Status   *string `query:"status"`
	Category *string `query:"category"`
	Severity *string `query:"severity"`
}

type IncidentOverviewOutput struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Type            string    `json:"type"`
	Severity        string    `json:"severity"`
	Summary         string    `json:"summary"`
	ImpactStartedAt time.Time `json:"impactStartedAt" db:"impact_started_at"`
	Status          string    `json:"status"`
	Lead            string    `json:"lead"`
	LeadAvatar      string    `json:"leadAvatar" db:"avatar_url"`
}

type IncidentSummary struct {
	ID      int    `json:"id" db:"id"`
	Summary string `json:"summary" db:"summary"`
}

type IncidentStatus struct {
	ID      int    `json:"id" db:"id"`
	Status string `json:"status" db:"status"`
}

type IncidentSeverity struct {
	ID       int    `json:"id" db:"id"`
	Severity string `json:"severity" db:"severity"`
}

type IncidentType struct {
	ID      int    `json:"id" db:"id"`
	Type 	string `json:"type" db:"type"`
}

type IncidentRoles struct {
	ID      int		`json:"id" db:"id"`
	Lead    *int    `json:"lead" db:"lead"`
	QE      *int    `json:"qe" db:"qe"`
}

type IncidentOutput struct {
	ID                    int               `json:"id" db:"id"`
	Reference             *int              `json:"reference" db:"reference"`
	Status                string            `json:"status" db:"status"`
	Type                  string            `json:"type" db:"type"`
	Lead                  *int              `json:"lead" db:"lead"`
	Reporter              *int              `json:"reporter" db:"reporter"`
	QE                    *int              `json:"qe" db:"qe"`
	ReporterName          string            `json:"reporterName" db:"reporter_name"`
	LeadName              *string           `json:"leadName" db:"lead_name"`
	QeName                *string           `json:"QEName" db:"qe_name"`
	ReporterAvatar        string            `json:"reporterAvatar" db:"reporter_avatar"`
	LeadAvatar            *string           `json:"leadAvatar" db:"lead_avatar"`
	QEAvatar              *string           `json:"QEAvatar" db:"qe_avatar"`
	Title                 string            `json:"title" db:"title"`
	Summary               string            `json:"summary" db:"summary"`
	Severity              string            `json:"severity" db:"severity"`
	Impact                *string           `json:"impact" db:"impact"`
	PostMortem            *string           `json:"postMortem" db:"post_mortem"`
	ImpactStartedAt       *CustomTime        `json:"impactStartedAt" db:"impact_started_at"`
	ImpactStoppedAt       *CustomTime        `json:"impactStoppedAt" db:"impact_stopped_at"`
	ReportedAt            *CustomTime        `json:"reportedAt" db:"reported_at"`
	IdentifiedAt          *CustomTime        `json:"identifiedAt" db:"identified_at"`
	FixedAt               *CustomTime        `json:"fixedAt" db:"fixed_at"`
	ResolvedAt            *CustomTime        `json:"resolvedAt" db:"resolved_at"`
	DocumentationAt       *CustomTime        `json:"documentationAt" db:"documentation_at"`
	InReviewAt            *CustomTime        `json:"inReviewAt" db:"in_review_at"`
	ClosedAt              *CustomTime        `json:"closedAt" db:"closed_at"`
	AcceptedAt            *CustomTime        `json:"acceptedAt" db:"accepted_at"`
	DeclinedAt            *CustomTime        `json:"declinedAt" db:"declined_at"`
	MergedAt              *CustomTime        `json:"mergedAt" db:"merged_at"`
	CanceledAt            *CustomTime        `json:"canceledAt" db:"canceled_at"`
	TriagedBy             *CustomTime        `json:"triagedBy" db:"triaged_by"`
	Treatment             *string           `json:"treatment" db:"treatment"`
	Mitigator             *string           `json:"mitigator" db:"mitigator"`
	SlackChannel          *string           `json:"slackChannel" db:"slack_channel"`
	RelatedIncident       *int              `json:"relatedIncident" db:"related_incident"`
	IncidentSource        *string           `json:"incidentSource" db:"incident_source"`
	ThreadOnSlack         *string           `json:"threadOnSlack" db:"thread_on_slack"`
	CleanedUpAt           *CustomTime        `json:"cleanedUpAt" db:"cleaned_up_at"`
	MonitoredAt           *CustomTime        `json:"monitoredAt" db:"monitored_at"`
	InvestigatingAt       *CustomTime        `json:"investigatingAt" db:"investigating_at"`
	FixingAt              *CustomTime        `json:"fixingAt" db:"fixing_at"`
	MonitoringAt          *CustomTime        `json:"monitoringAt" db:"monitoring_at"`
	CleaningUpAt          *CustomTime        `json:"cleaningUpAt" db:"cleaning_up_at"`
	PostToStatusPage      *bool             `json:"postToStatusPage" db:"post_to_status_page"`
	DocumentedAt          *CustomTime        `json:"documentedAt" db:"documented_at"`
	ReviewedAt            *CustomTime        `json:"reviewedAt" db:"reviewed_at"`
	Category              *string           `json:"category" db:"category"`
	Products              []RelatedItem   	`json:"products" db:"-"`
	Areas                 []RelatedItem   	`json:"areas" db:"-"`
	Causes                []RelatedItem   	`json:"causes" db:"-"`
	FaultySystems         []RelatedItem    	`json:"faultySystems" db:"-"`
	PerformanceIndicators []RelatedItem  	`json:"performanceIndicators" db:"-"`
}

type IncidentCustomFieldsUpdate struct {
	ID                    int               `json:"id" db:"id"`
	Products              []int             `json:"products,omitempty"`
	Areas                 []int             `json:"areas,omitempty"`
	Causes                []int             `json:"causes,omitempty"`
	FaultySystems         []int             `json:"faultySystems,omitempty"`
	PerformanceIndicators []int             `json:"performanceIndicators,omitempty"`
	Impact                *string           `json:"impact,omitempty" db:"impact"`
	Treatment             *string           `json:"treatment,omitempty" db:"treatment"`
	Mitigator             *string           `json:"mitigator,omitempty" db:"mitigator"`
}
