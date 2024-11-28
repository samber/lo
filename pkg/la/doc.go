// Package la represents a collection of functions that can be evaluated lazily
// by the golang's iterator feature.
//
// They targeted to be a mirror for main functions provided by the
// [lo](https://github.com/samber/lo) package but for iterators.
//
// There are few functions useful if you use iterator as a slice (like backward
// iteration helper) because if you do that – prefer to convert them to
// slices/maps and use the great [lo](https://github.com/samber/lo) package. As
// you can see in the provided benchmarks, if you have only a few layers of Map
// transformers, they will be not too much slower than when you use an iterator,
// but they're easier to use.
package la
