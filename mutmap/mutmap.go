package mutmap

func ForceNil[Map ~map[K]V, K comparable, V any](m Map) Map {
	if len(m) == 0 {
		return nil
	}

	return m
}

func ForceEmpty[Map ~map[K]V, K comparable, V any](m Map) Map {
	if m == nil {
		return make(Map)
	}

	return m
}

func ToPointers[Map ~map[K]V, K comparable, V any](m Map) (res map[K]*V) {
	pointers := make([]V, len(m))
	res = make(map[K]*V, len(m))

	var k K
	i := 0

	for k, pointers[i] = range m {
		res[k] = &pointers[i]
		i++
	}

	return res
}

func Clear[Map ~map[K]V, K comparable, V any](m Map) {
	for k := range m {
		delete(m, k)
	}
}
