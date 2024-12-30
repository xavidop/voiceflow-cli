package voiceflow

type InteractionResponse struct {
	Type    string                 `json:"type"`
	Payload map[string]interface{} `json:"payload"`
}

// MAIN TYPES

type Speak struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Src     string `json:"src,omitempty"`
}

type Audio struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Src     string `json:"src,omitempty"`
}

type Visual struct {
	Type    string        `json:"type"`
	Payload VisualPayload `json:"payload"`
}

type CardV2 struct {
	Type    string        `json:"type"`
	Time    int64         `json:"time"`
	Payload CardV2Payload `json:"payload"`
}

type NoReply struct {
	Type    string         `json:"type"`
	Time    int64          `json:"time"`
	Payload NoReplyPayload `json:"payload"`
}

type Carousel struct {
	Type    string          `json:"type"`
	Time    int64           `json:"time"`
	Payload CarouselPayload `json:"payload"`
}

type End struct {
	Type    string      `json:"type"`
	Time    int64       `json:"time"`
	Payload interface{} `json:"payload"` // Can be null, so use interface{} to allow nil
}

// Specific payload types
type VisualPayload struct {
	VisualType       string     `json:"visualType"`
	Image            string     `json:"image"`
	Dimensions       Dimensions `json:"dimensions"`
	CanvasVisibility string     `json:"canvasVisibility"`
}

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type CardV2Payload struct {
	ImageUrl    string      `json:"imageUrl"`
	Description Description `json:"description"`
	Buttons     []Button    `json:"buttons"`
	Title       string      `json:"title"`
}

type Description struct {
	Slate []Slate `json:"slate"`
	Text  string  `json:"text"`
}

type Slate struct {
	Children []Child `json:"children"`
}

type Child struct {
	Text string `json:"text"`
}

type Button struct {
	Name    string        `json:"name"`
	Request ButtonRequest `json:"request"`
}

type ButtonRequest struct {
	Type    string               `json:"type"`
	Payload ButtonRequestPayload `json:"payload"`
}

type ButtonRequestPayload struct {
	Actions []interface{} `json:"actions"` // Adjust the type if `actions` has a specific structure
	Label   string        `json:"label"`
}

type NoReplyPayload struct {
	Timeout int `json:"timeout"` // Timeout in seconds
}

type CarouselPayload struct {
	Layout string `json:"layout"`
	Cards  []Card `json:"cards"`
}

type Card struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Description Description `json:"description"`
	ImageUrl    string      `json:"imageUrl"`
	Buttons     []Button    `json:"buttons"`
}
