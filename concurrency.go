package lo

// Async executes a function in a goroutine and returns the result in a channel.
func Async[A any](f func() A) chan A {
	ch := make(chan A)
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

// Async2 has the same behavior as Async, but returns the 2 results as a tuple inside the channel.
func Async2[A any, B any](f func() (A, B)) chan Tuple2[A, B] {
	ch := make(chan Tuple2[A, B])
	go func() {
		ch <- T2(f())
	}()
	return ch
}

// Async3 has the same behavior as Async, but returns the 3 results as a tuple inside the channel.
func Async3[A any, B any, C any](f func() (A, B, C)) chan Tuple3[A, B, C] {
	ch := make(chan Tuple3[A, B, C])
	go func() {
		ch <- T3(f())
	}()
	return ch
}

// Async4 has the same behavior as Async, but returns the 4 results as a tuple inside the channel.
func Async4[A any, B any, C any, D any](f func() (A, B, C, D)) chan Tuple4[A, B, C, D] {
	ch := make(chan Tuple4[A, B, C, D])
	go func() {
		ch <- T4(f())
	}()
	return ch
}

// Async5 has the same behavior as Async, but returns the 5 results as a tuple inside the channel.
func Async5[A any, B any, C any, D any, E any](f func() (A, B, C, D, E)) chan Tuple5[A, B, C, D, E] {
	ch := make(chan Tuple5[A, B, C, D, E])
	go func() {
		ch <- T5(f())
	}()
	return ch
}

// Async6 has the same behavior as Async, but returns the 6 results as a tuple inside the channel.
func Async6[A any, B any, C any, D any, E any, F any](f func() (A, B, C, D, E, F)) chan Tuple6[A, B, C, D, E, F] {
	ch := make(chan Tuple6[A, B, C, D, E, F])
	go func() {
		ch <- T6(f())
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
