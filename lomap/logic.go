package lomap

import "github.com/samber/lo/loslice"

func ContainsVal[Map ~map[K]V, K, V comparable](m Map, v V) bool {
	for _, val := range m {
		if val == v {
			return true
		}
	}

	return false
}

func WithoutVal[Map ~map[K]V, K, V comparable](m Map, v V) bool {
	return !ContainsVal(m, v)
}

func ContainsKey[Map ~map[K]V, K comparable, V any](m Map, pred func(K) bool) bool {
	for k := range m {
		if pred(k) {
			return true
		}
	}

	return false
}

func WithoutKeyPred[Map ~map[K]V, K comparable, V any](m Map, pred func(K) bool) bool {
	return !ContainsKey(m, pred)
}

func Contains[Map ~map[K]V, K comparable, V any](m Map, pred func(V) bool) bool {
	for _, v := range m {
		if pred(v) {
			return true
		}
	}

	return false
}

func Without[Map ~map[K]V, K comparable, V any](m Map, pred func(V) bool) bool {
	return !Contains(m, pred)
}

func KContains[Map ~map[K]V, K comparable, V any](m Map, pred func(K, V) bool) bool {
	for k, v := range m {
		if pred(k, v) {
			return true
		}
	}

	return false
}

func KWithout[Map ~map[K]V, K comparable, V any](m Map, pred func(K, V) bool) bool {
	return !KContains(m, pred)
}

func EveryVal[Map ~map[K]V, K, V comparable](m Map, v V) bool {
	for _, val := range m {
		if val != v {
			return false
		}
	}

	return true
}

func EveryKey[Map ~map[K]V, K comparable, V any](m Map, pred func(K) bool) bool {
	for k := range m {
		if !pred(k) {
			return false
		}
	}

	return true
}

func Every[Map ~map[K]V, K comparable, V any](m Map, pred func(V) bool) bool {
	for _, v := range m {
		if !pred(v) {
			return false
		}
	}

	return true
}

func KEvery[Map ~map[K]V, K comparable, V any](m Map, pred func(K, V) bool) bool {
	for k, v := range m {
		if !pred(k, v) {
			return false
		}
	}

	return true
}

func Disjoint[Map ~map[K]V, K comparable, V any](a, b Map) bool {
	if len(b) < len(a) {
		a, b = b, a // Ensure a is the smaller map
	}

	for k := range a {
		if _, exists := b[k]; exists {
			return false
		}
	}

	return true
}

func GroupDisjoint[Map ~map[K]V, K comparable, V any](ms ...Map) bool {
	if len(ms) == 0 || loslice.Every(ms, IsEmpty) {
		return true
	}

	seen := make(map[K]struct{}, loslice.MapSum(ms, Len))

	for _, m := range ms {
		for k := range m {
			if _, exists := seen[k]; exists {
				return false // Found a duplicate key
			}
			seen[k] = struct{}{} // Mark the key as seen
		}
	}

	return true
}
