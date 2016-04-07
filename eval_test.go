package predicate

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestEval(t *testing.T) {
	assert := assert.New(t)

	assert.True(Eval(T).Evaluate())
	assert.False(Eval(F).Evaluate())
}
