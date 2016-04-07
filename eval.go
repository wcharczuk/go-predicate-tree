package predicate

import (
	"fmt"

	"github.com/blendlabs/go-util"
)

const (
	NodeTypeEval = "eval"
)

func Eval(predicate Predicate) *EvalNode {
	return &EvalNode{
		id:            util.UUIDv4().ToShortString(),
		predicateType: predicate.Type(),
		predicate:     predicate,
	}
}

type EvalNode struct {
	id            string
	predicateType string
	predicate     Predicate
}

func (en EvalNode) ID() string {
	return en.id
}

func (en *EvalNode) SetID(id string) {
	en.id = id
}

func (en EvalNode) Type() string {
	return NodeTypeEval
}

func (en EvalNode) Children() []Node {
	return nil
}

func (en EvalNode) AddChild(node Node) {}

func (en EvalNode) RemoveChild(nodeID string) {}

func (en *EvalNode) Evaluate(args ...interface{}) bool {
	return en.predicate.Evaluate(args...)
}

func (en *EvalNode) String() string {
	return fmt.Sprintf("EVAL(%s)", en.predicateType)
}
