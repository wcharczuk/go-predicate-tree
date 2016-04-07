package predicate

import (
	"fmt"
	"strings"
)

const (
	NodeTypeAnd = "and"
)

func And(nodes ...Node) *AndNode {
	return &AndNode{
		children: append([]Node{}, nodes...),
	}
}

type AndNode struct {
	children []Node
}

func (an AndNode) Type() string {
	return NodeTypeAnd
}

func (an AndNode) Children() []Node {
	return an.children
}

func (an *AndNode) AddChild(node Node) {
	an.children = append(an.children, node)
}

func (an *AndNode) Evaluate(args ...interface{}) bool {
	result := true
	for _, n := range an.children {
		result = result && n.Evaluate(args...)
	}
	return result
}

func (an *AndNode) String() string {
	childStrings := []string{}
	for _, n := range an.children {
		childStrings = append(childStrings, n.String())
	}
	return fmt.Sprintf("AND(%s)", strings.Join(childStrings, ", "))
}
