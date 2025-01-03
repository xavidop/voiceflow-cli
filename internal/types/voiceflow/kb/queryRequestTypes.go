package kb

type Settings struct {
	Model       string  `json:"model,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	System      string  `json:"system,omitempty"`
}

type TagOperator struct {
	Operator string   `json:"operator,omitempty"`
	Items    []string `json:"items,omitempty"`
}

type Tags struct {
	Include             *TagOperator `json:"include,omitempty"`
	Exclude             *TagOperator `json:"exclude,omitempty"`
	IncludeAllTagged    bool         `json:"includeAllTagged,omitempty"`
	IncludeAllNonTagged bool         `json:"includeAllNonTagged,omitempty"`
}

type Query struct {
	ChunkLimit int       `json:"chunkLimit,omitempty"`
	Synthesis  bool      `json:"synthesis,omitempty"`
	Settings   *Settings `json:"settings,omitempty"`
	Tags       *Tags     `json:"tags,omitempty"`
	Question   string    `json:"question"`
}
