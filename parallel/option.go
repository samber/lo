package parallel

// ParallelOption provides some useful configure options for parallel, such as conurrency limit.
type ParallelOption struct {
	concurrency       int
	concurrencySetted bool
}

// Option() create an empty ParallelOption
func Option() *ParallelOption {
	return &ParallelOption{}
}

// Concurrency() set the maximum number of concurrent `iteratee` goroutines running at the same time.
func (o *ParallelOption) Concurrency(concurrency int) *ParallelOption {
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
