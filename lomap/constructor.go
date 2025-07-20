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
