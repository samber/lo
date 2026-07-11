package benchmark

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/samber/lo"
)

func BenchmarkChunk(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Chunk(strs, 5)
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Chunk(ints, 5)
			}
		})
	}
}

func BenchmarkFlatten(b *testing.B) {
	for _, n := range lengths {
		ints := make([][]int, 0, n)
		for i := 0; i < n; i++ {
			ints = append(ints, genSliceInt(n))
		}
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Flatten(ints)
			}
		})
	}

	for _, n := range lengths {
		strs := make([][]string, 0, n)
		for i := 0; i < n; i++ {
			strs = append(strs, genSliceString(n))
		}
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Flatten(strs)
			}
		})
	}
}

func BenchmarkDrop(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Drop(strs, n/4)
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Drop(ints, n/4)
			}
		})
	}
}

func BenchmarkDropRight(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.DropRight(strs, n/4)
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.DropRight(ints, n/4)
			}
		})
	}
}

func BenchmarkDropWhile(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.DropWhile(strs, func(v string) bool { return len(v) < 4 })
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.DropWhile(ints, func(v int) bool { return i < 10_000 })
			}
		})
	}
}

func BenchmarkDropRightWhile(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.DropRightWhile(strs, func(v string) bool { return len(v) < 4 })
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.DropRightWhile(ints, func(v int) bool { return i < 10_000 })
			}
		})
	}
}

func BenchmarkDropByIndex(b *testing.B) {
	for _, n := range lengths {
		for _, indexes := range [][]int{
			{0},
			{0, n / 2, n / 4, n - 1},
			lo.Range(n),
		} {
			name := fmt.Sprintf("size_%d/indexes_%d/", n, len(indexes))

			strs := genSliceString(n)
			b.Run(name+"strings", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = lo.DropByIndex(strs, indexes...)
				}
			})

			ints := genSliceInt(n)
			b.Run(name+"ints", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = lo.DropByIndex(ints, indexes...)
				}
			})

			heavy := genSliceHeavy(n)
			b.Run(name+"heavy", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = lo.DropByIndex(heavy, indexes...)
				}
			})
		}
	}
}

func BenchmarkReplace(b *testing.B) {
	lengths := []int{1_000, 10_000, 100_000}
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Replace(strs, strs[n/4], "123123", 10)
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Replace(ints, ints[n/4], 123123, 10)
			}
		})
	}
}

func BenchmarkReject(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Reject(strs, func(v string, _ int) bool { return len(v) < 3 })
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Reject(ints, func(v, _ int) bool { return v < 50000 })
			}
		})
	}
}

func BenchmarkRejectErr(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.RejectErr(strs, func(v string, _ int) (bool, error) { return len(v) < 3, nil })
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.RejectErr(ints, func(v, _ int) (bool, error) { return v < 50000, nil })
			}
		})
	}
}

func BenchmarkRejectMap(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.RejectMap(strs, func(v string, _ int) (string, bool) { return v, len(v) < 3 })
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.RejectMap(ints, func(v, _ int) (int, bool) { return v, v < 50000 })
			}
		})
	}
}

func BenchmarkUniqMap(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.UniqMap(ints, func(v, _ int) int { return v % 50 })
			}
		})
	}

	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.UniqMap(strs, func(v string, _ int) string { return v })
			}
		})
	}
}

func BenchmarkRepeatBy(b *testing.B) {
	for _, n := range lengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.RepeatBy(n, func(index int) int { return index * 2 })
			}
		})
	}

	for _, n := range lengths {
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.RepeatBy(n, strconv.Itoa)
			}
		})
	}
}

func BenchmarkFill(b *testing.B) {
	for _, n := range lengths {
		collection := make([]clonableString, n)
		for i := range collection {
			collection[i] = clonableString{strconv.Itoa(i)}
		}
		b.Run(fmt.Sprintf("size_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Fill(collection, clonableString{"hello"})
			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for _, n := range lengths {
		b.Run(fmt.Sprintf("size_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Repeat(n, clonableString{"hello"})
			}
		})
	}
}

func BenchmarkFilter(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Filter(ints, func(v, _ int) bool { return v%2 == 0 })
			}
		})
	}
}

func BenchmarkFilterErr(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.FilterErr(ints, func(v, _ int) (bool, error) { return v%2 == 0, nil })
			}
		})
	}
}

func BenchmarkSliceMap(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Map(ints, func(v, _ int) int { return v * 2 })
			}
		})
	}
}

func BenchmarkMapErr(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.MapErr(ints, func(v, _ int) (int, error) { return v * 2, nil })
			}
		})
	}
}

func BenchmarkFilterMap(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FilterMap(ints, func(v, _ int) (int, bool) { return v * 2, v%2 == 0 })
			}
		})
	}
}

func BenchmarkFlatMap(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FlatMap(ints, func(v, _ int) []int { return []int{v, v + 1} })
			}
		})
	}
}

func BenchmarkReduce(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Reduce(ints, func(agg, item, _ int) int { return agg + item }, 0)
			}
		})
	}
}

func BenchmarkReduceRight(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.ReduceRight(ints, func(agg, item, _ int) int { return agg + item }, 0)
			}
		})
	}
}

func BenchmarkForEach(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				lo.ForEach(ints, func(_, _ int) {})
			}
		})
	}
}

func BenchmarkForEachWhile(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				lo.ForEachWhile(ints, func(_, _ int) bool { return true })
			}
		})
	}
}

func BenchmarkTimes(b *testing.B) {
	for _, n := range lengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Times(n, func(i int) int { return i })
			}
		})
	}
}

func BenchmarkUniq(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Uniq(ints)
			}
		})
	}
}

func BenchmarkUniqBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.UniqBy(ints, func(v int) int { return v % 100 })
			}
		})
	}
}

func BenchmarkGroupBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.GroupBy(ints, func(v int) int { return v % 10 })
			}
		})
	}
}

func BenchmarkGroupByMap(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.GroupByMap(ints, func(v int) (int, string) { return v % 10, strconv.Itoa(v) })
			}
		})
	}
}

func BenchmarkPartitionBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.PartitionBy(ints, func(v int) int { return v % 5 })
			}
		})
	}
}

func BenchmarkConcat(b *testing.B) {
	for _, n := range lengths {
		a := genSliceInt(n)
		c := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Concat(a, c)
			}
		})
	}
}

func BenchmarkWindow(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Window(ints, 5)
			}
		})
	}
}

func BenchmarkSliding(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Sliding(ints, 5, 2)
			}
		})
	}
}

func BenchmarkInterleave(b *testing.B) {
	for _, n := range lengths {
		a := genSliceInt(n)
		c := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Interleave(a, c)
			}
		})
	}
}

func BenchmarkShuffle(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Shuffle(ints) //nolint:staticcheck
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Reverse(ints) //nolint:staticcheck
			}
		})
	}
}

func BenchmarkRepeatByErr(b *testing.B) {
	for _, n := range lengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.RepeatByErr(n, func(index int) (int, error) { return index * 2, nil })
			}
		})
	}
}

func BenchmarkKeyBy(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.KeyBy(strs, func(v string) string { return v })
			}
		})
	}
}

func BenchmarkAssociate(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Associate(ints, func(v int) (int, string) { return v, strconv.Itoa(v) })
			}
		})
	}
}

func BenchmarkSliceToMap(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.SliceToMap(ints, func(v int) (int, string) { return v, strconv.Itoa(v) })
			}
		})
	}
}

func BenchmarkFilterSliceToMap(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FilterSliceToMap(ints, func(v int) (int, string, bool) { return v, strconv.Itoa(v), v%2 == 0 })
			}
		})
	}
}

func BenchmarkKeyify(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Keyify(ints)
			}
		})
	}
}

func BenchmarkTake(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Take(ints, n/4)
			}
		})
	}
}

func BenchmarkTakeWhile(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.TakeWhile(ints, func(v int) bool { return v < 50000 })
			}
		})
	}
}

func BenchmarkTakeFilter(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.TakeFilter(ints, 5, func(v, _ int) bool { return v%2 == 0 })
			}
		})
	}
}

func BenchmarkFilterReject(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.FilterReject(ints, func(v, _ int) bool { return v%2 == 0 })
			}
		})
	}
}

func BenchmarkCount(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Count(ints, 42)
			}
		})
	}
}

func BenchmarkCountBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.CountBy(ints, func(v int) bool { return v%2 == 0 })
			}
		})
	}
}

func BenchmarkCountValues(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.CountValues(ints)
			}
		})
	}
}

func BenchmarkCountValuesBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.CountValuesBy(ints, func(v int) int { return v % 100 })
			}
		})
	}
}

func BenchmarkSubset(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Subset(ints, n/4, uint(n/2))
			}
		})
	}
}

func BenchmarkSlice(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Slice(ints, n/4, n*3/4)
			}
		})
	}
}

func BenchmarkReplaceAll(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.ReplaceAll(ints, ints[n/4], 123123)
			}
		})
	}
}

func BenchmarkClone(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Clone(ints)
			}
		})
	}
}

func BenchmarkCompact(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Compact(ints)
			}
		})
	}
}

func BenchmarkIsSorted(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.IsSorted(ints)
			}
		})
	}
}

func BenchmarkIsSortedBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.IsSortedBy(ints, func(v int) int { return v })
			}
		})
	}
}

func BenchmarkIsSortedBySorted(b *testing.B) {
	for _, n := range lengths {
		data := genSliceInt(n)
		sort.Ints(data)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				lo.IsSortedBy(data, func(v int) int { return v })
			}
		})
	}
}

func BenchmarkSplice(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		extra := []int{1, 2, 3}
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Splice(ints, n/2, extra...)
			}
		})
	}
}

func BenchmarkCut(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		sep := ints[n/4 : n/4+3]
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = lo.Cut(ints, sep)
			}
		})
	}
}

func BenchmarkCutPrefix(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		prefix := ints[:3]
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.CutPrefix(ints, prefix)
			}
		})
	}
}

func BenchmarkCutSuffix(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		suffix := ints[n-3:]
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.CutSuffix(ints, suffix)
			}
		})
	}
}

func BenchmarkTrim(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		cutset := ints[:3]
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Trim(ints, cutset)
			}
		})
	}
}

func BenchmarkTrimLeft(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		cutset := ints[:3]
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.TrimLeft(ints, cutset)
			}
		})
	}
}

func BenchmarkTrimRight(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		cutset := ints[n-3:]
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.TrimRight(ints, cutset)
			}
		})
	}
}

func BenchmarkTrimPrefix(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		prefix := ints[:3]
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.TrimPrefix(ints, prefix)
			}
		})
	}
}

func BenchmarkTrimSuffix(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		suffix := ints[n-3:]
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.TrimSuffix(ints, suffix)
			}
		})
	}
}

func BenchmarkFilterTakeVsFilterAndTake(b *testing.B) {
	n := 1000
	ints := genSliceInt(n)

	b.Run("lo.TakeFilter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = lo.TakeFilter(ints, 5, func(v, _ int) bool {
				return v%2 == 0
			})
		}
	})

	b.Run("lo.Filter+lo.Take", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = lo.Take(lo.Filter(ints, func(v, _ int) bool { return v%2 == 0 }), 5)
		}
	})

	b.Run("lo.Filter+native_slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			filtered := lo.Filter(ints, func(v, _ int) bool { return v%2 == 0 })
			takeN := 5
			if takeN > len(filtered) {
				_ = filtered
			} else {
				_ = filtered[:takeN]
			}
		}
	})

	b.Run("manual_loop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := make([]int, 0, 5)
			count := 0
			for _, v := range ints {
				if v%2 == 0 {
					result = append(result, v)
					count++
					if count >= 5 {
						break
					}
				}
			}
			_ = result
		}
	})
}

func BenchmarkDifference(b *testing.B) {
	for _, n := range lengths {
		ints1 := genSliceInt(n)
		ints2 := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.Difference(ints1, ints2)
			}
		})
	}

	for _, n := range lengths {
		strs1 := genSliceString(n)
		strs2 := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.Difference(strs1, strs2)
			}
		})
	}
}

func BenchmarkFromSlicePtr(b *testing.B) {
	for _, n := range lengths {
		ptrs := lo.ToSlicePtr(genSliceInt(n))
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FromSlicePtr(ptrs)
			}
		})
	}

	for _, n := range lengths {
		ptrs := lo.ToSlicePtr(genSliceString(n))
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FromSlicePtr(ptrs)
			}
		})
	}
}

func BenchmarkFromSlicePtrOr(b *testing.B) {
	for _, n := range lengths {
		ptrs := lo.ToSlicePtr(genSliceInt(n))
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FromSlicePtrOr(ptrs, -1)
			}
		})
	}

	for _, n := range lengths {
		ptrs := lo.ToSlicePtr(genSliceString(n))
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FromSlicePtrOr(ptrs, "default")
			}
		})
	}
}
