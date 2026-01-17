package loset

import "iter"

type Set[T comparable] = map[T]struct{}

func FromSlice[Slice ~[]T, T comparable](xs Slice) Set[T] {
	set := make(Set[T], len(xs))

	for _, x := range xs {
		set[x] = struct{}{}
	}

	return set
}

func Collect[T comparable, Seq iter.Seq[T]](xs Seq) Set[T] {
	set := make(Set[T])

	for x := range xs {
		set[x] = struct{}{}
	}

	return set
}

func ToSlice[S ~Set[T], T comparable](set S) []T {
	slice := make([]T, 0, len(set))

	for x := range set {
		slice = append(slice, x)
	}

	return slice
}
