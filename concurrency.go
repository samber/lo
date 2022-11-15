package lo

import "sync"

type synchronize struct {
	locker sync.Locker
}

func (s *synchronize) Do(cb func()) {
	s.locker.Lock()
	Try0(cb)
	s.locker.Unlock()
}

// Synchronize wraps the underlying callback in a mutex. It receives an optional mutex.
func Synchronize(opt ...sync.Locker) *synchronize {
	if len(opt) > 1 {
		panic("unexpected arguments")
	} else if len(opt) == 0 {
		opt = append(opt, &sync.Mutex{})
	}

	return &synchronize{
		locker: opt[0],
	}
}

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

type rpc[T any, R any] struct {
	C chan Tuple2[T, func(R)]
}

// NewRPC synchronizes goroutines for a bidirectionnal request-response communication.
func NewRPC[T any, R any](ch chan<- T) *rpc[T, R] {
	return &rpc[T, R]{
		C: make(chan Tuple2[T, func(R)]),
	}
}

// Send blocks until response is triggered.
func (rpc *rpc[T, R]) Send(request T) R {
	done := make(chan R)
	defer close(done)

	once := sync.Once{}

	rpc.C <- T2(request, func(response R) {
		once.Do(func() {
			done <- response
		})
	})

	return <-done
}
