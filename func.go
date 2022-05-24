package lo

// Pipeline takes a list of functions and returns a function
// that takes a value as its argument and runs it through
// a pipeline of the original functions given in this function.
func Pipeline[T any](funcs ...func(T) T) func(T) T {
	return func(t T) (result T) {
		result = t
		for _, f := range funcs {
			result = f(result)
		}
		return
	}
}
