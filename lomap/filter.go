package lomap

func FilterKeys[Map ~map[K]V, K comparable, V any](m Map, pred func(K) bool) (filtered Map) {
	filtered = make(Map, len(m))
	for k, v := range m {
		if pred(k) {
			filtered[k] = v
		}
	}

	return filtered
}

func FilterValues[Map ~map[K]V, K comparable, V any](m Map, pred func(V) bool) (filtered Map) {
	filtered = make(Map, len(m))
	for k, v := range m {
		if pred(v) {
			filtered[k] = v
		}
	}

	return filtered
}

func FilterPairs[Map ~map[K]V, K comparable, V any](m Map, pred func(K, V) bool) (filtered Map) {
	filtered = make(Map, len(m))
	for k, v := range m {
		if pred(k, v) {
			filtered[k] = v
		}
	}

	return filtered
}
