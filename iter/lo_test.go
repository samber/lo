package iter

import (
	"iter"
	"slices"
)

func values[T any](v ...T) iter.Seq[T] { return slices.Values(v) }

type foo struct {
	bar string
}

func (f foo) Clone() foo {
	return foo{f.bar}
}
