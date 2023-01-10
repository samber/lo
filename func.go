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

// Compose returns new function that, when called, returns the result of calling g and then f with the result from g
func Compose[T, U, V any](f func(U) V, g func(T) U) func (T) V {
	return func(t T) V {
		return f(g(t))
	}
}

// Pipe returns new function that, when called, returns the result of calling f and then g with the result from f
func Pipe[T, U, V any](f func(T) U, g func(U) V) func (T) V {
	return Compose(g, f)
}
