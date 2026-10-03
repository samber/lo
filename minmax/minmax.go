package minmax

import "golang.org/x/exp/constraints"

// Min returns the minimum value of items.
func Min[T constraints.Ordered](items ...T) T {
	var min T

	if len(items) == 0 {
		return min
	}

	min = items[0]

	for i := 1; i < len(items); i++ {
		if items[i] < min {
			min = items[i]
		}
	}

	return min
}

// Max returns the maximum value of items.
func Max[T constraints.Ordered](items ...T) T {
	var max T

	if len(items) == 0 {
		return max
	}

	max = items[0]

	for i := 1; i < len(items); i++ {

		if items[i] > max {
			max = items[i]
		}
	}

	return max
}
