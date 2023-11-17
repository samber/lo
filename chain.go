package lo

// Middleware represents a function that takes an input of type T and returns an output of type T.
// It is used as a building block for creating middleware chains.
type Middleware[T any] func(T) T

// Chain creates a new middleware chain by combining the provided middlewares.
// The outer middleware is the first one to be executed, followed by the rest of the middlewares in the order they are provided.
// The resulting middleware chain takes an input of type T and returns an output of type T.
// If no middlewares are provided, the outer middleware is returned as is.
// Each middleware in the chain is invoked with the output of the previous middleware as its input.
// The final output of the chain is the result of invoking the outer middleware with the input.
func Chain[T any](outer Middleware[T], middlewares ...Middleware[T]) Middleware[T] {
	if len(middlewares) == 0 {
		return outer
	}

	return func(t T) T {
		return Chain(middlewares[0], middlewares[1:]...)(outer(t))
	}
}
