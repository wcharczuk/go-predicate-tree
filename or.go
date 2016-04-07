package predicate

import (
	"fmt"
	"strings"
)

const (
	NodeTypeOr = "or"
)

func Or(nodes ...Node) *OrNode {
	return &OrNode{
		children: append([]Node{}, nodes...),
	}
}

type OrNode struct {
	children []Node `json:"children"`
}

func (on OrNode) Type() string {
	return NodeTypeOr
}

func (on OrNode) Children() []Node {
	return on.children
}

func (on *OrNode) AddChild(node Node) {
	on.children = append(on.children, node)
}

func (on *OrNode) Evaluate() bool {
	result := false
	for _, n := range on.children {
		result = result || n.Evaluate()
	}
	return result
}

func (on *OrNode) String() string {
	childStrings := []string{}
	for _, n := range on.children {
		childStrings = append(childStrings, n.String())
	}
	return fmt.Sprintf("OR(%s)", strings.Join(childStrings, ", "))
}
