package lo

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Unwind is a function that orders a collection based on predefined order and returns the sorted collection plus the previous order of the items, making it possible to reverse the process
func Unwind[T constraints.Integer, R any](order []T, collection []R) (sortedCollection []R, stochasticTenet []T) {
	if len(collection) <= 0 {
		sortedCollection = make([]R, 0)
		stochasticTenet = make([]T, 0)
		return
	}

	intermediary := []struct {
		x     T
		y     T
		value R
	}{}

	for i, v := range collection {
		intermediary = append(intermediary, struct {
			x     T
			y     T
			value R
		}{x: order[i], y: T(i), value: v})
	}

	sort.Slice(intermediary, func(i, j int) bool {
		return intermediary[i].x < intermediary[j].x
	})

	for _, v := range intermediary {
		sortedCollection = append(sortedCollection, v.value)
		stochasticTenet = append(stochasticTenet, v.y)
	}

	return
}
