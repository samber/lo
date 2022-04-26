package lo

// Async executes a function in a goroutine and returns the result in a channel.
func Async[T any](f func() T) chan T {
	ch := make(chan T)
	go func() {
		ch <- f()
	}()
	return ch
}

// Async0 executes a function in a goroutine and returns a channel set once the function finishes.
func Async0(f func()) chan struct{} {
	ch := make(chan struct{})
	go func() {
		f()
		ch <- struct{}{}
	}()
	return ch
}

// Async1 is an alias to Async.
func Async1[A any](f func() A) chan A {
	return Async(f)
}
