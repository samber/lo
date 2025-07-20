package lomap

func Enumerate[K comparable, V any, Map ~map[K]V](m Map) (indices map[K]int) {
	indices = make(map[K]int, len(m))
	i := 0
	for k := range m {
		indices[k] = i
		i++
	}

	return indices
}

func Swap[K, V comparable, Map ~map[K]V](m Map) (swapped map[V]K) {
	swapped = make(map[V]K, len(m))
	for k, v := range m {
		swapped[v] = k
	}

	return swapped
}
