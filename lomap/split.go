package lomap

func SplitKeys[K comparable, V any, Map ~map[K]V](m Map, pred func(K) bool) (taken, rest Map) {
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
