package predicate

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestNot(t *testing.T) {
	assert := assert.New(t)

	assert.True(Not(Eval(F)).Evaluate())
	assert.False(Not(Eval(T)).Evaluate())
}
