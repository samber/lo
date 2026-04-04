package parallel

import "testing"

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
