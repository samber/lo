package lomap

import (
	"github.com/samber/lo/loslice"
	"maps"
	"slices"
)

func IsNil[Map ~map[K]V, K comparable, V any](m Map) bool {
	return m == nil
}

func IsEmpty[Map ~map[K]V, K comparable, V any](m Map) bool {
	return len(m) == 0
}

func Len[Map ~map[K]V, K comparable, V any](m Map) int {
	return len(m)
}

func MapValues[Map ~map[K]V, K comparable, V, R any](m Map, fmap func(V) R) map[K]R {
	if m == nil {
		return nil
	}

	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = fmap(v)
	}

	return result
}

func MapKeys[Map ~map[K]V, K comparable, V, R any](m Map, fmap func(K) R) map[K]R {
	if m == nil {
		return nil
	}

	result := make(map[K]R, len(m))
	for k := range m {
		result[k] = fmap(k)
	}

	return result
}

func MapPairs[Map ~map[K]V, K comparable, V, R any](m Map, fmap func(K, V) R) map[K]R {
	if m == nil {
		return nil
	}

	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = fmap(k, v)
	}

	return result
}

func FullMap[Map ~map[K]V, K, Q comparable, V, R any](m Map, fmap func(K, V) (Q, R)) map[Q]R {
	if m == nil {
		return nil
	}

	result := make(map[Q]R, len(m))
	for k, v := range m {
		q, r := fmap(k, v)
		result[q] = r
	}

	return result
}

func FilterMapValues[Map ~map[K]V, K comparable, V, R any](m Map, fmap func(V) (R, bool)) map[K]R {
	if m == nil {
		return nil
	}

	result := make(map[K]R, len(m))
	for k, v := range m {
		if r, ok := fmap(v); ok {
			result[k] = r
		}
	}

	return result
}

func FilterMapKeys[Map ~map[K]V, K comparable, V, R any](m Map, fmap func(K) (R, bool)) map[K]R {
	if m == nil {
		return nil
	}

	result := make(map[K]R, len(m))
	for k := range m {
		if r, ok := fmap(k); ok {
			result[k] = r
		}
	}

	return result
}

func FilterMapPairs[Map ~map[K]V, K comparable, V, R any](m Map, fmap func(K, V) (R, bool)) map[K]R {
	if m == nil {
		return nil
	}

	result := make(map[K]R, len(m))
	for k, v := range m {
		if r, ok := fmap(k, v); ok {
			result[k] = r
		}
	}

	return result
}

func FilterFullMap[Map ~map[K]V, K, Q comparable, V, R any](m Map, fmap func(K, V) (Q, R, bool)) map[Q]R {
	if m == nil {
		return nil
	}

	result := make(map[Q]R, len(m))
	for k, v := range m {
		q, r, ok := fmap(k, v)
		if ok {
			result[q] = r
		}
	}

	return result
}

func Merge[Map ~map[K]V, K comparable, V any](a, b Map) Map {
	if a == nil && b == nil {
		return nil
	} else if a == nil {
		return maps.Clone(b)
	} else if b == nil {
		return maps.Clone(a)
	}

	result := make(Map, len(a)+len(b))
	maps.Copy(result, a)
	maps.Copy(result, b)

	return result
}

func GroupMerge[Map ~map[K]V, K comparable, V any](ms ...Map) Map {
	if len(ms) == 0 || loslice.Every(ms, IsNil) {
		return nil
	}

	result := make(Map, loslice.MapSum(ms, Len))
	for _, m := range ms {
		maps.Copy(result, m)
	}

	return result
}

func Common[Map ~map[K]V, K comparable, V any](a, b Map) (ia Map, ib Map) {
	if a == nil || b == nil {
		return nil, nil
	}

	swapped := false
	if len(b) < len(a) {
		swapped = true
		a, b = b, a // Ensure a is the smaller map
	}

	ia, ib = make(Map, len(a)), make(Map, len(a))
	for k, va := range a {
		if vb, ok := b[k]; ok {
			ia[k] = va
			ib[k] = vb
		}
	}

	if swapped {
		ia, ib = ib, ia // Swap back if we swapped a and b
	}

	return
}

func GroupCommon[Map ~map[K]V, K comparable, V any](maps ...Map) map[K][]V {
	if len(maps) == 0 || loslice.Every(maps, IsNil) {
		return nil
	}

	// Find the smallest map to reduce iterations
	smallest, size := loslice.Min(maps, Len)

	if size == 0 {
		return make(map[K][]V)
	}

	intersection := make(map[K][]V, size)
	alternatives := make([]V, 0, len(maps))

keysLoop:
	for k := range smallest {
		alternatives = alternatives[:0] // Reset alternatives for each key
		for _, m := range maps {
			if v, ok := m[k]; ok {
				alternatives = append(alternatives, v)
			} else {
				continue keysLoop // If any map does not have the key, skip to the next key
			}
		}

		intersection[k] = slices.Clone(alternatives)
	}

	return intersection
}

func Distinct[Map ~map[K]V, K comparable, V any](a, b Map) Map {
	if a == nil && b == nil {
		return nil
	} else if a == nil {
		return maps.Clone(b)
	} else if b == nil {
		return maps.Clone(a)
	}

	if len(a) < len(b) {
		// swap a and b to ensure b is the smaller map
		a, b = b, a
	}

	result := make(Map, len(a)+len(b))
	maps.Copy(result, a) // clone larger map

	for k := range b { // iterate over smaller map
		if _, exists := result[k]; exists {
			delete(result, k)
		} else {
			result[k] = b[k]
		}
	}

	return result
}

func GroupDistinct[Map ~map[K]V, K comparable, V any](ms ...Map) Map {
	if len(ms) == 0 || loslice.Every(ms, IsNil) {
		return nil
	}

	size := loslice.MapSum(ms, Len)
	result := make(Map, size)
	seen := make(map[K]struct{}, size)
	for _, m := range ms {
		for k, v := range m {
			if _, exists := seen[k]; exists {
				continue // Skip if the key has already been seen
			} else if _, exists := result[k]; exists {
				// If the key exists in the result, delete it from the result
				// and mark it as seen to avoid duplicates
				delete(result, k)
				seen[k] = struct{}{}
			} else {
				result[k] = v
			}
		}
	}

	return result
}

// Unique returns a new map containing keys from `a` that are not present in `b`.
func Unique[Map ~map[K]V, K comparable, V any](a, b Map) Map {
	if a == nil && b == nil {
		return nil
	} else if a == nil {
		return maps.Clone(b)
	} else if b == nil {
		return maps.Clone(a)
	}

	result := make(Map, len(a))
	for k := range a {
		if _, exists := b[k]; !exists {
			result[k] = a[k] // Add only keys from a that are not in b
		}
	}

	return result
}
