package predicate

import "fmt"

const (
	NodeTypeEval = "eval"
)

func Eval(predicate Predicate) *EvalNode {
	return &EvalNode{
		PredicateType: predicate.Type(),
		Predicate:     predicate,
	}
}

type EvalNode struct {
	PredicateType string
	Predicate     Predicate
}

func (en EvalNode) Type() string {
	return NodeTypeEval
}

func (en EvalNode) Children() []Node {
	return nil
}

func (en EvalNode) AddChild(node Node) {}

func (en *EvalNode) Evaluate(args ...interface{}) bool {
	return en.Predicate.Evaluate(args...)
}

func (en *EvalNode) String() string {
	return fmt.Sprintf("EVAL(%s)", en.PredicateType)
}
