package loany

func Map[Map ~map[K]V, K comparable, V any](m Map) map[K]any {
	result := make(map[K]any, len(m))
	for k, v := range m {
		result[k] = v
	}

	return result
}

// TypedMap converts a map of any type to a map of a specific type V.
// Values that cannot be converted to V will be omitted.
func TypedMap[Map ~map[K]any, V any, K comparable](m Map) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		if val, ok := v.(V); ok {
			result[k] = val
		}
	}

	return result
}
