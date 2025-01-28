package lo

import "reflect"

// IsNil checks if a value is nil or if it's a reference type with a nil underlying value.
func IsNil(x any) bool {
	defer func() { recover() }() // nolint:errcheck
	return x == nil || reflect.ValueOf(x).IsNil()
}

// IsNotNil checks if a value is not nil or if it's not a reference type with a nil underlying value.
func IsNotNil(x any) bool {
	return !IsNil(x)
}

// ToPtr returns a pointer copy of value.
func ToPtr[T any](x T) *T {
	return &x
}

// Nil returns a nil pointer of type.
func Nil[T any]() *T {
	return nil
}

// EmptyableToPtr returns a pointer copy of value if it's nonzero.
// Otherwise, returns nil pointer.
func EmptyableToPtr[T any](x T) *T {
	// 🤮
	isZero := reflect.ValueOf(&x).Elem().IsZero()
	if isZero {
		return nil
	}

	return &x
}

// FromPtr returns the pointer value or empty.
func FromPtr[T any](x *T) T {
	if x == nil {
		return Empty[T]()
	}

	return *x
}

// FromPtrOr returns the pointer value or the fallback value.
func FromPtrOr[T any](x *T, fallback T) T {
	if x == nil {
		return fallback
	}

	return *x
}

// ToSlicePtr returns a slice of pointer copy of value.
func ToSlicePtr[T any](collection []T) []*T {
	result := make([]*T, len(collection))

	for i := range collection {
		result[i] = &collection[i]
	}
	return result
}

// FromSlicePtr returns a slice with the pointer values.
// Returns a zero value in case of a nil pointer element.
func FromSlicePtr[T any](collection []*T) []T {
	return Map(collection, func(x *T, _ int) T {
		if x == nil {
			return Empty[T]()
		}
		return *x
	})
}

// FromSlicePtrOr returns a slice with the pointer values or the fallback value.
// Play: https://go.dev/play/p/lbunFvzlUDX
func FromSlicePtrOr[T any](collection []*T, fallback T) []T {
	return Map(collection, func(x *T, _ int) T {
		if x == nil {
			return fallback
		}
		return *x
	})
}

// ToAnySlice returns a slice with all elements mapped to `any` type
func ToAnySlice[T any](collection []T) []any {
	result := make([]any, len(collection))
	for i := range collection {
		result[i] = collection[i]
	}
	return result
}

// FromAnySlice returns an `any` slice with all elements mapped to a type.
// Returns false in case of type conversion failure.
func FromAnySlice[T any](in []any) (out []T, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			out = []T{}
			ok = false
		}
	}()

	result := make([]T, len(in))
	for i := range in {
		result[i] = in[i].(T)
	}
	return result, true
}

// Empty returns the zero value (https://go.dev/ref/spec#The_zero_value).
func Empty[T any]() T {
	var zero T
	return zero
}

// IsEmpty returns true if argument is a zero value.
func IsEmpty[T comparable](v T) bool {
	var zero T
	return zero == v
}

// IsNotEmpty returns true if argument is not a zero value.
func IsNotEmpty[T comparable](v T) bool {
	var zero T
	return zero != v
}

// Coalesce returns the first non-empty arguments. Arguments must be comparable.
func Coalesce[T comparable](values ...T) (result T, ok bool) {
	for i := range values {
		if values[i] != result {
			result = values[i]
			ok = true
			return
		}
	}

	return
}

// CoalesceOrEmpty returns the first non-empty arguments. Arguments must be comparable.
func CoalesceOrEmpty[T comparable](v ...T) T {
	result, _ := Coalesce(v...)
	return result
}

// CoalesceSlice returns the first non-zero slice.
func CoalesceSlice[T any](v ...[]T) ([]T, bool) {
	for i := range v {
		if v[i] != nil && len(v[i]) > 0 {
			return v[i], true
		}
	}
	return []T{}, false
}

// CoalesceSliceOrEmpty returns the first non-zero slice.
func CoalesceSliceOrEmpty[T any](v ...[]T) []T {
	for i := range v {
		if v[i] != nil && len(v[i]) > 0 {
			return v[i]
		}
	}
	return []T{}
}

// CoalesceMap returns the first non-zero map.
func CoalesceMap[K comparable, V any](v ...map[K]V) (map[K]V, bool) {
	for i := range v {
		if v[i] != nil && len(v[i]) > 0 {
			return v[i], true
		}
	}
	return map[K]V{}, false
}

// CoalesceMapOrEmpty returns the first non-zero map.
func CoalesceMapOrEmpty[K comparable, V any](v ...map[K]V) map[K]V {
	for i := range v {
		if v[i] != nil && len(v[i]) > 0 {
			return v[i]
		}
	}
	return map[K]V{}
}
