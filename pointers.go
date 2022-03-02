package lo

// ToPtr returns a pointer copy of value.
func ToPtr[T any](x T) *T {
	return &x
}

// ToPtr returns a slice of pointer copy of value.
func ToSlicePtr[T any](collection []T) []*T {
	return Map(collection, ToPtr[T])
}
