package mutmap

import (
	"github.com/samber/lo/lomap"
	"maps"
)

func MapValues[Map ~map[K]V, K comparable, V any](m Map, fmap func(V) V) {
	for k, v := range m {
		m[k] = fmap(v)
	}
}

func MapKeys[Map ~map[K]V, K comparable, V any](m Map, fmap func(K) V) {
	for k := range m {
		m[k] = fmap(k)
	}
}

func MapPairs[Map ~map[K]V, K comparable, V any](m Map, fmap func(K, V) V) {
	for k, v := range m {
		m[k] = fmap(k, v)
	}
}

func FullMap[Map ~map[K]V, K comparable, V any](m Map, fmap func(K, V) (K, V)) {
	res := lomap.FullMap(m, fmap)

	Clear(m)
	maps.Copy(m, res)
}

func FilterMapValues[Map ~map[K]V, K comparable, V any](m Map, fmap func(V) (V, bool)) {
	for k, v := range m {
		if nv, ok := fmap(v); ok {
			m[k] = nv
		} else {
			delete(m, k)
		}
	}
}

func FilterMapKeys[Map ~map[K]V, K comparable, V any](m Map, fmap func(K) (V, bool)) {
	for k := range m {
		if nv, ok := fmap(k); ok {
			m[k] = nv
		} else {
			delete(m, k)
		}
	}
}

func FilterMapPairs[Map ~map[K]V, K comparable, V any](m Map, fmap func(K, V) (V, bool)) {
	for k, v := range m {
		if nv, ok := fmap(k, v); ok {
			m[k] = nv
		} else {
			delete(m, k)
		}
	}
}

func FilterFullMap[Map ~map[K]V, K comparable, V any](m Map, fmap func(K, V) (K, V, bool)) {
	res := lomap.FilterFullMap(m, fmap)

	Clear(m)
	maps.Copy(m, res)
}
