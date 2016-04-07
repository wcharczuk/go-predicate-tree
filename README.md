Go Predicate
============

`go-predicate` is a simple package that implements a basic boolean evaluation tree that can be serialized to json.

It allows for custom predicates to be serialized with the tree, preserving relevant state for later evaluation.

##Example

Lets say you have the following functions:
```golang

func TheFooIsBar(foo string) bool {
    return foo == "bar"
}

func TheBarIsFoo(bar string) bool {
    return bar == "foo"
}

func Validate(foo, bar) bool {
    return TheFooIsBar(foo) && TheBarIsFoo(bar)
}

```

You could represent this as a predicate tree:

```golang
foo := "something"
bar := "somethingElse"

tree := And(
        Eval(ExpectedValue(foo, "bar")),
        Eval(ExpectedValue(bar, "foo")),
    )
```

This assumes the following predicate was implemented:

```golang
type expectedValuePredicate struct {
    Expected string `json:"expected"`
    Value string `json:"value"`
}

func (evp *expectedValuePredicate) Type() {
    return "expected_value"
}

func (evp *expectedValuePredicate) Evaluate() bool {
    return evp.Expected == evp.Value
}
```

So that's fine and feels super complicated to do some basic validation.

The power of this comes from serialization. If we register the predicate:

```golang
predicate.Register("expected_value", func() Predicate { return &expectedValue{} })
```

We could then do the following:

```golang
blob := Serialize(
        And (
            Eval(ExpectedValue(foo, "bar")),
            Eval(ExpectedValue(bar, "foo")),
        )
    )

tree := predicate.Deserialize(blob)
fmt.Printf("%v\n", tree.Evaluate())
```