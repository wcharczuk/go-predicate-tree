package predicate

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func makeTree() Node {
	return Or(
		And(
			Eval(T),
			Not(Eval(F)),
		),
		Or(
			Eval(T),
			Eval(F),
		),
	)
}

func TestPredicateTreeEvaluate(t *testing.T) {
	assert := assert.New(t)

	tree := makeTree()
	assert.True(tree.Evaluate())
}

func TestPredicateTreeString(t *testing.T) {
	assert := assert.New(t)

	tree := makeTree()
	assert.NotEmpty(tree.String())
	assert.Equal("OR(AND(EVAL(true), NOT(EVAL(false))), OR(EVAL(true), EVAL(false)))", tree.String())
}

func TestSerialize(t *testing.T) {
	assert := assert.New(t)

	tree := makeTree()

	serialized, err := Serialize(tree)
	assert.Nil(err)
	assert.NotEmpty(serialized)

	deserialized, err := Deserialize(serialized)

	assert.Nil(err)
	assert.NotNil(deserialized)
	assert.True(deserialized.Evaluate())
}
