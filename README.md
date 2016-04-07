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
    return TheFooIsBar(foo) || TheBarIsFoo(bar)
}

```

You could represent this as a predicate tree:

```golang
foo := "something"
bar := "somethingElse"

tree := Or(
        Eval(Equals("bar")),
        Eval(Equals("foo")),
    )
```

Note: This uses a pre-built predicate ("Equals") that tests if the first argument passed to `Evaluate(...)` is equal to the given value.

The power of this comes from serialization.

We could then do the following:

```golang
blob := Serialize(
        Or(
            Eval(Equals("bar")),
            Eval(Equals("foo")),
        )
    )

tree := predicate.Deserialize(blob)
fmt.Printf("%v\n", tree.Evaluate("foo")) // "true"
```

This does a round trip serialize => deserialize, preserving the 'state' of the predicates, and allowing us to evaluate the tree after deserialization.