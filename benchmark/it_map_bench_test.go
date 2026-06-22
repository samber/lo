//go:build go1.23

package benchmark

import (
	"fmt"
	"strings"
	"testing"

	"github.com/samber/lo/it"
)

func BenchmarkItKeys(b *testing.B) {
	for _, n := range itLengths {
		m := genMapStringInt(n)
		b.Run(fmt.Sprintf("map_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Keys(m) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItUniqKeys(b *testing.B) {
	for _, n := range itLengths {
		m := genMapStringInt(n)
		b.Run(fmt.Sprintf("map_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.UniqKeys(m) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItValues(b *testing.B) {
	for _, n := range itLengths {
		m := genMapStringInt(n)
		b.Run(fmt.Sprintf("map_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Values(m) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItUniqValues(b *testing.B) {
	for _, n := range itLengths {
		m := genMapStringInt(n)
		b.Run(fmt.Sprintf("map_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.UniqValues(m) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItEntries(b *testing.B) {
	for _, n := range itLengths {
		m := genMapStringInt(n)
		b.Run(fmt.Sprintf("map_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Entries(m) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItFromEntries(b *testing.B) {
	for _, n := range itLengths {
		m := genMapStringInt(n)
		entries := it.Entries(m)
		b.Run(fmt.Sprintf("map_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.FromEntries(entries)
			}
		})
	}
}

func BenchmarkItInvert(b *testing.B) {
	for _, n := range itLengths {
		m := genMapStringInt(n)
		entries := it.Entries(m)
		b.Run(fmt.Sprintf("map_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Invert(entries) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItAssign(b *testing.B) {
	for _, n := range itLengths {
		seq := genMapSeq(n)
		b.Run(fmt.Sprintf("maps_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Assign(seq)
			}
		})
	}
}

func BenchmarkItFilterKeys(b *testing.B) {
	for _, n := range itLengths {
		m := genMapStringInt(n)
		b.Run(fmt.Sprintf("map_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.FilterKeys(m, func(_ string, v int) bool { return v%2 == 0 }) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItFilterValues(b *testing.B) {
	for _, n := range itLengths {
		m := genMapStringInt(n)
		b.Run(fmt.Sprintf("map_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.FilterValues(m, func(_ string, v int) bool { return v%2 == 0 }) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItChunkString(b *testing.B) {
	for _, n := range itLengths {
		var sb strings.Builder
		for range n {
			sb.WriteString("a")
		}
		s := sb.String()
		b.Run(fmt.Sprintf("len_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.ChunkString(s, 5) { //nolint:revive
				}
			}
		})
	}
}
