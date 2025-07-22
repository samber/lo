package lomap

import "github.com/samber/lo/lotup"

func FromKeysAndValues[K comparable, V any](keys []K, values []V) (m map[K]V) {
	size := min(len(keys), len(values))

	m = make(map[K]V, size)
	for i := range size {
		m[keys[i]] = values[i]
	}

	return m
}

func FromPairs[K comparable, V any](pairs []lotup.Tuple2[K, V]) (m map[K]V) {
	m = make(map[K]V, len(pairs))
	for _, p := range pairs {
		m[p.V1] = p.V2
	}

	return m
}

func FromKeysAndDefault[K comparable, V any](keys []K, value V) (m map[K]V) {
	m = make(map[K]V, len(keys))
	for _, k := range keys {
		m[k] = value
	}

	return m
}

func MapKeys[K comparable, V any](keys []K, fmap func(K) V) (m map[K]V) {
	m = make(map[K]V, len(keys))
	for _, k := range keys {
		m[k] = fmap(k)
	}

	return m
}

func IMapKeys[K comparable, V any](keys []K, imap func(int, K) V) (m map[K]V) {
	m = make(map[K]V, len(keys))
	for i, k := range keys {
		m[k] = imap(i, k)
	}

	return m
}

func MapValues[K comparable, V any](values []V, fmap func(V) K) (m map[K]V) {
	m = make(map[K]V, len(values))
	for _, v := range values {
		m[fmap(v)] = v
	}

	return m
}

func IMapValues[K comparable, V any](values []V, imap func(int, V) K) (m map[K]V) {
	m = make(map[K]V, len(values))
	for i, v := range values {
		m[imap(i, v)] = v
	}

	return m
}

func Fill[K comparable, V any](size int, get func(int) (K, V)) (m map[K]V) {
	if size < 0 {
		return nil
	}

	m = make(map[K]V, size)
	for i := 0; i < size; i++ {
		k, v := get(i)
		m[k] = v
	}

	return m
}
