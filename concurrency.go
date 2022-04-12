package lo

// Async executes a function in a goroutine and returns the result in a channel.
func Async[T any](f func() T) chan T {
	ch := make(chan T)
	go func() {
		ch <- f()
	}()
	return ch
}
