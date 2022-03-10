package lo

import "golang.org/x/exp/constraints"

// Range creates an array of numbers (positive and/or negative) with given length.
func Range(elementNum int) []int {
	if elementNum >= 0 {
		result := make([]int, elementNum)
		for i := 0; i < elementNum; i++ {
			result[i] = i
		}
		return result
	}
	result := make([]int, -elementNum)
	for i := 0; i < -elementNum; i++ {
		result[i] = -i
	}
	return result
}

// RangeFrom creates an array of numbers from start with specified length.
func RangeFrom[T constraints.Integer | constraints.Float](start T, elementNum int) []T {
	if elementNum >= 0 {
		result := make([]T, elementNum)
		for i := 0; i < elementNum; i++ {
			result[i] = T(i) + start
		}
		return result
	}
	result := make([]T, -elementNum)
	for i := 0; i < -elementNum; i++ {
		result[i] = start - T(i)
	}
	return result
}

// RangeWithSteps creates an array of numbers (positive and/or negative) progressing from start up to, but not including end.
// step set to zero will return empty array.
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
