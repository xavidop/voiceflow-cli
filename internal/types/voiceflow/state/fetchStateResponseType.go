package state

type StackItem struct {
	ProgramID string `json:"programID"`
	DiagramID string `json:"diagramID"`
	NodeID   string `json:"nodeID"`
}

type State struct {
	Stack     []StackItem       `json:"stack"`
	Storage   map[string]any    `json:"storage"`
	Variables map[string]any    `json:"variables"`
}