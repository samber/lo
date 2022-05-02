package lo

const nilPointerPanic = "runtime error: invalid memory address or nil pointer dereference"

// ToPtr returns a pointer copy of value.
func ToPtr[T any](x T) *T {
	return &x
}

// ToSlicePtr returns a slice of pointer copy of value.
func ToSlicePtr[T any](collection []T) []*T {
	return Map(collection, func(x T, _ int) *T {
		return &x
	})
}

// Empty returns an empty value.
func Empty[T any]() T {
	var t T
	return t
}

// Coalesce returns the first non-empty arguments. Arguments must be comparable.
func Coalesce[T comparable](v ...T) (result T, ok bool) {
	for _, e := range v {
		if e != result {
			result = e
			ok = true
			return
		}
	}

	return
}

func Safe[T any](cb func() T) (result T, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok && err.Error() == nilPointerPanic {
				// catch error
			} else {
				panic(r)
			}
		}
	}()

	result = cb()
	ok = true
	return
}
