package predicate

import (
	"fmt"
	"strings"

	"github.com/blendlabs/go-util"
)

const (
	NodeTypeAnd = "and"
)

func And(nodes ...Node) *AndNode {
	return &AndNode{
		id:       util.UUIDv4().ToShortString(),
		children: append([]Node{}, nodes...),
	}
}

type AndNode struct {
	id       string
	children []Node
}

func (an AndNode) ID() string {
	return an.id
}

func (an *AndNode) SetID(id string) {
	an.id = id
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

func (an *AndNode) RemoveChild(nodeID string) {
	var newChildren []Node
	for _, c := range an.children {
		if c.ID() != nodeID {
			newChildren = append(newChildren, c)
		}
	}
	an.children = newChildren
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
