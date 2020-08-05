package model

// NodeType is type for usecase unit object
type NodeType string

const (
	// Title is usecase main title
	Title = "Title"
	// Context is unit of usecase context name
	Context = "Context"
	// Behavior is unit of usecase behavior such as Request, Response and ViewModel
	Behavior = "Behavior"
)

// Node is usecase unit object
type Node struct {
	Type     NodeType
	Name     string
	Children []Node
}
