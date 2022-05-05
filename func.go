package lo

// Bind returns new function that, when called, has its first argument set to the provided value.
func Bind[T1, T2, R any](f func(T1, T2) R, arg1 T1) func(T2) R {
	return func(t2 T2) R {
		return f(arg1, t2)
	}
}
