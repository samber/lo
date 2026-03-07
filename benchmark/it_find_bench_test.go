//go:build go1.23

package benchmark

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/samber/lo/it"
)

func BenchmarkItFind(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_, _ = it.Find(ints, func(x int) bool { return x == 0 })
			}
		})
	}
}

func BenchmarkItContains(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Contains(ints, 42)
			}
		})
	}
}

func BenchmarkItContainsBy(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			target := rand.IntN(100_000)
			for range b.N {
				_ = it.ContainsBy(ints, func(x int) bool { return x == target })
			}
		})
	}
}

func BenchmarkItEvery(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Every(ints, 1, 2, 3)
			}
		})
	}
}

func BenchmarkItEveryBy(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.EveryBy(ints, func(x int) bool { return x >= 0 })
			}
		})
	}
}

func BenchmarkItSome(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Some(ints, 1, 2, 3)
			}
		})
	}
}

func BenchmarkItSomeBy(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			target := rand.IntN(100_000)
			for range b.N {
				_ = it.SomeBy(ints, func(x int) bool { return x == target })
			}
		})
	}
}

func BenchmarkItNone(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.None(ints, -1, -2, -3)
			}
		})
	}
}

func BenchmarkItNoneBy(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			target := rand.IntN(100_000)
			for range b.N {
				_ = it.NoneBy(ints, func(x int) bool { return x == target })
			}
		})
	}
}

func BenchmarkItIntersect(b *testing.B) {
	for _, n := range itLengths {
		a := genInts(n)
		c := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Intersect(a, c) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItUnion(b *testing.B) {
	for _, n := range itLengths {
		a := genInts(n)
		c := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Union(a, c) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItWithout(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Without(ints, 1, 2, 3, 4, 5) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItWithoutNth(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.WithoutNth(ints, 0, n/2, n-1) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItIndexOf(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.IndexOf(ints, -1)
			}
		})
	}
}

func BenchmarkItLastIndexOf(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.LastIndexOf(ints, -1)
			}
		})
	}
}

func BenchmarkItHasPrefix(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.HasPrefix(ints, -1, -2, -3)
			}
		})
	}
}

func BenchmarkItHasSuffix(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.HasSuffix(ints, -1, -2, -3)
			}
		})
	}
}

func BenchmarkItFindIndexOf(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_, _, _ = it.FindIndexOf(ints, func(x int) bool { return x == -1 })
			}
		})
	}
}

func BenchmarkItFindOrElse(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.FindOrElse(ints, -1, func(x int) bool { return x == -1 })
			}
		})
	}
}

func BenchmarkItFindUniques(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.FindUniques(ints) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItFindDuplicates(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.FindDuplicates(ints) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItMin(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Min(ints)
			}
		})
	}
}

func BenchmarkItMax(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Max(ints)
			}
		})
	}
}

func BenchmarkItMinBy(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.MinBy(ints, func(a, b int) bool { return a < b })
			}
		})
	}
}

func BenchmarkItMaxBy(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.MaxBy(ints, func(a, b int) bool { return a > b })
			}
		})
	}
}

func BenchmarkItFirst(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_, _ = it.First(ints)
			}
		})
	}
}

func BenchmarkItLast(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_, _ = it.Last(ints)
			}
		})
	}
}

func BenchmarkItNth(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_, _ = it.Nth(ints, n/2)
			}
		})
	}
}

func BenchmarkItSample(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Sample(ints)
			}
		})
	}
}

func BenchmarkItSamples(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Samples(ints, 5) { //nolint:revive
				}
			}
		})
	}
}
