package parallel

import "context"

type options struct {
	concurrency int
	ctx         context.Context
}

// ErrOption configures Err-variant parallel functions (MapErr, ForEachErr).
type ErrOption interface {
	apply(*options)
}

// Option configures all parallel functions. It extends ErrOption with a
// cannotFail marker, meaning the option does not require error reporting
// to work correctly. Options like WithContext that need an error return
// only implement ErrOption and will be rejected at compile time by
// non-Err functions.
type Option interface {
	ErrOption
	cannotFail()
}

// ConcurrencyOpt is the option returned by WithConcurrency.
type ConcurrencyOpt int

func (c ConcurrencyOpt) apply(o *options) { o.concurrency = int(c) }
func (c ConcurrencyOpt) cannotFail()      {}

// WithConcurrency sets the maximum number of goroutines running at the same time.
// A value <= 0 means unbounded (one goroutine per item).
func WithConcurrency(n int) ConcurrencyOpt { return ConcurrencyOpt(n) }

// ContextOpt is the option returned by WithContext.
type ContextOpt struct{ ctx context.Context }

func (c ContextOpt) apply(o *options) { o.ctx = c.ctx }

// WithContext sets a context for cancellation support.
// When the context is cancelled, no new items are processed and the first
// cancellation error is returned. Only usable with Err variants (MapErr,
// ForEachErr); passing it to non-Err functions is a compile error.
func WithContext(ctx context.Context) ContextOpt { return ContextOpt{ctx: ctx} }

func buildOptions[O ErrOption](opts []O) options {
	o := options{ctx: context.Background()}
	for _, opt := range opts {
		opt.apply(&o)
	}
	return o
}
