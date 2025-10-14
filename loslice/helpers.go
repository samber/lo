package loslice

func allocateCapacity[Slice ~[]T, T any](mode AllocateMode, size int, precount func() int) (result Slice) {
	switch mode {
	default:
		return nil // do nothing
	case AllocateZero:
		return make(Slice, 0)
	case AllocateAll:
		return make(Slice, 0, size)
	case AllocatePrecount:
		return make(Slice, 0, precount())
	}
}
