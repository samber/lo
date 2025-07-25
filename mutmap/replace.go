package mutmap

import (
	"github.com/samber/lo/lomap"
	"maps"
)

func ReplaceVal[Map ~map[K]V, K, V comparable](m Map, oldVal, newVal V) {
	for k, v := range m {
		if v == oldVal {
			m[k] = newVal
		}
	}
}

func Replace[Map ~map[K]V, K comparable, V any](m Map, pred func(V) bool, fmap func(V) V) {
	for k, v := range m {
		if pred(v) {
			m[k] = fmap(v)
		}
	}
}

func KReplace[Map ~map[K]V, K comparable, V any](m Map, pred func(K, V) bool, fmap func(K, V) V) {
	for k, v := range m {
		if pred(k, v) {
			m[k] = fmap(k, v)
		}
	}
}

func FullReplace[Map ~map[K]V, K comparable, V any](m Map, pred func(K, V) bool, fmap func(K, V) (K, V)) {
	res := lomap.FullMap(m, func(k K, v V) (K, V) {
		if pred(k, v) {
			return fmap(k, v)
		}

		return k, v
	})

	Clear(m)
	maps.Copy(m, res)
}
