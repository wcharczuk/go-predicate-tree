package predicate

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	// PredicateTypeTrue is the type name for the true predicate.
	PredicateTypeTrue = "true"

	// PredicateTypeFalse is the type name for the false predicate.
	PredicateTypeFalse = "false"

	// PredicateEquals is the type name for the equals predicate.
	PredicateEquals = "equals"

	// PredicateStringContains is the type name for the string contains predicate.
	PredicateStringContains = "contains"

	// PredicateStringHasPrefix is the type name for the string contains predicate.
	PredicateStringHasPrefix = "has_prefix"

	// PredicateStringHasSuffix is the type name for the string contains predicate.
	PredicateStringHasSuffix = "has_suffix"
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

// StringContains returns a predicate for strings.Contains.
func StringContains(value string) Predicate {
	return &StringContainsPredicate{
		Value: value,
	}
}

// StringContains is the predicate for strings.Contains.
type StringContainsPredicate struct {
	Value string `json:"value"`
}

func (scp *StringContainsPredicate) Type() string {
	return PredicateStringContains
}

func (scp *StringContainsPredicate) Evaluate(args ...interface{}) bool {
	if len(args) == 0 {
		return false
	}

	result := true
	for _, v := range args {
		if typed, isString := v.(string); isString {
			result = result && strings.Contains(typed, scp.Value)
		} else {
			typed := fmt.Sprintf("%v", v)
			result = result && strings.Contains(typed, scp.Value)
		}
	}

	return result
}

// HasPrefix is the predicate for strings.HasPrefix.
func HasPrefix(value string) Predicate {
	return &StringHasPrefixPredicate{
		Value: value,
	}
}

// StringHasPrefixPredicate is the predicate for strings.HasPrefix.
type StringHasPrefixPredicate struct {
	Value string `json:"value"`
}

// Type returns the predicate type.
func (shp *StringHasPrefixPredicate) Type() string {
	return PredicateStringHasPrefix
}

// Evaluate evaluates the predicate.s
func (shp *StringHasPrefixPredicate) Evaluate(args ...interface{}) bool {
	if len(args) == 0 {
		return false
	}

	result := true
	for _, v := range args {
		if typed, isString := v.(string); isString {
			result = result && strings.HasPrefix(typed, shp.Value)
		} else {
			typed := fmt.Sprintf("%v", v)
			result = result && strings.HasPrefix(typed, shp.Value)
		}
	}

	return result
}

// HasSuffix returns a predicate for strings.HasSuffix.
func HasSuffix(value string) Predicate {
	return &StringHasSuffixPredicate{
		Value: value,
	}
}

// StringHasSuffixPredicate is the predicate for strings.HasSuffix.
type StringHasSuffixPredicate struct {
	Value string `json:"value"`
}

// Type returns the predicate type.
func (shp *StringHasSuffixPredicate) Type() string {
	return PredicateStringHasSuffix
}

// Evaluate evaluates the predicate.s
func (shp *StringHasSuffixPredicate) Evaluate(args ...interface{}) bool {
	if len(args) == 0 {
		return false
	}

	result := true
	for _, v := range args {
		if typed, isString := v.(string); isString {
			result = result && strings.HasSuffix(typed, shp.Value)
		} else {
			typed := fmt.Sprintf("%v", v)
			result = result && strings.HasSuffix(typed, shp.Value)
		}
	}

	return result
}
