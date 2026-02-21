package lomap

func NewGetter[Map ~map[K]V, K comparable, V any](m Map) func(K) (V, bool) {
	return func(k K) (v V, ok bool) {
		v, ok = m[k]
		return
	}
}

func NewGetterWithDefault[Map ~map[K]V, K comparable, V any](m Map, defaultValue V) func(K) V {
	return func(k K) V {
		if v, ok := m[k]; ok {
			return v
		}

		return defaultValue
	}
}

func NewGetterOrDefault[Map ~map[K]V, K comparable, V any](m Map) func(K, V) V {
	return func(k K, defaultValue V) V {
		if v, ok := m[k]; ok {
			return v
		}

		return defaultValue
	}
}

func NewGetterWithFallback[Map ~map[K]V, K comparable, V any](m Map, fallback func(K) V) func(K) V {
	return func(k K) V {
		if v, ok := m[k]; ok {
			return v
		}

		return fallback(k)
	}
}

func NewGetterOrFallback[Map ~map[K]V, K comparable, V any](m Map) func(K, func(K) V) V {
	return func(k K, fallback func(K) V) V {
		if v, ok := m[k]; ok {
			return v
		}

		return fallback(k)
	}
}
