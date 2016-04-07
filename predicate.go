package predicate

const (
	// PredicateTypeTrue is the type name for the true predicate.
	PredicateTypeTrue = "true"

	// PredicateTypeFalse is the type name for the false predicate.
	PredicateTypeFalse = "false"
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
	Evaluate() bool
}

// TruePredicate is a pre-built predicate that returns true
type TruePredicate struct{}

// Evaluate returns true
func (tp *TruePredicate) Evaluate() bool {
	return true
}

// Type returns the predicate type.
func (tp *TruePredicate) Type() string {
	return PredicateTypeTrue
}

// FalsePredicate is a pre-built predicate that returns false.
type FalsePredicate struct{}

// Evaluate returns false.
func (fp *FalsePredicate) Evaluate() bool {
	return false
}

// Type returns the predicate type.
func (fp *FalsePredicate) Type() string {
	return PredicateTypeFalse
}
