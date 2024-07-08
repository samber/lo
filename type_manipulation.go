package lo

import (
	"unsafe"
)

// IsNil checks if a value is nil or if it's a reference type with a nil underlying value.
func IsNil(x any) bool {
	defer func() { recover() }() // nolint:errcheck
	return x == nil || reflect.ValueOf(x).IsNil()
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
	// You cannot directly use `==` to determine whether it is a zero value,
	// because T is not constrained to be `comparable`
	//   var zero T
	//   if x == zero { ... } // error: T is not comparable

	// If x is a zero value, all of its memory is 0.
	// Leveraging this property, you can convert x into []uint8 to implement a zero-value check.
	size := unsafe.Sizeof(x)
	bytes := unsafe.Slice((*uint8)(unsafe.Pointer(&x)), size)
	for _, b := range bytes {
		if b != 0 {
			return &x
		}
	}
	return nil
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

// Empty returns an empty value.
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
