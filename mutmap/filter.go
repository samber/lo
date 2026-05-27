package mutmap

// In-place reject: removes all entries with value == val.
func RejectVal[Map ~map[K]V, K, V comparable](m Map, val V) {
	for k, v := range m {
		if v == val {
			delete(m, k)
		}
	}
}

// In-place filter: keeps only keys for which pred returns true.
func FilterKeys[Map ~map[K]V, K comparable, V any](m Map, pred func(K) bool) {
	for k := range m {
		if !pred(k) {
			delete(m, k)
		}
	}
}

// In-place filter: keeps only values for which pred returns true.
func FilterValues[Map ~map[K]V, K comparable, V any](m Map, pred func(V) bool) {
	for k, v := range m {
		if !pred(v) {
			delete(m, k)
		}
	}
}

// In-place filter: keeps only pairs for which pred returns true.
func FilterPairs[Map ~map[K]V, K comparable, V any](m Map, pred func(K, V) bool) {
	for k, v := range m {
		if !pred(k, v) {
			delete(m, k)
		}
	}
}
