package lomap

func SplitKeys[Map ~map[K]V, K comparable, V any](m Map, pred func(K) bool) (taken, rest Map) {
	taken = make(Map, len(m))
	rest = make(Map, len(m))

	for k, v := range m {
		if pred(k) {
			taken[k] = v
		} else {
			rest[k] = v
		}
	}

	return taken, rest
}

func SplitValues[Map ~map[K]V, K comparable, V any](m Map, pred func(V) bool) (taken, rest Map) {
	taken = make(Map, len(m))
	rest = make(Map, len(m))

	for k, v := range m {
		if pred(v) {
			taken[k] = v
		} else {
			rest[k] = v
		}
	}

	return taken, rest
}

func SplitPairs[Map ~map[K]V, K comparable, V any](m Map, pred func(K, V) bool) (taken, rest Map) {
	taken = make(Map, len(m))
	rest = make(Map, len(m))

	for k, v := range m {
		if pred(k, v) {
			taken[k] = v
		} else {
			rest[k] = v
		}
	}

	return taken, rest
}
