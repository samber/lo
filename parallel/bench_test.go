package parallel

import (
	"context"
	"testing"
)

func BenchmarkMap(b *testing.B) {
	items := make([]int, 100)
	for i := range items {
		items[i] = i
	}
	for b.Loop() {
		Map(items, func(x, _ int) int { return x * 2 })
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
