//go:build go1.23

package benchmark

import (
	"fmt"
	"iter"
	"math/rand/v2"
	"strconv"
	"strings"
	"testing"

	"github.com/samber/lo/it"
)

var itLengths = []int{10, 100, 1000}

func BenchmarkItChunk(b *testing.B) {
	for _, n := range itLengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Chunk(strs, 5) { //nolint:revive
				}
			}
		})
	}

	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Chunk(ints, 5) { //nolint:revive
				}
			}
		})
	}
}

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

func BenchmarkItFlatten(b *testing.B) {
	for _, n := range itLengths {
		ints := make([]iter.Seq[int], 0, n)
		for range n {
			ints = append(ints, genInts(n))
		}
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Flatten(ints) { //nolint:revive
				}
			}
		})
	}

	for _, n := range itLengths {
		strs := make([]iter.Seq[string], 0, n)
		for range n {
			strs = append(strs, genStrings(n))
		}
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Flatten(strs) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItDrop(b *testing.B) {
	for _, n := range itLengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Drop(strs, n/4) { //nolint:revive
				}
			}
		})
	}

	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Drop(ints, n/4) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItDropWhile(b *testing.B) {
	for _, n := range itLengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.DropWhile(strs, func(v string) bool { return len(v) < 4 }) { //nolint:revive
				}
			}
		})
	}

	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := range b.N {
				for range it.DropWhile(ints, func(v int) bool { return i < 10_000 }) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItDropLastWhile(b *testing.B) {
	for _, n := range itLengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.DropLastWhile(strs, func(v string) bool { return len(v) < 4 }) { //nolint:revive
				}
			}
		})
	}

	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for range b.N {
				for range it.DropLastWhile(ints, func(v int) bool { return v < 10_000 }) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItDropByIndex(b *testing.B) {
	for _, n := range itLengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.DropByIndex(strs, n/4) { //nolint:revive
				}
			}
		})
	}

	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for range b.N {
				for range it.DropByIndex(ints, n/4) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItReplace(b *testing.B) {
	lengths := []int{1_000, 10_000, 100_000}

	for _, n := range lengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Replace(strs, "321321", "123123", 10) { //nolint:revive
				}
			}
		})
	}

	for _, n := range lengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Replace(ints, 321321, 123123, 10) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItTrim(b *testing.B) {
	for _, n := range itLengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Trim(strs, "123", "456") { //nolint:revive
				}
			}
		})
	}

	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Trim(ints, 123, 456) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItTrimSuffix(b *testing.B) {
	for _, n := range itLengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.TrimSuffix(strs, []string{""}) { //nolint:revive
				}
			}
		})
	}

	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for range b.N {
				for range it.TrimSuffix(ints, []int{0}) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItCountBy(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.CountBy(ints, func(x int) bool { return x%2 == 0 })
			}
		})
	}
}

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

func BenchmarkItFilter(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Filter(ints, func(x int) bool { return x%2 == 0 }) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItMap(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Map(ints, func(x int) int { return x * 2 }) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItUniqMap(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.UniqMap(ints, func(x int) int { return x % 50 }) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItFilterMap(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.FilterMap(ints, func(x int) (int, bool) { return x * 2, x%2 == 0 }) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItFlatMap(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.FlatMap(ints, func(x int) iter.Seq[int] { //nolint:revive
					return func(yield func(int) bool) {
						yield(x)
					}
				}) {
				}
			}
		})
	}
}

func BenchmarkItReduce(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Reduce(ints, func(agg, item int) int { return agg + item }, 0)
			}
		})
	}
}

func BenchmarkItForEach(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				it.ForEach(ints, func(_ int) {})
			}
		})
	}
}

func BenchmarkItForEachWhile(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				it.ForEachWhile(ints, func(_ int) bool { return true })
			}
		})
	}
}

func BenchmarkItTimes(b *testing.B) {
	for _, n := range itLengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Times(n, func(i int) int { return i * 2 }) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItUniq(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Uniq(ints) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItUniqBy(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.UniqBy(ints, func(x int) int { return x % 50 }) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItGroupBy(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.GroupBy(ints, func(x int) int { return x % 10 })
			}
		})
	}
}

func BenchmarkItPartitionBy(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.PartitionBy(ints, func(x int) int { return x % 10 })
			}
		})
	}
}

func BenchmarkItConcat(b *testing.B) {
	for _, n := range itLengths {
		a := genInts(n)
		c := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Concat(a, c) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItInterleave(b *testing.B) {
	for _, n := range itLengths {
		a := genInts(n)
		c := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Interleave(a, c) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItShuffle(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Shuffle(ints) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItReverse(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Reverse(ints) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItRepeatBy(b *testing.B) {
	for _, n := range itLengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.RepeatBy(n, func(i int) int { return i * 2 }) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItKeyBy(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.KeyBy(ints, func(x int) int { return x })
			}
		})
	}
}

func BenchmarkItAssociate(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Associate(ints, func(x int) (int, int) { return x, x * 2 })
			}
		})
	}
}

func BenchmarkItTake(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Take(ints, n/2) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItTakeWhile(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.TakeWhile(ints, func(x int) bool { return x < 90_000 }) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItTakeFilter(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.TakeFilter(ints, 5, func(x int) bool { return x%2 == 0 }) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItReject(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Reject(ints, func(x int) bool { return x%2 == 0 }) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItRejectMap(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.RejectMap(ints, func(x int) (int, bool) { return x * 2, x%2 == 0 }) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItCount(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Count(ints, 42)
			}
		})
	}
}

func BenchmarkItCountValues(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.CountValues(ints)
			}
		})
	}
}

func BenchmarkItCountValuesBy(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.CountValuesBy(ints, func(x int) int { return x % 10 })
			}
		})
	}
}

func BenchmarkItSubset(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Subset(ints, n/4, n/2) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItSlice(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Slice(ints, n/4, n*3/4) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItReplaceAll(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.ReplaceAll(ints, 42, 99) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItCompact(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Compact(ints) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItIsSorted(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.IsSorted(ints)
			}
		})
	}
}

func BenchmarkItSplice(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Splice(ints, n/2, 1, 2, 3) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItCutPrefix(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				after, _ := it.CutPrefix(ints, []int{-1, -2})
				for range after { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItBuffer(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Buffer(ints, 5) { //nolint:revive
				}
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

func BenchmarkItSum(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Sum(ints)
			}
		})
	}
}

func BenchmarkItSumBy(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.SumBy(ints, func(x int) int { return x })
			}
		})
	}
}

func BenchmarkItProduct(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Product(ints)
			}
		})
	}
}

func BenchmarkItMean(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Mean(ints)
			}
		})
	}
}

func BenchmarkItMode(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Mode(ints)
			}
		})
	}
}

func BenchmarkItRange(b *testing.B) {
	for _, n := range itLengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Range(n) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItLength(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Length(ints)
			}
		})
	}
}

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

func BenchmarkItToSeqPtr(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.ToSeqPtr(ints) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItFromSeqPtr(b *testing.B) {
	for _, n := range itLengths {
		ptrs := genIntPtrSeq(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.FromSeqPtr(ptrs) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItToAnySeq(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.ToAnySeq(ints) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItChunkString(b *testing.B) {
	for _, n := range itLengths {
		s := ""
		var sSb1288 strings.Builder
		for range n {
			sSb1288.WriteString("a")
		}
		s += sSb1288.String()
		b.Run(fmt.Sprintf("len_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.ChunkString(s, 5) { //nolint:revive
				}
			}
		})
	}
}
