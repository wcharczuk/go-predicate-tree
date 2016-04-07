package predicate

// Node is a node in the predicate tree.
type Node interface {
	Type() string
	AddChild(node Node)
	Children() []Node
	String() string

	Evaluate(args ...interface{}) bool
}
