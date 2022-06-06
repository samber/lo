package parallel

// ParallelOption provides some useful configure options for parallel, such as conurrency limit.
type ParallelOption struct {
	concurrency       int
	concurrencySetted bool
}

// WithConcurrency create an ParallelOption and set the maximum number of concurrent `iteratee` goroutines running
// at the same time.
func WithConcurrency(concurrency int) *ParallelOption {
	return Option().WithConcurrency(concurrency)
}

// Option() create an empty ParallelOption
func Option() *ParallelOption {
	return &ParallelOption{}
}

// WithConcurrency() set the maximum number of concurrent `iteratee` goroutines running at the same time.
func (o *ParallelOption) WithConcurrency(concurrency int) *ParallelOption {
	o.concurrency = concurrency
	o.concurrencySetted = true
	return o
}

func mergeOptions(options []*ParallelOption) ParallelOption {
	ret := ParallelOption{}
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
