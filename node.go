package predicate

// Node is a node in the predicate tree.
type Node interface {
	Type() string
	AddChild(node Node)
	Children() []Node
	Evaluate() bool
	String() string
}
