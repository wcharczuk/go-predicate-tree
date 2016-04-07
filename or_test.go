package predicate

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestOr(t *testing.T) {
	assert := assert.New(t)

	assert.True(Or(Eval(T), Eval(T)).Evaluate())
	assert.True(Or(Eval(T), Eval(F)).Evaluate())
	assert.True(Or(Eval(F), Eval(T)).Evaluate())
	assert.False(Or(Eval(F), Eval(F)).Evaluate())
}
