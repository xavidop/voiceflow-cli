package analytics

import "time"

// Query represents the root structure of the JSON
type Query struct {
	Query []QueryItem `json:"query"`
}

// QueryItem represents an individual query with a name and a filter
type QueryItem struct {
	Name   string `json:"name"`
	Filter Filter `json:"filter"`
}

// Filter represents the filter criteria for the query
type Filter struct {
	ProjectID string    `json:"projectID"`
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
	Limit     int       `json:"limit,omitempty"`
}
