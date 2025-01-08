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
	StartTime CustomTime `json:"startTime,omitempty"`
	EndTime   CustomTime `json:"endTime,omitempty"`
	Limit     int       `json:"limit,omitempty"`
}

// CustomTime is a wrapper around time.Time that formats to ISO-8601 with milliseconds
type CustomTime struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.Format("2006-01-02T15:04:05.000Z") + `"`), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02T15:04:05.000Z", s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}
