package predicate

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestRegister(t *testing.T) {
	assert := assert.New(t)

	Register("true_predicate", func() Predicate { return &TruePredicate{} })
	Register("false_predicate", func() Predicate { return &FalsePredicate{} })

	assert.NotNil(CreatePredicate("true_predicate"))
	assert.NotNil(CreatePredicate("false_predicate"))
}
