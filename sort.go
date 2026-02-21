package lo

import "sort"

// IndexMap build a map[T]int which the map value was the element index
func IndexMap[T comparable](collection []T) map[T]int {
	m := make(map[T]int, len(collection))
	for i, v := range collection {
		m[v] = i
	}
	return m
}

// SortAs sort collection as index of key which extract by keyFn
func SortAs[T any, K comparable](collection []T, keys []K, keyFn func(T) K) {
	indexesMap := IndexMap(keys)
	sort.Slice(collection, func(i, j int) bool {
		return indexesMap[keyFn(collection[i])] < indexesMap[keyFn(collection[j])]
	})
}
