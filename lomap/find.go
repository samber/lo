package lomap

func FindVal[Map ~map[K]V, K, V comparable](m Map, v V) (k K, ok bool) {
	for k, val := range m {
		if val == v {
			return k, true
		}
	}

	return
}

func FindKey[Map ~map[K]V, K comparable, V any](m Map, pred func(K) bool) (k K, ok bool) {
	for k := range m {
		if pred(k) {
			return k, true
		}
	}

	return
}

func Find[Map ~map[K]V, K comparable, V any](m Map, pred func(V) bool) (k K, ok bool) {
	for k, v := range m {
		if pred(v) {
			return k, true
		}
	}

	return
}

func KFind[Map ~map[K]V, K comparable, V any](m Map, pred func(K, V) bool) (k K, ok bool) {
	for k, v := range m {
		if pred(k, v) {
			return k, true
		}
	}

	return
}
