package lomap

import "github.com/samber/lo/lotup"

func Keys[Map ~map[K]V, K comparable, V any](m Map) (keys []K) {
	keys = make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func Values[Map ~map[K]V, K comparable, V any](m Map) (values []V) {
	values = make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func Pairs[Map ~map[K]V, K comparable, V any](m Map) (pairs []lotup.Tuple2[K, V]) {
	pairs = make([]lotup.Tuple2[K, V], 0, len(m))
	for k, v := range m {
		pairs = append(pairs, lotup.Of2(k, v))
	}

	return pairs
}

func Unzip[Map ~map[K]V, K comparable, V any](m Map) (keys []K, values []V) {
	keys = make([]K, 0, len(m))
	values = make([]V, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}

	return keys, values
}
