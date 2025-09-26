package lo

import (
	"github.com/samber/lo/internal/constraints"
)

// Range creates a slice of numbers (positive and/or negative) with given length.
// Play: https://go.dev/play/p/0r6VimXAi9H
func Range(elementNum int) []int {
	length := If(elementNum < 0, -elementNum).Else(elementNum)
	result := make([]int, length)
	step := If(elementNum < 0, -1).Else(1)
	for i, j := 0, 0; i < length; i, j = i+1, j+step {
		result[i] = j
	}
	return result
}

// RangeFrom creates a slice of numbers from start with specified length.
// Play: https://go.dev/play/p/0r6VimXAi9H
func RangeFrom[T constraints.Integer | constraints.Float](start T, elementNum int) []T {
	length := If(elementNum < 0, -elementNum).Else(elementNum)
	result := make([]T, length)
	step := If(elementNum < 0, -1).Else(1)
	for i, j := 0, start; i < length; i, j = i+1, j+T(step) {
		result[i] = j
	}
	return result
}

// RangeWithSteps creates a slice of numbers (positive and/or negative) progressing from start up to, but not including end.
// step set to zero will return an empty slice.
// Play: https://go.dev/play/p/0r6VimXAi9H
func RangeWithSteps[T constraints.Integer | constraints.Float](start, end, step T) []T {
	result := []T{}
	if start == end || step == 0 {
		return result
	}
	if start < end {
		if step < 0 {
			return result
		}
		for i := start; i < end; i += step {
			result = append(result, i)
		}
		return result
	}
	if step > 0 {
		return result
	}
	for i := start; i > end; i += step {
		result = append(result, i)
	}
	return result
}

// Clamp clamps number within the inclusive lower and upper bounds.
// Play: https://go.dev/play/p/RU4lJNC2hlI
func Clamp[T constraints.Ordered](value, mIn, mAx T) T {
	if value < mIn {
		return mIn
	} else if value > mAx {
		return mAx
	}
	return value
}

// Sum sums the values in a collection. If collection is empty 0 is returned.
// Play: https://go.dev/play/p/upfeJVqs4Bt
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T {
	var sum T
	for i := range collection {
		sum += collection[i]
	}
	return sum
}

// SumBy summarizes the values in a collection using the given return value from the iteration function. If collection is empty 0 is returned.
// Play: https://go.dev/play/p/Dz_a_7jN_ca
func SumBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection []T, iteratee func(item T) R) R {
	var sum R
	for i := range collection {
		sum += iteratee(collection[i])
	}
	return sum
}

// Product gets the product of the values in a collection. If collection is empty 1 is returned.
// Play: https://go.dev/play/p/2_kjM_smtAH
func Product[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T {
	if collection == nil {
		return 1
	}

	if len(collection) == 0 {
		return 1
	}

	var product T = 1
	for i := range collection {
		product *= collection[i]
	}
	return product
}

// ProductBy summarizes the values in a collection using the given return value from the iteration function. If collection is empty 1 is returned.
// Play: https://go.dev/play/p/wadzrWr9Aer
func ProductBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection []T, iteratee func(item T) R) R {
	if collection == nil {
		return 1
	}

	if len(collection) == 0 {
		return 1
	}

	var product R = 1
	for i := range collection {
		product *= iteratee(collection[i])
	}
	return product
}

// Mean calculates the mean of a collection of numbers.
// Play: https://go.dev/play/p/tPURSuteUsP
func Mean[T constraints.Float | constraints.Integer](collection []T) T {
	length := T(len(collection))
	if length == 0 {
		return 0
	}
	sum := Sum(collection)
	return sum / length
}

// MeanBy calculates the mean of a collection of numbers using the given return value from the iteration function.
// Play: https://go.dev/play/p/j7TsVwBOZ7P
func MeanBy[T any, R constraints.Float | constraints.Integer](collection []T, iteratee func(item T) R) R {
	length := R(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumBy(collection, iteratee)
	return sum / length
}

// Mode returns the mode (most frequent value) of a collection.
// If multiple values have the same highest frequency, then multiple values are returned.
// If the collection is empty, then the zero value of T is returned.
func Mode[T constraints.Integer | constraints.Float](collection []T) []T {
	length := T(len(collection))
	if length == 0 {
		return []T{}
	}

	mode := make([]T, 0)
	maxFreq := 0
	frequency := make(map[T]int)

	for _, item := range collection {
		frequency[item]++
		count := frequency[item]

		if count > maxFreq {
			maxFreq = count
			mode = []T{item}
		} else if count == maxFreq {
			mode = append(mode, item)
		}
	}

	return mode
}
