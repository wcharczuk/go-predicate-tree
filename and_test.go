package predicate

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestAnd(t *testing.T) {
	assert := assert.New(t)

	assert.True(And(Eval(T), Eval(T)).Evaluate())
	assert.False(And(Eval(T), Eval(F)).Evaluate())
	assert.False(And(Eval(F), Eval(T)).Evaluate())
	assert.False(And(Eval(F), Eval(F)).Evaluate())
}
