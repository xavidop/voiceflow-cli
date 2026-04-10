package transcript

type TranscriptInformation struct {
	ID            string `json:"id,omitempty"`
	ProjectID     string `json:"projectID,omitempty"`
	SessionID     string `json:"sessionID,omitempty"`
	EnvironmentID string `json:"environmentID,omitempty"`
	CreatedAt     string `json:"createdAt,omitempty"`
	UpdatedAt     string `json:"updatedAt,omitempty"`
}

// SearchTranscriptsRequest is the request body for the search transcripts API.
type SearchTranscriptsRequest struct {
	StartDate     string `json:"startDate,omitempty"`
	EndDate       string `json:"endDate,omitempty"`
	SessionID     string `json:"sessionID,omitempty"`
	EnvironmentID string `json:"environmentID,omitempty"`
}

// SearchTranscriptsResponse is the response from the search transcripts API.
type SearchTranscriptsResponse struct {
	Transcripts []TranscriptInformation `json:"transcripts"`
}
