package parallel

type ParallelOption struct {
	concurrency       int
	concurrencySetted bool
}

func Option() *ParallelOption {
	return &ParallelOption{}
}

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
