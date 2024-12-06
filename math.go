package lo

import (
	"fmt"
	"github.com/samber/lo/internal/constraints"
	"math"
	"strconv"
)

// Range creates an array of numbers (positive and/or negative) with given length.
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

// RangeFrom creates an array of numbers from start with specified length.
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

// RangeWithSteps creates an array of numbers (positive and/or negative) progressing from start up to, but not including end.
// step set to zero will return empty array.
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
func Clamp[T constraints.Ordered](value T, min T, max T) T {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

// Sum sums the values in a collection. If collection is empty 0 is returned.
// Play: https://go.dev/play/p/upfeJVqs4Bt
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T {
	var sum T = 0
	for i := range collection {
		sum += collection[i]
	}
	return sum
}

// SumBy summarizes the values in a collection using the given return value from the iteration function. If collection is empty 0 is returned.
// Play: https://go.dev/play/p/Dz_a_7jN_ca
func SumBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection []T, iteratee func(item T) R) R {
	var sum R = 0
	for i := range collection {
		sum = sum + iteratee(collection[i])
	}
	return sum
}

// Mean calculates the mean of a collection of numbers.
func Mean[T constraints.Float | constraints.Integer](collection []T) T {
	var length = T(len(collection))
	if length == 0 {
		return 0
	}
	var sum = Sum(collection)
	return sum / length
}

// MeanBy calculates the mean of a collection of numbers using the given return value from the iteration function.
func MeanBy[T any, R constraints.Float | constraints.Integer](collection []T, iteratee func(item T) R) R {
	var length = R(len(collection))
	if length == 0 {
		return 0
	}
	var sum = SumBy(collection, iteratee)
	return sum / length
}

// Round returns the float32/float64 of rounding half away from the specified precision
func Round[T float64 | float32](f T, n ...int) T {
	var nn = 3
	if len(n) > 0 {
		if n[0] >= 0 && n[0] <= 15 {
			nn = n[0]
		}
	}
	r, _ := strconv.ParseFloat(fmt.Sprintf("%.*f", nn, f), 64)
	return T(r)
}

// Truncate returns the float32/float64 of the specified precision
func Truncate[T float64 | float32](f T, n ...int) T {
	var nn = 3
	if len(n) > 0 {
		nn = n[0]
		if n[0] >= 0 && n[0] <= 15 {
			nn = n[0]
		}
	}
	pow10N := math.Pow10(nn)
	integer, fractional := math.Modf(float64(f))
	r, _ := strconv.ParseFloat(fmt.Sprintf("%.*f", nn, integer+math.Trunc(fractional*pow10N)/pow10N), 64)
	return T(r)
}
