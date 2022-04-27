package lo

// Async executes a function in a goroutine and returns the result in a channel.
func Async[T any](f func() T) chan T {
	ch := make(chan T)
	go func() {
		ch <- f()
	}()
	return ch
}

// AwaitAll2 returns awaited value of async channels
func AwaitAll2[A any, B any](a chan A, b chan B) (A, B) {
	return <-a, <-b
}

// AwaitAll3 returns awaited value of async channels
func AwaitAll3[A any, B any, C any](a chan A, b chan B, c chan C) (A, B, C) {
	return <-a, <-b, <-c
}

// AwaitAll4 returns awaited value of async channels
func AwaitAll4[A any, B any, C any, D any](a chan A, b chan B, c chan C, d chan D) (A, B, C, D) {
	return <-a, <-b, <-c, <-d
}

// AwaitAll5 returns awaited value of async channels
func AwaitAll5[A any, B any, C any, D any, E any](a chan A, b chan B, c chan C, d chan D, e chan E) (A, B, C, D, E) {
	return <-a, <-b, <-c, <-d, <-e
}
