package predicate

import "fmt"

const (
	NodeTypeNot = "not"
)

func Not(node Node) *NotNode {
	return &NotNode{
		child: node,
	}
}

type NotNode struct {
	child Node `json:"child"`
}

func (nn NotNode) Type() string {
	return NodeTypeNot
}

func (nn NotNode) Children() []Node {
	return []Node{nn.child}
}

func (nn *NotNode) AddChild(node Node) {
	nn.child = node
}

func (nn *NotNode) Evaluate() bool {
	return !nn.child.Evaluate()
}

func (nn *NotNode) String() string {
	return fmt.Sprintf("NOT(%s)", nn.child.String())
}
