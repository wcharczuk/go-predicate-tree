package predicate

import "reflect"

const (
	// PredicateTypeTrue is the type name for the true predicate.
	PredicateTypeTrue = "true"

	// PredicateTypeFalse is the type name for the false predicate.
	PredicateTypeFalse = "false"

	// PredicateEquals is the type name for the equals predicate.
	PredicateEquals = "equals"
)

var (
	// T is a prebuilt predicate that returns true.
	T = &TruePredicate{}

	// F is a prebuilt predicate that returns false.
	F = &FalsePredicate{}
)

// Predicate is a value provider for the Eval node.
type Predicate interface {
	// Type is the predicate registry key for the predicate.
	Type() string

	// Evaluate runs the predicate.
	Evaluate(args ...interface{}) bool
}

// TruePredicate is a pre-built predicate that returns true
type TruePredicate struct{}

// Evaluate returns true
func (tp *TruePredicate) Evaluate(args ...interface{}) bool {
	return true
}

// Type returns the predicate type.
func (tp *TruePredicate) Type() string {
	return PredicateTypeTrue
}

// FalsePredicate is a pre-built predicate that returns false.
type FalsePredicate struct{}

// Evaluate returns false.
func (fp *FalsePredicate) Evaluate(args ...interface{}) bool {
	return false
}

// Type returns the predicate type.
func (fp *FalsePredicate) Type() string {
	return PredicateTypeFalse
}

// Equals is a simple predicate that evaluates if a value equals something.
func Equals(value interface{}) Predicate {
	return &EqualsPredicate{
		Value: value,
	}
}

// EqualsPredicate is a pre-built predicate that tests equality.
type EqualsPredicate struct {
	Value interface{} `json:"value"`
}

// Evaluate returns a bool
func (ep *EqualsPredicate) Evaluate(args ...interface{}) bool {
	if len(args) == 0 {
		return false
	}

	actual := args[0]
	return reflect.DeepEqual(ep.Value, actual)
}

// Type returns the type name.
func (ep *EqualsPredicate) Type() string {
	return PredicateEquals
}
