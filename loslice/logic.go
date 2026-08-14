package loslice

import (
	"github.com/samber/lo/lotup"
	"slices"
)

// ContainsVal checks if there is at least one element in the slice that equals the given value.
// Supports early exit, so it will stop checking as soon as it finds a match.
func ContainsVal[Slice ~[]T, T comparable](xs Slice, val T) bool {
	return slices.Contains(xs, val)
}

func WithoutVal[Slice ~[]T, T comparable](xs Slice, val T) bool {
	return !ContainsVal(xs, val)
}

// Contains checks if there is at least one element in the slice that satisfies the predicate.
func Contains[Slice ~[]T, T any](xs Slice, pred func(item T) bool) bool {
	return lotup.Second(Index(xs, pred))
}

func Without[Slice ~[]T, T any](xs Slice, pred func(item T) bool) bool {
	return !Contains(xs, pred)
}

// IContains checks if there is at least one element in the slice that satisfies the indexed predicate.
func IContains[Slice ~[]T, T any](xs Slice, ipred func(int, T) bool) bool {
	return lotup.Second(IIndex(xs, ipred))
}

func IWithout[Slice ~[]T, T any](xs Slice, ipred func(int, T) bool) bool {
	return !IContains(xs, ipred)
}

func EveryVal[Slice ~[]T, T comparable](xs Slice, val T) bool {
	for _, x := range xs {
		if x != val {
			return false
		}
	}

	return true
}

func Every[Slice ~[]T, T any](xs Slice, pred func(item T) bool) bool {
	for _, x := range xs {
		if !pred(x) {
			return false
		}
	}

	return true
}

func IEvery[Slice ~[]T, T any](xs Slice, ipred func(int, T) bool) bool {
	for i, x := range xs {
		if !ipred(i, x) {
			return false
		}
	}

	return true
}
