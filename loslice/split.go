package loslice

func Split[T any, Slice ~[]T](xs Slice, pred func(T) bool) (taken, rest Slice) {
	taken = make(Slice, 0, Count(xs, pred))
	rest = make(Slice, 0, len(xs)-cap(taken))

	for _, x := range xs {
		if pred(x) {
			taken = append(taken, x)
		} else {
			rest = append(rest, x)
		}
	}

	return taken, rest
}

func ISplit[T any, Slice ~[]T](xs Slice, ipred func(int, T) bool) (taken, rest Slice) {
	taken = make(Slice, 0, ICount(xs, ipred))
	rest = make(Slice, 0, len(xs)-cap(taken))

	for i, x := range xs {
		if ipred(i, x) {
			taken = append(taken, x)
		} else {
			rest = append(rest, x)
		}
	}

	return taken, rest
}
