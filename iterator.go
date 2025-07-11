//go:build goexperiment.rangefunc

package lo

// Iterator TODO
type Iterator[V any] func(func(int, V) bool)

// ToIterator TODO
func ToIterator[V any](collection ...V) Iterator[V] {
	return func(yield func(int, V) bool) {
		for i, item := range collection {
			yield(i, item)
		}
	}
}

// Len TODO
func (iter Iterator[V]) Len() int {
	var n int
	for _, _ = range iter {
		n++
	}
	return n
}

// Slice TODO
func (iter Iterator[V]) Slice() []V {
	var result []V
	for _, item := range iter {
		result = append(result, item)
	}
	return result
}

// FilterI returns an iterator of all elements that match the specified predicate callback.
// Play: https://go.dev/play/p/TODO
func FilterI[V any](iter Iterator[V], predicate func(item V, index int) bool) Iterator[V] {
	return func(yield func(int, V) bool) {
		for i, item := range iter {
			if predicate(item, i) {
				if !yield(i, item) {
					break
				}
			}
		}
	}
}

// MapI transforms an iterator into an iterator of another type.
// Play: https://go.dev/play/p/TODO
func MapI[T any, R any](iter Iterator[T], iteratee func(item T, index int) R) Iterator[R] {
	return func(yield func(int, R) bool) {
		for i, item := range iter {
			if !yield(i, iteratee(item, i)) {
				break
			}
		}
	}
}

// FilterMapI returns an iterator obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
//
// Play: https://go.dev/play/p/TODO
func FilterMapI[T any, R any](iter Iterator[T], callback func(item T, index int) (R, bool)) Iterator[R] {
	return func(yield func(int, R) bool) {
		for i, item := range iter {
			if r, ok := callback(item, i); ok {
				if !yield(i, r) {
					break
				}
			}
		}
	}
}
