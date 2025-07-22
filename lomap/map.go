package lomap

func Enumerate[Map ~map[K]V, K comparable, V any](m Map) (indices map[K]int) {
	indices = make(map[K]int, len(m))
	i := 0
	for k := range m {
		indices[k] = i
		i++
	}

	return indices
}

func Swap[Map ~map[K]V, K, V comparable](m Map) (swapped map[V]K) {
	swapped = make(map[V]K, len(m))
	for k, v := range m {
		swapped[v] = k
	}

	return swapped
}

func Any[Map ~map[K]V, K comparable, V any](m Map) (k K, v V) {
	for k, v = range m {
		return k, v
	}

	return
}
