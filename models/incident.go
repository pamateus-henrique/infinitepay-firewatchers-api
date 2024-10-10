package models

import "time"

type IncidentInput struct {
	Title           string     `json:"title" db:"title"`
	Type            int        `json:"type" db:"type"`
	Severity        string     `json:"severity" db:"severity"`
	Summary         string     `json:"summary" db:"summary"`
	Impact          *string    `json:"impact,omitempty" db:"impact"`
	Source          *string    `json:"source,omitempty" db:"incident_source"`
	Products        []int      `json:"products,omitempty" db:"product_id"`
	Areas           []int      `json:"areas,omitempty" db:"area_id"`
	Indicators      []int      `json:"indicators,omitempty" db:"indicator_id"`
	ImpactStartedAt *time.Time `json:"started_at" db:"impact_started_at"`
	SlackThread     *string    `json:"slack_thread,omitempty" db:"slack_thread"`
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
	ID      int    `json:"id" db:"id"`
	Severity string `json:"severity" db:"severity"`
}


type IncidentOutput struct {
	ID                    int               `json:"id" db:"id"`
	Reference             *int              `json:"reference" db:"reference"`
	Status                string            `json:"status" db:"status"`
	Type                  int               `json:"type" db:"type"`
	Lead                  *int              `json:"lead" db:"lead"`
	Reporter              *int              `json:"reporter" db:"reporter"`
	QE                    *int              `json:"qe" db:"qe"`
	ReporterName          string            `json:"reporterName" db:"reporter_name"`
	LeadName              *string           `json:"leadName" db:"lead_name"`
	QeName                *string           `json:"qeName" db:"qe_name"`
	ReporterAvatar        string            `json:"reporterAvatar" db:"reporter_avatar"`
	LeadAvatar            *string           `json:"leadAvatar" db:"lead_avatar"`
	QEAvatar              *string           `json:"QEAvatar" db:"qe_avatar"`
	Title                 string            `json:"title" db:"title"`
	Summary               string            `json:"summary" db:"summary"`
	Severity              string            `json:"severity" db:"severity"`
	Impact                *string           `json:"impact" db:"impact"`
	PostMortem            *string           `json:"postMortem" db:"post_mortem"`
	ImpactStartedAt       *time.Time        `json:"impactStartedAt" db:"impact_started_at"`
	ImpactStoppedAt       *time.Time        `json:"impactStoppedAt" db:"impact_stopped_at"`
	ReportedAt            *time.Time        `json:"reportedAt" db:"reported_at"`
	IdentifiedAt          *time.Time        `json:"identifiedAt" db:"identified_at"`
	FixedAt               *time.Time        `json:"fixedAt" db:"fixed_at"`
	ResolvedAt            *time.Time        `json:"resolvedAt" db:"resolved_at"`
	DocumentationAt       *time.Time        `json:"documentationAt" db:"documentation_at"`
	InReviewAt            *time.Time        `json:"inReviewAt" db:"in_review_at"`
	ClosedAt              *time.Time        `json:"closedAt" db:"closed_at"`
	AcceptedAt            *time.Time        `json:"acceptedAt" db:"accepted_at"`
	DeclinedAt            *time.Time        `json:"declinedAt" db:"declined_at"`
	MergedAt              *time.Time        `json:"mergedAt" db:"merged_at"`
	CanceledAt            *time.Time        `json:"canceledAt" db:"canceled_at"`
	TriagedBy             *time.Time        `json:"triagedBy" db:"triaged_by"`
	Treatment             *string           `json:"treatment" db:"treatment"`
	Mitigator             *string           `json:"mitigator" db:"mitigator"`
	SlackChannel          *string           `json:"slackChannel" db:"slack_channel"`
	RelatedIncident       *int              `json:"relatedIncident" db:"related_incident"`
	IncidentSource        *string           `json:"incidentSource" db:"incident_source"`
	ThreadOnSlack         *string           `json:"threadOnSlack" db:"thread_on_slack"`
	CleanedUpAt           *time.Time        `json:"cleanedUpAt" db:"cleaned_up_at"`
	MonitoredAt           *time.Time        `json:"monitoredAt" db:"monitored_at"`
	InvestigatingAt       *time.Time        `json:"investigatingAt" db:"investigating_at"`
	FixingAt              *time.Time        `json:"fixingAt" db:"fixing_at"`
	MonitoringAt          *time.Time        `json:"monitoringAt" db:"monitoring_at"`
	CleaningUpAt          *time.Time        `json:"cleaningUpAt" db:"cleaning_up_at"`
	PostToStatusPage      *bool             `json:"postToStatusPage" db:"post_to_status_page"`
	DocumentedAt          *time.Time        `json:"documentedAt" db:"documented_at"`
	ReviewedAt            *time.Time        `json:"reviewedAt" db:"reviewed_at"`
	Category              *string           `json:"category" db:"category"`
	Products              map[int]string    `json:"products" db:"-"`
	Areas                 map[int]string    `json:"areas" db:"-"`
	Causes                map[int]string    `json:"causes" db:"-"`
	FaultySystems         map[int]string    `json:"faultySystems" db:"-"`
	PerformanceIndicators map[int]string    `json:"performanceIndicators" db:"-"`
}