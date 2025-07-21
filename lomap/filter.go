package lomap

func FilterKeys[K comparable, V any, Map ~map[K]V](m Map, pred func(K) bool) (filtered Map) {
	filtered = make(Map, len(m))
	for k, v := range m {
		if pred(k) {
			filtered[k] = v
		}
	}

	return filtered
}

func FilterValues[K comparable, V any, Map ~map[K]V](m Map, pred func(V) bool) (filtered Map) {
	filtered = make(Map, len(m))
	for k, v := range m {
		if pred(v) {
			filtered[k] = v
		}
	}

	return filtered
}

func FilterPairs[K comparable, V any, Map ~map[K]V](m Map, pred func(K, V) bool) (filtered Map) {
	filtered = make(Map, len(m))
	for k, v := range m {
		if pred(k, v) {
			filtered[k] = v
		}
	}

	return filtered
}
