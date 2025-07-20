package lomap

func NewGetter[K comparable, V any, Map ~map[K]V](m Map) func(K) (V, bool) {
	return func(k K) (v V, ok bool) {
		v, ok = m[k]
		return
	}
}

func NewGetterWithDefault[K comparable, V any, Map ~map[K]V](m Map, defaultValue V) func(K) V {
	return func(k K) V {
		if v, ok := m[k]; ok {
			return v
		}

		return defaultValue
	}
}

func NewGetterWithFallback[K comparable, V any, Map ~map[K]V](m Map, fallback func(K) V) func(K) V {
	return func(k K) V {
		if v, ok := m[k]; ok {
			return v
		}

		return fallback(k)
	}
}
