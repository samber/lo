package parallel

// Option provides some useful configure options for parallel, such as conurrency limit.
type Option struct {
	concurrency       int
	concurrencySetted bool
}

// WithConcurrency create an ParallelOption and set the maximum number of concurrent `iteratee` goroutines running
// at the same time.
func WithConcurrency(concurrency int) *Option {
	return NewOption().WithConcurrency(concurrency)
}

// NewOption() create an empty ParallelOption
func NewOption() *Option {
	return &Option{}
}

// WithConcurrency() set the maximum number of concurrent `iteratee` goroutines running at the same time.
func (o *Option) WithConcurrency(concurrency int) *Option {
	o.concurrency = concurrency
	o.concurrencySetted = true
	return o
}

func mergeOptions(options []*Option) Option {
	ret := Option{}
	for _, option := range options {
		if option == nil {
			continue
		}
		if option.concurrencySetted {
			ret.concurrency = option.concurrency
			ret.concurrencySetted = true
		}
	}
	return ret
}
