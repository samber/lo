package lomap

func CountKeys[K comparable, V any, Map ~map[K]V](m Map, pred func(K) bool) (count int) {
	for k := range m {
		if pred(k) {
			count++
		}
	}

	return count
}

func CountValues[K comparable, V any, Map ~map[K]V](m Map, pred func(V) bool) (count int) {
	for _, v := range m {
		if pred(v) {
			count++
		}
	}

	return count
}

func CountPairs[K comparable, V any, Map ~map[K]V](m Map, pred func(K, V) bool) (count int) {
	for k, v := range m {
		if pred(k, v) {
			count++
		}
	}

	return count
}
