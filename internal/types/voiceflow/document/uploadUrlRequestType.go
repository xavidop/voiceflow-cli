package document

// Data represents the nested "data" structure
type Data struct {
	Type string `json:"type"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Payload represents the top-level structure
type URLDocument struct {
	Data Data `json:"data"`
}
