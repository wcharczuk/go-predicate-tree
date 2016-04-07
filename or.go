package predicate

import (
	"fmt"
	"strings"

	"github.com/blendlabs/go-util"
)

const (
	NodeTypeOr = "or"
)

func Or(nodes ...Node) *OrNode {
	return &OrNode{
		id:       util.UUIDv4().ToShortString(),
		children: append([]Node{}, nodes...),
	}
}

type OrNode struct {
	id       string
	children []Node
}

func (or OrNode) ID() string {
	return or.id
}

func (on *OrNode) SetID(id string) {
	on.id = id
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

func (on *OrNode) RemoveChild(nodeID string) {
	var newChildren []Node
	for _, c := range on.children {
		if c.ID() != nodeID {
			newChildren = append(newChildren, c)
		}
	}
	on.children = newChildren
}

func (on *OrNode) Evaluate(args ...interface{}) bool {
	result := false
	for _, n := range on.children {
		result = result || n.Evaluate(args...)
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
