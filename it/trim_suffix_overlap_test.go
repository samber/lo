//go:build go1.23

package it

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/samber/lo"
)

func TestTrimSuffixOverlappingCandidates(t *testing.T) {
	for n := 0; n <= 10; n++ {
		for bits := 0; bits < 1<<n; bits++ {
			input := make([]int, n)
			for i := range input {
				input[i] = (bits >> i) & 1
			}
			for m := 1; m <= 5; m++ {
				for mask := 0; mask < 1<<m; mask++ {
					suffix := make([]int, m)
					for i := range suffix {
						suffix[i] = (mask >> i) & 1
					}
					got := slices.Collect(TrimSuffix(slices.Values(input), suffix))
					want := lo.TrimSuffix(input, suffix)
					if !slices.Equal(got, want) {
						t.Fatalf("TrimSuffix(%v,%v)=%v; want %v", input, suffix, got, want)
					}
				}
			}
		}
	}
}

func TestTrimSuffixOverlapControls(t *testing.T) {
	tests := []struct{ input, suffix, want string }{
		{"ababa", "aba", "ab"},
		{"aaaaa", "aa", "a"},
		{"ababaaba", "aba", "ab"},
		{"ababax", "aba", "ababax"},
		{"abababa", "aba", "abab"},
		{"aaaaab", "aa", "aaaaab"},
		{"abcababcab", "abcab", ""},
		{"abcababc", "abcab", "abcababc"},
	}
	for _, tt := range tests {
		t.Run(tt.input+"/"+tt.suffix, func(t *testing.T) {
			input, suffix := []byte(tt.input), []byte(tt.suffix)
			seq := TrimSuffix(slices.Values(input), suffix)
			for range 2 {
				if got := string(slices.Collect(seq)); got != tt.want {
					t.Fatalf("got %q; want %q", got, tt.want)
				}
			}
			for limit := 0; limit <= len(tt.want); limit++ {
				source := func(yield func(byte) bool) {
					for _, v := range input {
						if !yield(v) {
							return
						}
					}
				}
				got := slices.Collect(Take(TrimSuffix(source, suffix), limit))
				if string(got) != tt.want[:limit] {
					t.Fatalf("limit %d: got %q; want %q", limit, got, tt.want[:limit])
				}
			}
		})
	}
}

func BenchmarkTrimSuffixOverlap(b *testing.B) {
	for _, tc := range []struct{ name, input, suffix string }{
		{"unmatched", "the quick brown fox jumps over the lazy dog", "xyz"},
		{"repeated", "prefixabababababababababababababababababababababababababababababababab", "ab"},
		{"overlapping", "ababababababababababababababababababababababababababababababababababa", "aba"},
	} {
		b.Run(tc.name, func(b *testing.B) {
			input, suffix := []byte(tc.input), []byte(tc.suffix)
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				for range TrimSuffix(slices.Values(input), suffix) {
				}
			}
		})
	}
}

func TestTrimSuffixOverlapRandomParity(t *testing.T) {
	random := rand.New(rand.NewSource(1))
	for range 10000 {
		input := make([]int, random.Intn(128))
		suffix := make([]int, 1+random.Intn(12))
		for i := range input {
			input[i] = random.Intn(3)
		}
		for i := range suffix {
			suffix[i] = random.Intn(3)
		}
		got := slices.Collect(TrimSuffix(slices.Values(input), suffix))
		want := lo.TrimSuffix(input, suffix)
		if !slices.Equal(got, want) {
			t.Fatalf("TrimSuffix(%v,%v)=%v; want %v", input, suffix, got, want)
		}
	}
}

func TestTrimSuffixOverlapStopsSource(t *testing.T) {
	read, closed := 0, false
	source := func(yield func(byte) bool) {
		defer func() { closed = true }()
		for _, v := range []byte("xabababababa") {
			read++
			if !yield(v) {
				return
			}
		}
	}
	for range TrimSuffix(source, []byte("aba")) {
		break
	}
	if read != 1 || !closed {
		t.Fatalf("read=%d, closed=%v; want 1, true", read, closed)
	}
}
