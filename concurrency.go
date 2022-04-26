package lo

// Async executes a function in a goroutine and returns the result in a channel.
func Async[T any](f func() T) chan T {
	ch := make(chan T)
	go func() {
		ch <- f()
	}()
	return ch
}

// Async0 executes a function in a goroutine and returns a channel closed after execution.
func Async0(f func()) chan struct{} {
	ch := make(chan struct{})
	go func() {
		f()
		close(ch)
	}()
	return ch
}

// Async1 executes a function in a goroutine and returns the result in a channel.
func Async1[A any](f func() A) chan A {
	return Async(f)
}
