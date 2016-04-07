package predicate

import (
	"fmt"

	"github.com/blendlabs/go-util"
)

const (
	NodeTypeNot = "not"
)

func Not(node Node) *NotNode {
	return &NotNode{
		id:    util.UUIDv4().ToShortString(),
		child: node,
	}
}

type NotNode struct {
	id    string
	child Node
}

func (nn NotNode) ID() string {
	return nn.id
}

func (nn *NotNode) SetID(id string) {
	nn.id = id
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

func (nn *NotNode) RemoveChild(nodeID string) {
	nn.child = nil
}

func (nn *NotNode) Evaluate(args ...interface{}) bool {
	return !nn.child.Evaluate(args...)
}

func (nn *NotNode) String() string {
	return fmt.Sprintf("NOT(%s)", nn.child.String())
}
