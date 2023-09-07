package lo

// Partial returns new function that, when called, has its first argument set to the provided value.
func Partial[T1, T2, R any](f func(a T1, b T2) R, arg1 T1) func(T2) R {
	return func(t2 T2) R {
		return f(arg1, t2)
	}
}

// Partial1 returns new function that, when called, has its first argument set to the provided value.
func Partial1[T1, T2, R any](f func(T1, T2) R, arg1 T1) func(T2) R {
	return Partial(f, arg1)
}

// Partial2 returns new function that, when called, has its first argument set to the provided value.
func Partial2[T1, T2, T3, R any](f func(T1, T2, T3) R, arg1 T1) func(T2, T3) R {
	return func(t2 T2, t3 T3) R {
		return f(arg1, t2, t3)
	}
}

// Partial3 returns new function that, when called, has its first argument set to the provided value.
func Partial3[T1, T2, T3, T4, R any](f func(T1, T2, T3, T4) R, arg1 T1) func(T2, T3, T4) R {
	return func(t2 T2, t3 T3, t4 T4) R {
		return f(arg1, t2, t3, t4)
	}
}

// Partial4 returns new function that, when called, has its first argument set to the provided value.
func Partial4[T1, T2, T3, T4, T5, R any](f func(T1, T2, T3, T4, T5) R, arg1 T1) func(T2, T3, T4, T5) R {
	return func(t2 T2, t3 T3, t4 T4, t5 T5) R {
		return f(arg1, t2, t3, t4, t5)
	}
}

// Partial5 returns new function that, when called, has its first argument set to the provided value
func Partial5[T1, T2, T3, T4, T5, T6, R any](f func(T1, T2, T3, T4, T5, T6) R, arg1 T1) func(T2, T3, T4, T5, T6) R {
	return func(t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) R {
		return f(arg1, t2, t3, t4, t5, t6)
	}
}

// Compose returns new function that is the composition of 2 functions passed as arguments
func Compose[T1, T2, T3 any](f1 func(T1) T2, f2 func(T2) T3) func(T1) T3 {
	return func(t1 T1) T3 {
		return f2(f1(t1))
	}
}

// Compose2 returns new function that is the composition of 2 functions passed as arguments
func Compose2[T1, T2, T3 any](f1 func(T1) T2, f2 func(T2) T3) func(T1) T3 {
	return Compose(f1, f2)
}

// Compose3 returns new function that is the composition of 3 functions passed as arguments
func Compose3[T1, T2, T3, T4 any](f1 func(T1) T2, f2 func(T2) T3, f3 func(T3) T4) func(T1) T4 {
	return Compose(Compose(f1, f2), f3)
}

// Compose4 returns new function that is the composition of 4 functions passed as arguments
func Compose4[T1, T2, T3, T4, T5 any](f1 func(T1) T2, f2 func(T2) T3, f3 func(T3) T4, f4 func(T4) T5) func(T1) T5 {
	return Compose(Compose(Compose(f1, f2), f3), f4)
}

// Compose5 returns new function that is the composition of 5 functions passed as arguments
func Compose5[T1, T2, T3, T4, T5, T6 any](f1 func(T1) T2, f2 func(T2) T3, f3 func(T3) T4, f4 func(T4) T5, f5 func(T5) T6) func(T1) T6 {
	return Compose(Compose(Compose(Compose(f1, f2), f3), f4), f5)
}
