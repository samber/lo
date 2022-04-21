package lo

// MustOk has the same behavior than Must, but receives bool instead of error.
func MustOk[T any](val T, ok bool) T {
	if !ok{
		panic("not ok")
	}

	return val
}

// MustOk0 has the same behavior than MustOk, but callback returns no variable.
func MustOk0(ok bool) {
	if !ok {
		panic("not ok")
	}
}

// MustOk1 is an alias to MustOk
func MustOk1[T any](val T, ok bool) T {
	return MustOk(val, ok)
}

// MustOk2 has the same behavior than MustOk, but callback returns 2 variable.
func MustOk2[T1 any, T2 any](val1 T1, val2 T2, ok bool) (T1, T2) {
	if !ok {
		panic("not ok")
	}

	return val1, val2
}

// MustOk3 has the same behavior than MustOk, but callback returns 3 variable.
func MustOk3[T1 any, T2 any, T3 any](val1 T1, val2 T2, val3 T3, ok bool) (T1, T2, T3) {
	if !ok {
		panic("not ok")
	}

	return val1, val2, val3
}

// MustOk4 has the same behavior than MustOk, but callback returns 4 variable.
func MustOk4[T1 any, T2 any, T3 any, T4 any](val1 T1, val2 T2, val3 T3, val4 T4, ok bool) (T1, T2, T3, T4) {
	if !ok {
		panic("not ok")
	}

	return val1, val2, val3, val4
}

// MustOk5 has the same behavior than MustOk, but callback returns 5 variable.
func MustOk5[T1 any, T2 any, T3 any, T4 any, T5 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5, ok bool) (T1, T2, T3, T4, T5) {
	if !ok {
		panic("not ok")
	}

	return val1, val2, val3, val4, val5
}

// MustOk6 has the same behavior than MustOk, but callback returns 6 variable.
func MustOk6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5, val6 T6, ok bool) (T1, T2, T3, T4, T5, T6) {
	if !ok {
		panic("not ok")
	}

	return val1, val2, val3, val4, val5, val6
}
