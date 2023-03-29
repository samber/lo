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

// Partial returns new function that, when called, has its last argument set to the provided value.
func PartialRight[T1, T2, R any](f func(a T1, b T2) R, lastArg T2) func(T1) R {
	return func(t1 T1) R {
		return f(t1, lastArg)
	}
}

// PartialRight1 returns new function that, when called, has its last argument set to the provided value.
func PartialRight1[T1, T2, R any](f func(T1, T2) R, lastArg T2) func(T1) R {
	return PartialRight(f, lastArg)
}

// PartialRight2 returns new function that, when called, has its last argument set to the provided value.
func PartialRight2[T1, T2, T3, R any](f func(T1, T2, T3) R, lastArg T3) func(T1, T2) R {
	return func(t1 T1, t2 T2) R {
		return f(t1, t2, lastArg)
	}
}

// PartialRight3 returns new function that, when called, has its last argument set to the provided value.
func PartialRight3[T1, T2, T3, T4, R any](f func(T1, T2, T3, T4) R, lastArg T4) func(T1, T2, T3) R {
	return func(t1 T1, t2 T2, t3 T3) R {
		return f(t1, t2, t3, lastArg)
	}
}

// PartialRight4 returns new function that, when called, has its last argument set to the provided value.
func PartialRight4[T1, T2, T3, T4, T5, R any](f func(T1, T2, T3, T4, T5) R, lastArg T5) func(T1, T2, T3, T4) R {
	return func(t1 T1, t2 T2, t3 T3, t4 T4) R {
		return f(t1, t2, t3, t4, lastArg)
	}
}

// PartialRight5 returns new function that, when called, has its last argument set to the provided value
func PartialRight5[T1, T2, T3, T4, T5, T6, R any](f func(T1, T2, T3, T4, T5, T6) R, lastArg T6) func(T1, T2, T3, T4, T5) R {
	return func(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5) R {
		return f(t1, t2, t3, t4, t5, lastArg)
	}
}

// Not returns a function which is the logical inverse of the provided function
func Not[T any](f func(T) bool) func(T) bool {
	return func(t T) bool {
		return !f(t)
	}
}

// Not2 returns a function which is the logical inverse of the provided function
func Not2[T1, T2 any](f func(T1, T2) bool) func(T1, T2) bool {
	return func(t1 T1, t2 T2) bool {
		return !f(t1, t2)
	}
}

// Do we need to support more than 2 parameters? I struggle to think of meaningful examples.
