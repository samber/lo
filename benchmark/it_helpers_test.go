//go:build go1.23

package benchmark

import (
	"iter"
	"math/rand/v2"
	"strconv"
)

var itLengths = []int{10, 100, 1000}

func genStrings(n int) iter.Seq[string] {
	return func(yield func(string) bool) {
		for range n {
			if !yield(strconv.Itoa(rand.IntN(100_000))) {
				break
			}
		}
	}
}

func genInts(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for range n {
			if !yield(rand.IntN(100_000)) {
				break
			}
		}
	}
}

func genIntPtrSeq(n int) iter.Seq[*int] {
	return func(yield func(*int) bool) {
		for range n {
			v := rand.IntN(100_000)
			if !yield(&v) {
				break
			}
		}
	}
}

func genMapStringInt(n int) map[string]int {
	m := make(map[string]int, n)
	for i := range n {
		m[strconv.Itoa(i)] = rand.IntN(100_000)
	}
	return m
}

func genMapSeq(n int) iter.Seq[map[string]int] {
	return func(yield func(map[string]int) bool) {
		for range n {
			m := map[string]int{strconv.Itoa(rand.IntN(100_000)): rand.IntN(100_000)}
			if !yield(m) {
				break
			}
		}
	}
}
