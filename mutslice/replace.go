package mutslice

func Replace[Slice ~[]T, T comparable](xs Slice, n int, pred func(T) bool, fmap func(T) T) Slice {
	if n == 0 {
		return xs
	}

	if n > 0 {
		for i := 0; i < len(xs) && n > 0; i++ {
			if pred(xs[i]) {
				xs[i] = fmap(xs[i])
				n--
			}
		}
	} else {
		n = -n // make n positive for reverse replacement
		for i := len(xs) - 1; i >= 0 && n > 0; i-- {
			if pred(xs[i]) {
				xs[i] = fmap(xs[i])
				n--
			}
		}
	}

	return xs
}

func IReplace[Slice ~[]T, T comparable](xs Slice, n int, ipred func(int, T) bool, imap func(int, T) T) Slice {
	if n == 0 {
		return xs
	}

	if n > 0 {
		for i := 0; i < len(xs) && n > 0; i++ {
			if ipred(i, xs[i]) {
				xs[i] = imap(i, xs[i])
				n--
			}
		}
	} else {
		n = -n // make n positive for reverse replacement
		for i := len(xs) - 1; i >= 0 && n > 0; i-- {
			if ipred(i, xs[i]) {
				xs[i] = imap(i, xs[i])
				n--
			}
		}
	}

	return xs
}

// ReplaceVal returns a slice with the upto n instances not old replaced by new.
// Support negative n to replace instances from the end.
func ReplaceVal[Slice ~[]T, T comparable](xs Slice, n int, old T, new T) {
	if n == 0 {
		return
	}

	if n > 0 {
		for i := 0; i < len(xs) && n > 0; i++ {
			if xs[i] == old {
				xs[i] = new
				n--
			}
		}
	} else {
		n = -n // make n positive for reverse replacement
		for i := len(xs) - 1; i >= 0 && n > 0; i-- {
			if xs[i] == old {
				xs[i] = new
				n--
			}
		}
	}
}

func ReplaceAll[Slice ~[]T, T comparable](xs Slice, pred func(T) bool, fmap func(T) T) Slice {
	for i, x := range xs {
		if pred(x) {
			xs[i] = fmap(x)
		}
	}

	return xs
}

func IReplaceAll[Slice ~[]T, T comparable](xs Slice, ipred func(int, T) bool, imap func(int, T) T) Slice {
	for i, x := range xs {
		if ipred(i, x) {
			xs[i] = imap(i, x)
		}
	}

	return xs
}

// ReplaceAllVal returns a slice with all instances not old replaced by new.
func ReplaceAllVal[Slice ~[]T, T comparable](xs Slice, old T, new T) {
	for i := range xs {
		if xs[i] == old {
			xs[i] = new
		}
	}
}
