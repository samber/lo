# lo

[![Build Status](https://secure.travis-ci.org/samber/lo.svg?branch=master)](http://travis-ci.org/samber/lo)
[![GoDoc](https://godoc.org/github.com/samber/lo?status.svg)](https://pkg.go.dev/github.com/samber/lo)
[![Go report](https://goreportcard.com/badge/github.com/samber/lo)](https://goreportcard.com/report/github.com/samber/lo)

`lo` is a Lodash-style Go library based on Go 1.18+ Generics.

This project have started as an experiment to discover generics implementation. It may look like Lodash in some aspects. I used to code with the awesome [go-funk](https://github.com/thoas/go-funk) package, but it uses reflection and therefore is not typesafe.

As expected, my benchmarks demonstrate that generics will be much faster than reflect-based implementations.

In the future, some of these helpers will be available in the Go standard library (under package names "slice" and "maps").

## Why this name?

I wanted a short name, similar to "Lodash", and no Go package currently use this name.

## Install

```sh
go get github.com/samber/lo
```

## Usage

You can import ``lo`` using a basic statement:

```go
import "github.com/samber/lo"
```

Then use one of the helpers behind:

```go
names := lo.Uniq[string]([]string{"Samuel", "Marc", "Samuel"})
// names == []string{"Samuel", "Marc"}
```

## Spec

GoDoc: [https://godoc.org/github.com/samber/lo](https://godoc.org/github.com/samber/lo)

Supported helpers for slices:

- Filter
- Map
- Reduce
- ForEach
- Uniq
- UniqBy
- GroupBy
- Chunk
- Flatten
- Shuffle
- Reverse
- Fill
- ToMap

Supported helpers for maps:

- Keys
- Values
- Entries
- FromEntries
- Assign (maps merge)

Supported intersection helpers:

- Contains
- Every
- Some
- Intersect
- Difference

Supported search helpers:

- IndexOf
- LastIndexOf
- Find
- Min
- Max
- Last
- Nth

Other functional programming helpers:

- Ternary (1 line if/else statement)
- If / ElseIf / Else
- Switch / Case / Default
- ToPtr
- ToSlicePtr

Constraints:

- Clonable
- Ordered

### Map

Manipulates a slice and transforms it to a slice of another type:

```go
lo.Map[int64, string]([]int64{1, 2, 3, 4}, func(x int64) string {
    return strconv.FormatInt(x, 10)
})
// []string{"1", "2", "3", "4"}
```

### Filter

Iterates over elements of collection, returning an array of all elements predicate returns truthy for.

```go
odd := lo.Filter[int]([]int{1, 2, 3, 4}, func(x int) bool {
    return x%2 == 0
})
// []int{2, 4}
```

### Contains

Returns true if an element is present in a collection.

```go
present := lo.Contains[int]([]int{0, 1, 2, 3, 4, 5}, 5)
// true
```

### Reduce

Reduces collection to a value which is the accumulated result of running each element in collection through accumulator, where each successive invocation is supplied the return value of the previous.

```go
sum := lo.Reduce[int, int]([]int{1, 2, 3, 4}, func(agg int, item int) int {
    return agg + item
}, 0)
// 10
```

### ForEach

Iterates over elements of collection and invokes iteratee for each element.

```go
odd := lo.Filter[string]([]string{"hello", "world"}, func(x string) {
    println(x)
})
// prints "hello\nworld"
```

### Uniq

Returns a duplicate-free version of an array, in which only the first occurrence of each element is kept. The order of result values is determined by the order they occur in the array.

```go
uniqValues := lo.Uniq[int]([]int{1, 2, 2, 1})
// []int{1, 2}
```

### UniqBy

Returns a duplicate-free version of an array, in which only the first occurrence of each element is kept. The order of result values is determined by the order they occur in the array. It accepts `iteratee` which is invoked for each element in array to generate the criterion by which uniqueness is computed.

```go
uniqValues := lo.UniqBy[int, int]([]int{0, 1, 2, 3, 4, 5}, func (i int) int {
    return i%3
})
// []int{0, 1, 2}
```

### GroupBy

Returns an object composed of keys generated from the results of running each element of collection through iteratee.

```go
groups := GroupBy[int, int]([]int{0, 1, 2, 3, 4, 5}, func (i int) int {
    return i%3
})
// map[int][]int{0: []int{0, 3}, 1: []int{1, 4}, 2: []int{2, 5}}
```

### Chunk

Returns an array of elements split into groups the length of size. If array can't be split evenly, the final chunk will be the remaining elements.

```go
lo.Chunk[int]([]int{0, 1, 2, 3, 4, 5}, 2)
// [][]int{{0, 1}, {2, 3}, {4, 5}}

lo.Chunk[int]([]int{0, 1, 2, 3, 4, 5, 6}, 2)
// [][]int{{0, 1}, {2, 3}, {4, 5}, {6}}

lo.Chunk[int]([]int{}, 2)
// [][]int{}

lo.Chunk[int]([]int{0}, 2)
// [][]int{{0}}
```

### Flatten

Returns an array a single level deep.

```go
flat := lo.Flatten[int]([][]int{{0, 1}, {2, 3, 4, 5}})
// []int{0, 1, 2, 3, 4, 5}
```

### Shuffle

Returns an array of shuffled values.

```go
randomOrder := lo.Shuffle[int]([]int{0, 1, 2, 3, 4, 5})
// []int{0, 1, 2, 3, 4, 5}
```

### Reverse

Reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.

```go
reverseOder := lo.Reverse[int]([]int{0, 1, 2, 3, 4, 5})
// []int{5, 4, 3, 2, 1, 0}
```

### Fill

Fills elements of array with `initial` value.

```go
type foo struct {
	bar string
}

func (f foo) Clone() foo {
	return foo{f.bar}
}

initializedSlice := lo.Fill[foo]([]foo{foo{"a"}, foo{"a"}}, foo{"b"})
// []foo{foo{"b"}, foo{"b"}}
```

### ToMap

Transforms a slice or an array of structs to a map based on a pivot callback.

```go
m := lo.ToMap[int, string]([]string{"a", "aa", "aaa"}, func (str string) int {
    return len(str)
})
// map[int]string{1: "a", 2: "aa", 3: "aaa"}
```

### Keys

Creates an array of the map keys.

```go
keys := lo.Keys[string, int](map[string]int{"foo": 1, "bar": 2})
// []string{"bar", "foo"}
```

### Values

Creates an array of the map values.

```go
values := lo.Values[string, int](map[string]int{"foo": 1, "bar": 2})
// []int{1, 2}
```

### Entries

Transforms a map into array of key/value pairs.

```go
entries := lo.Entries[string, int](map[string]int{"foo": 1, "bar": 2})
// []lo.Entry[string, int]{
//     {
//         Key: "foo",
//         Value: 1,
//     },
//     {
//         Key: "bar",
//         Value: 2,
//     },
// }
```

### FromEntries

Transforms an array of key/value pairs into a map.

```go
m := lo.FromEntries[string, int]([]lo.Entry[string, int]{
    {
        Key: "foo", 
        Value: 1,
    },
    {
        Key: "bar",
        Value: 2,
    },
})
// map[string]int{"foo": 1, "bar": 2}
```

### Assign

Merges multiple maps from left to right.

```go
mergedMaps := lo.Assign[string, int](
    map[string]int{"a": 1, "b": 2},
    map[string]int{"b": 3, "c": 4},
)
// map[string]int{"a": 1, "b": 3, "c": 4}
```

### Every

Returns true if all elements of a subset are contained into a collection.

```go
ok := lo.Every[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
// true

ko := lo.Every[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
// false
```

### Some

Returns true if at least 1 element of a subset is contained into a collection.

```go
ok := lo.Some[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
// true

ko := lo.Some[int]([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
// false
```

### Intersect

Returns the intersection between two collections.

```go
result1 := lo.Intersect[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
// []int{0, 2}

result2 := lo.Intersect[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 6}
// []int{0}

result3 := lo.Intersect[int]([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
// []int{}
```

### Difference

Returns the difference between two collections.

- The first value is the collection of element absent of list2.
- The second value is the collection of element absent of list1.

```go
left, right := lo.Difference[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 6})
// []int{1, 3, 4, 5}, []int{6}

left, right := Difference[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4, 5})
// []int{}, []int{}
```

### IndexOf

Returns the index at which the first occurrence of a value is found in an array or return -1 if the value cannot be found.

```go
found := lo.IndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 2)
// 2

notFound := lo.IndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 6)
// -1
```

### LastIndex

Returns the index at which the last occurrence of a value is found in an array or return -1 if the value cannot be found.

```go
found := lo.LastIndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 2)
// 4

notFound := lo.LastIndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 6)
// -1
```

### Find

Search an element in a slice based on a predicate. It returns element and true if element was found.

```go
str, ok := lo.Find[string]([]string{"a", "b", "c", "d"}, func(i string) bool {
    return i == "b"
})
// "b", true

str, ok := lo.Find[string]([]string{"foobar"}, func(i string) bool {
    return i == "b"
})
// "", false
```

### Min

Search the minimum value of a collection.

```go
min := lo.Min[int]([]int{1, 2, 3})
// 1

min := lo.Min[int]([]int{})
// 0
```

### Max

Search the maximum value of a collection.

```go
min := lo.Max[int]([]int{1, 2, 3})
// 3

min := lo.Max[int]([]int{})
// 0
```

### Last

Returns the last element of a collection or panics if empty.

```go
last := lo.Last[int]([]int{1, 2, 3})
// 3
```

### Nth

Returns the element at index `nth` of collection. If `nth` is negative, the nth element from the end is returned.

```go
nth := lo.Nth[int]([]int{0, 1, 2, 3}, 2)
// 2

nth := lo.Nth[int]([]int{0, 1, 2, 3}, -2)
// 2
```

### Ternary

A 1 line if/else statement.

```go
result := lo.Ternary[string](true, "a", "b")
// "a"

result := lo.Ternary[string](false, "a", "b")
// "b"
```

### If / ElseIf / Else

```go
result := lo.If[int](true, 1).
    ElseIf(false, 2).
    Else(3)
// 1

result := lo.If[int](false, 1).
    ElseIf(true, 2).
    Else(3)
// 2

result := lo.If[int](false, 1).
    ElseIf(false, 2).
    Else(3)
// 3
```

### Switch / Case / Default

```go
result := lo.Switch[int, string](1).
    Case(1, "1").
    Case(2, "2").
    Default("3")
// "1"

result := lo.Switch[int, string](2).
    Case(1, "1").
    Case(2, "2").
    Default("3")
// "2"

result := lo.Switch[int, string](42).
    Case(1, "1").
    Case(2, "2").
    Default("3")
// "3"
```

Using callbacks:

```go
result := lo.Switch[int, string](1).
    CaseF(1, func () int { return "1" }).
    CaseF(2, func () int { return "2" }).
    DefaultF(func () int { return "3" })
// "1"
```

### ToPtr

Returns a pointer copy of value.

```go
ptr := lo.ToPtr[string]("hello world")
// *string{"hello world"}
```

### ToSlicePtr

Returns a slice of pointer copy of value.

```go
ptr := lo.ToSlicePtr[string]([]string{"hello", "world"})
// []*string{"hello", "world"}
```

## Performance

// TODO

## Contributing

* Ping me on twitter [@samuelberthe](https://twitter.com/samuelberthe) (DMs, mentions, whatever :))
* Fork the [project](https://github.com/samber/lo)
* Fix [open issues](https://github.com/samber/lo/issues) or request new features

Don't hesitate ;)

### Install go 1.18

```bash
make go1.18beta1
```

If your OS currently not default to Go 1.18, replace `BIN=go` by `BIN=go1.18beta1` in the Makefile.

### With Docker

```bash
docker-compose run --rm dev
```

### Without Docker

```bash
# Install some dev dependencies
make tools

# Run tests
make test
# or
make watch-test
```

## Authors

* Samuel Berthe
