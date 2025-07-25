package mutmap

func Safe[Map ~map[K]V, K comparable, V any](m Map) Map {
	if m == nil {
		return make(Map)
	}

	return m
}
