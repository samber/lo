package loslice

func Split[Slice ~[]T, T any](xs Slice, pred func(T) bool) (taken, rest Slice) {
	return splitImpl(xs, pred, nil, nil)
}

func SplitEx[Slice ~[]T, T any](mode AllocateMode, xs Slice, pred func(T) bool) (taken, rest Slice) {
	if xs == nil {
		return nil, nil
	}

	taken = allocateCapacity[Slice](mode, len(xs), func() int { return Count(xs, pred) })
	rest = allocateCapacity[Slice](mode, len(xs), func() int { return len(xs) - len(taken) })

	return splitImpl(xs, pred, taken, rest)
}

func splitImpl[Slice ~[]T, T any](xs Slice, pred func(T) bool, taken, rest Slice) (Slice, Slice) {
	for _, x := range xs {
		if pred(x) {
			taken = append(taken, x)
		} else {
			rest = append(rest, x)
		}
	}

	return taken, rest
}

func ISplit[Slice ~[]T, T any](xs Slice, ipred func(int, T) bool) (taken, rest Slice) {
	return isplitImpl(xs, ipred, nil, nil)
}

func ISplitEx[Slice ~[]T, T any](mode AllocateMode, xs Slice, ipred func(int, T) bool) (taken, rest Slice) {
	if xs == nil {
		return nil, nil
	}

	taken = allocateCapacity[Slice](mode, len(xs), func() int { return ICount(xs, ipred) })
	rest = allocateCapacity[Slice](mode, len(xs), func() int { return len(xs) - len(taken) })

	return isplitImpl(xs, ipred, taken, rest)
}

func isplitImpl[Slice ~[]T, T any](xs Slice, ipred func(int, T) bool, taken, rest Slice) (Slice, Slice) {
	for i, x := range xs {
		if ipred(i, x) {
			taken = append(taken, x)
		} else {
			rest = append(rest, x)
		}
	}

	return taken, rest
}
