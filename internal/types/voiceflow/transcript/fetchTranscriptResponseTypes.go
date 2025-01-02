package transcript

type Annotation struct {
	ID struct {
		UtteranceAddedCount int    `json:"utteranceAddedCount"`
		UtteranceAddedTo    string `json:"utteranceAddedTo"`
	} `json:"_id"`
}

type TranscriptInformation struct {
	ProjectID   string     `json:"projectID,omitempty"`
	SessionID   string     `json:"sessionID,omitempty"`
	Browser     string     `json:"browser,omitempty"`
	Device      string     `json:"device,omitempty"`
	OS          string     `json:"os,omitempty"`
	ReportTags  []string   `json:"reportTags,omitempty"`
	Unread      bool       `json:"unread,omitempty"`
	Annotations Annotation `json:"annotations,omitempty"`
	CreatorID   int        `json:"creatorID,omitempty"`
	CreatedAt   string     `json:"createdAt,omitempty"`
	UpdatedAt   string     `json:"updatedAt,omitempty"`
	Name        string     `json:"name,omitempty"`
	Image       string     `json:"image,omitempty"`
	ID          string     `json:"_id,omitempty"`
}
