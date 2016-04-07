package predicate

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

type customPredicate struct {
	StateValue string `json:"state_value"`
}

func (cp customPredicate) Type() string {
	return "custom_predicate"
}

func (cp customPredicate) Evaluate() bool {
	return len(cp.StateValue) != 0
}

func NewCustomPredicate(stateValue string) Predicate {
	return &customPredicate{
		StateValue: stateValue,
	}
}

func TestSerializerWithCustomPredicate(t *testing.T) {
	assert := assert.New(t)

	Register("custom_predicate", func() Predicate {
		return &customPredicate{}
	})

	tree := Or(Eval(NewCustomPredicate("things")), Eval(NewCustomPredicate("other things")))
	serialized, err := Serialize(tree)
	assert.Nil(err)
	assert.NotEmpty(serialized)
	assert.True(tree.Evaluate())

	deserialized, err := Deserialize(serialized)
	assert.Nil(err)
	assert.NotNil(deserialized)
	assert.True(deserialized.Evaluate())

	typed, isTyped := deserialized.(*OrNode)
	assert.True(isTyped)
	assert.NotNil(typed)
	assert.NotEmpty(typed.Children())

	evalNode, isEvalNode := typed.Children()[0].(*EvalNode)
	assert.True(isEvalNode)
	assert.NotNil(evalNode)

	predicate, isPredicate := evalNode.Predicate.(*customPredicate)
	assert.True(isPredicate)
	assert.NotNil(predicate)

	assert.NotEmpty(predicate.StateValue)
}
