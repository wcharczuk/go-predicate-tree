package predicate

import "sync"

var (
	registryLock = sync.RWMutex{}
	registry     = map[string]Factory{
		PredicateTypeTrue:        func() Predicate { return &TruePredicate{} },
		PredicateTypeFalse:       func() Predicate { return &FalsePredicate{} },
		PredicateEquals:          func() Predicate { return &EqualsPredicate{} },
		PredicateStringContains:  func() Predicate { return &StringContainsPredicate{} },
		PredicateStringHasPrefix: func() Predicate { return &StringHasPrefixPredicate{} },
		PredicateStringHasSuffix: func() Predicate { return &StringHasSuffixPredicate{} },
	}
)

// Factory is a method that returns a bare predicate.
type Factory func() Predicate

// Register registers a predicate factory for a given predicate type.
func Register(predicateType string, factory Factory) {
	registryLock.Lock()
	defer registryLock.Unlock()

	registry[predicateType] = factory
}

// CreatePredicate creates a predicate from a predicate type name.
func CreatePredicate(predicateType string) Predicate {
	registryLock.RLock()
	defer registryLock.RUnlock()
	if factory, hasFactory := registry[predicateType]; hasFactory {
		return factory()
	}
	return nil
}
