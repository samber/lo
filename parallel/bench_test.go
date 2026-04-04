package parallel

import (
	"context"
	"sync"
	"testing"
)

// Upstream master implementations inlined for direct comparison.
// Run `go test -bench=. ./parallel/` to see them side by side.

func masterMap(items []int, transform func(int, int) int) []int {
	result := make([]int, len(items))
	var wg sync.WaitGroup
	wg.Add(len(items))
	for i, item := range items {
		go func(_item, _i int) {
			result[_i] = transform(_item, _i)
			wg.Done()
		}(item, i)
	}
	wg.Wait()
	return result
}

func masterForEach(items []int, callback func(int, int)) {
	var wg sync.WaitGroup
	wg.Add(len(items))
	for i, item := range items {
		go func(_item, _i int) {
			callback(_item, _i)
			wg.Done()
		}(item, i)
	}
	wg.Wait()
}

func masterTimes(count int, iteratee func(int) int) []int {
	result := make([]int, count)
	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(_i int) {
			result[_i] = iteratee(_i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return result
}

func BenchmarkMap_Master(b *testing.B) {
	items := make([]int, 100)
	for i := range items {
		items[i] = i
	}
	for b.Loop() {
		masterMap(items, func(x, _ int) int { return x * 2 })
	}
}

func BenchmarkMap(b *testing.B) {
	items := make([]int, 100)
	for i := range items {
		items[i] = i
	}
	for b.Loop() {
		Map(items, func(x, _ int) int { return x * 2 })
	}
}

func BenchmarkForEach_Master(b *testing.B) {
	items := make([]int, 100)
	for i := range items {
		items[i] = i
	}
	for b.Loop() {
		masterForEach(items, func(_, _ int) {})
	}
}

func BenchmarkForEach(b *testing.B) {
	items := make([]int, 100)
	for i := range items {
		items[i] = i
	}
	for b.Loop() {
		ForEach(items, func(_, _ int) {})
	}
}

func BenchmarkTimes_Master(b *testing.B) {
	for b.Loop() {
		masterTimes(100, func(i int) int { return i * 2 })
	}
}

func BenchmarkTimes(b *testing.B) {
	for b.Loop() {
		Times(100, func(i int) int { return i * 2 })
	}
}

func BenchmarkMapErr_Unbounded(b *testing.B) {
	items := make([]int, 100)
	for i := range items {
		items[i] = i
	}
	for b.Loop() {
		_, _ = MapErr(items, func(x, _ int) (int, error) { return x * 2, nil })
	}
}

func BenchmarkMapErr_UnboundedWithContext(b *testing.B) {
	items := make([]int, 100)
	for i := range items {
		items[i] = i
	}
	ctx := context.Background()
	for b.Loop() {
		_, _ = MapErr(items, func(x, _ int) (int, error) { return x * 2, nil }, WithContext(ctx))
	}
}

func BenchmarkMapErr_Bounded(b *testing.B) {
	items := make([]int, 100)
	for i := range items {
		items[i] = i
	}
	for b.Loop() {
		_, _ = MapErr(items, func(x, _ int) (int, error) { return x * 2, nil }, WithConcurrency(10))
	}
}
