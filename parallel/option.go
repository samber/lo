package parallel

type ParallelOption struct {
	concurrency int
}

func Option() *ParallelOption {
	return &ParallelOption{}
}

func (o *ParallelOption) Concurrency(concurrency int) *ParallelOption {
	o.concurrency = concurrency
	return o
}
