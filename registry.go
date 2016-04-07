package predicate

import "sync"

var (
	registryLock = sync.RWMutex{}
	registry     = map[string]PredicateFactory{
		PredicateTypeTrue:  func() Predicate { return &TruePredicate{} },
		PredicateTypeFalse: func() Predicate { return &FalsePredicate{} },
	}
)

type PredicateFactory func() Predicate

func Register(predicateType string, factory PredicateFactory) {
	registryLock.Lock()
	defer registryLock.Unlock()

	registry[predicateType] = factory
}

func CreatePredicate(predicateType string) Predicate {
	registryLock.RLock()
	defer registryLock.RUnlock()
	if factory, hasFactory := registry[predicateType]; hasFactory {
		return factory()
	}
	return nil
}
