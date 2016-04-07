package predicate

// Node is a node in the predicate tree.
type Node interface {
	ID() string
	SetID(id string)
	Type() string
	AddChild(node Node)
	RemoveChild(nodeID string)
	Children() []Node
	String() string
	Evaluate(args ...interface{}) bool
}
