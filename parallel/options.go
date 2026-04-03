package parallel

import "context"

type options struct {
	concurrency int
	ctx         context.Context
}

// Option configures parallel execution.
type Option func(*options)

// WithConcurrency sets the maximum number of goroutines running at the same time.
// A value <= 0 means unbounded (one goroutine per item).
func WithConcurrency(n int) Option {
	return func(o *options) {
		o.concurrency = n
	}
}

// WithContext sets a context for cancellation support.
// When the context is cancelled, no new items are processed and the first
// cancellation error is returned. Only applies to Err variants.
func WithContext(ctx context.Context) Option {
	return func(o *options) {
		o.ctx = ctx
	}
}

func buildOptions(opts []Option) options {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}
