package benchmark

import (
	"strconv"
	"testing"

	"github.com/samber/lo"
)

var coreLengths = []int{10, 100, 1000}

func genMap(n int) map[string]int {
	m := make(map[string]int, n)
	for i := 0; i < n; i++ {
		m[strconv.Itoa(i)] = i
	}
	return m
}

// ---------------------------------------------------------------------------
// map.go
// ---------------------------------------------------------------------------

func BenchmarkKeys(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Keys(m)
			}
		})
	}
}

func BenchmarkCoreUniqKeys(b *testing.B) {
	for _, n := range coreLengths {
		m1 := genMap(n)
		m2 := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.UniqKeys(m1, m2)
			}
		})
	}
}

func BenchmarkHasKey(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		key := strconv.Itoa(n / 2)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.HasKey(m, key)
			}
		})
	}
}

func BenchmarkValues(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Values(m)
			}
		})
	}
}

func BenchmarkValueOr(b *testing.B) {
	m := genMap(100)
	b.Run("hit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = lo.ValueOr(m, "50", -1)
		}
	})
	b.Run("miss", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = lo.ValueOr(m, "missing", -1)
		}
	})
}

func BenchmarkPickBy(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.PickBy(m, func(_ string, v int) bool { return v%2 == 0 })
			}
		})
	}
}

func BenchmarkPickByKeys(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		keys := make([]string, n/2)
		for i := range keys {
			keys[i] = strconv.Itoa(i * 2)
		}
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.PickByKeys(m, keys)
			}
		})
	}
}

func BenchmarkPickByValues(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		vals := make([]int, n/2)
		for i := range vals {
			vals[i] = i * 2
		}
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.PickByValues(m, vals)
			}
		})
	}
}

func BenchmarkOmitBy(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.OmitBy(m, func(_ string, v int) bool { return v%2 == 0 })
			}
		})
	}
}

func BenchmarkOmitByKeys(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		keys := make([]string, n/4)
		for i := range keys {
			keys[i] = strconv.Itoa(i)
		}
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.OmitByKeys(m, keys)
			}
		})
	}
}

func BenchmarkOmitByValues(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		vals := make([]int, n/4)
		for i := range vals {
			vals[i] = i
		}
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.OmitByValues(m, vals)
			}
		})
	}
}

func BenchmarkEntries(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Entries(m)
			}
		})
	}
}

func BenchmarkFromEntries(b *testing.B) {
	for _, n := range coreLengths {
		entries := make([]lo.Entry[string, int], n)
		for i := 0; i < n; i++ {
			entries[i] = lo.Entry[string, int]{Key: strconv.Itoa(i), Value: i}
		}
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FromEntries(entries)
			}
		})
	}
}

func BenchmarkInvert(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Invert(m)
			}
		})
	}
}

func BenchmarkAssign(b *testing.B) {
	for _, n := range coreLengths {
		m1 := genMap(n)
		m2 := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Assign(m1, m2)
			}
		})
	}
}

func BenchmarkChunkEntries(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.ChunkEntries(m, 5)
			}
		})
	}
}

func BenchmarkMapKeys(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.MapKeys(m, func(_ int, k string) string { return k + "_x" })
			}
		})
	}
}

func BenchmarkMapValues(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.MapValues(m, func(v int, _ string) int { return v * 2 })
			}
		})
	}
}

func BenchmarkMapEntries(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.MapEntries(m, func(k string, v int) (string, int) { return k, v * 2 })
			}
		})
	}
}

func BenchmarkMapToSlice(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.MapToSlice(m, func(k string, v int) string { return k + "=" + strconv.Itoa(v) })
			}
		})
	}
}

func BenchmarkFilterMapToSlice(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FilterMapToSlice(m, func(k string, v int) (string, bool) { return k, v%2 == 0 })
			}
		})
	}
}

// ---------------------------------------------------------------------------
// find.go
// ---------------------------------------------------------------------------

func BenchmarkIndexOf(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		target := ints[n-1] // worst case: last element
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.IndexOf(ints, target)
			}
		})
	}
}

func BenchmarkLastIndexOf(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		target := ints[0]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.LastIndexOf(ints, target)
			}
		})
	}
}

func BenchmarkHasPrefix(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		prefix := ints[:n/10+1]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.HasPrefix(ints, prefix)
			}
		})
	}
}

func BenchmarkHasSuffix(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		suffix := ints[n-n/10-1:]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.HasSuffix(ints, suffix)
			}
		})
	}
}

func BenchmarkFind(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		target := ints[n-1]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.Find(ints, func(v int) bool { return v == target })
			}
		})
	}
}

func BenchmarkFindIndexOf(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		target := ints[n-1]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = lo.FindIndexOf(ints, func(v int) bool { return v == target })
			}
		})
	}
}

func BenchmarkFindLastIndexOf(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		target := ints[0]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = lo.FindLastIndexOf(ints, func(v int) bool { return v == target })
			}
		})
	}
}

func BenchmarkFindOrElse(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FindOrElse(ints, -1, func(v int) bool { return v == -999 })
			}
		})
	}
}

func BenchmarkFindKey(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.FindKey(m, n/2)
			}
		})
	}
}

func BenchmarkFindKeyBy(b *testing.B) {
	for _, n := range coreLengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.FindKeyBy(m, func(_ string, v int) bool { return v == n/2 })
			}
		})
	}
}

func BenchmarkFindUniques(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FindUniques(ints)
			}
		})
	}
}

func BenchmarkFindUniquesBy(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FindUniquesBy(ints, func(v int) int { return v % 50 })
			}
		})
	}
}

func BenchmarkFindDuplicates(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FindDuplicates(ints)
			}
		})
	}
}

func BenchmarkFindDuplicatesBy(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FindDuplicatesBy(ints, func(v int) int { return v % 50 })
			}
		})
	}
}

func BenchmarkMin(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Min(ints)
			}
		})
	}
}

func BenchmarkMinIndex(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.MinIndex(ints)
			}
		})
	}
}

func BenchmarkMinBy(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.MinBy(ints, func(a, b int) bool { return a < b })
			}
		})
	}
}

func BenchmarkMinIndexBy(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.MinIndexBy(ints, func(a, b int) bool { return a < b })
			}
		})
	}
}

func BenchmarkMax(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Max(ints)
			}
		})
	}
}

func BenchmarkMaxIndex(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.MaxIndex(ints)
			}
		})
	}
}

func BenchmarkMaxBy(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.MaxBy(ints, func(a, b int) bool { return a > b })
			}
		})
	}
}

func BenchmarkMaxIndexBy(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.MaxIndexBy(ints, func(a, b int) bool { return a > b })
			}
		})
	}
}

func BenchmarkFirst(b *testing.B) {
	ints := genSliceInt(100)
	for i := 0; i < b.N; i++ {
		_, _ = lo.First(ints)
	}
}

func BenchmarkFirstOrEmpty(b *testing.B) {
	ints := genSliceInt(100)
	for i := 0; i < b.N; i++ {
		_ = lo.FirstOrEmpty(ints)
	}
}

func BenchmarkLast(b *testing.B) {
	ints := genSliceInt(100)
	for i := 0; i < b.N; i++ {
		_, _ = lo.Last(ints)
	}
}

func BenchmarkLastOrEmpty(b *testing.B) {
	ints := genSliceInt(100)
	for i := 0; i < b.N; i++ {
		_ = lo.LastOrEmpty(ints)
	}
}

func BenchmarkNth(b *testing.B) {
	ints := genSliceInt(100)
	for i := 0; i < b.N; i++ {
		_, _ = lo.Nth(ints, 50)
	}
}

func BenchmarkSample(b *testing.B) {
	ints := genSliceInt(100)
	for i := 0; i < b.N; i++ {
		_ = lo.Sample(ints)
	}
}

func BenchmarkSamples(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Samples(ints, n/4)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// intersect.go
// ---------------------------------------------------------------------------

func BenchmarkContains(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		target := ints[n-1]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Contains(ints, target)
			}
		})
	}
}

func BenchmarkContainsBy(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		target := ints[n-1]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.ContainsBy(ints, func(v int) bool { return v == target })
			}
		})
	}
}

func BenchmarkEvery(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		subset := ints[:n/2]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Every(ints, subset)
			}
		})
	}
}

func BenchmarkEveryBy(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.EveryBy(ints, func(v int) bool { return v >= 0 })
			}
		})
	}
}

func BenchmarkSome(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		subset := []int{ints[n-1]}
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Some(ints, subset)
			}
		})
	}
}

func BenchmarkSomeBy(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.SomeBy(ints, func(v int) bool { return v < 0 })
			}
		})
	}
}

func BenchmarkNone(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		subset := []int{-1, -2, -3}
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.None(ints, subset)
			}
		})
	}
}

func BenchmarkNoneBy(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.NoneBy(ints, func(v int) bool { return v < 0 })
			}
		})
	}
}

func BenchmarkIntersect(b *testing.B) {
	for _, n := range coreLengths {
		a := genSliceInt(n)
		c := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Intersect(a, c)
			}
		})
	}
}

func BenchmarkIntersectBy(b *testing.B) {
	for _, n := range coreLengths {
		a := genSliceInt(n)
		c := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.IntersectBy(func(v int) int { return v }, a, c)
			}
		})
	}
}

func BenchmarkUnion(b *testing.B) {
	for _, n := range coreLengths {
		a := genSliceInt(n)
		c := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Union(a, c)
			}
		})
	}
}

func BenchmarkWithout(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Without(ints, 1, 2, 3, 4, 5)
			}
		})
	}
}

func BenchmarkWithoutBy(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.WithoutBy(ints, func(v int) int { return v % 100 }, 1, 2, 3, 4, 5)
			}
		})
	}
}

func BenchmarkWithoutEmpty(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		// sprinkle some zeroes
		for j := 0; j < n/10; j++ {
			ints[j*10] = 0
		}
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.WithoutEmpty(ints) //nolint:staticcheck
			}
		})
	}
}

func BenchmarkWithoutNth(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.WithoutNth(ints, 0, n/2, n-1)
			}
		})
	}
}

func BenchmarkElementsMatch(b *testing.B) {
	for _, n := range coreLengths {
		a := genSliceInt(n)
		c := make([]int, n)
		copy(c, a)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.ElementsMatch(a, c)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// math.go
// ---------------------------------------------------------------------------

func BenchmarkRange(b *testing.B) {
	for _, n := range coreLengths {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Range(n)
			}
		})
	}
}

func BenchmarkRangeFrom(b *testing.B) {
	for _, n := range coreLengths {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.RangeFrom(0, n)
			}
		})
	}
}

func BenchmarkRangeWithSteps(b *testing.B) {
	for _, n := range coreLengths {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.RangeWithSteps(0, n, 1)
			}
		})
	}
}

func BenchmarkClamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.Clamp(15, 0, 10)
	}
}

func BenchmarkSum(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Sum(ints)
			}
		})
	}
}

func BenchmarkSumBy(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.SumBy(ints, func(v int) int { return v })
			}
		})
	}
}

func BenchmarkProduct(b *testing.B) {
	for _, n := range coreLengths {
		floats := make([]float64, n)
		for j := range floats {
			floats[j] = 1.0001
		}
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Product(floats)
			}
		})
	}
}

func BenchmarkProductBy(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.ProductBy(ints, func(v int) float64 { return float64(v) * 0.001 })
			}
		})
	}
}

func BenchmarkMean(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Mean(ints)
			}
		})
	}
}

func BenchmarkMeanBy(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.MeanBy(ints, func(v int) int { return v })
			}
		})
	}
}

func BenchmarkMode(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Mode(ints)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// string.go
// ---------------------------------------------------------------------------

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.RandomString(64, lo.AlphanumericCharset)
	}
}

func BenchmarkSubstring(b *testing.B) {
	s := lo.RandomString(1000, lo.LettersCharset)
	for i := 0; i < b.N; i++ {
		_ = lo.Substring(s, 100, 200)
	}
}

func BenchmarkChunkString(b *testing.B) {
	s := lo.RandomString(1000, lo.LettersCharset)
	for i := 0; i < b.N; i++ {
		_ = lo.ChunkString(s, 10)
	}
}

func BenchmarkRuneLength(b *testing.B) {
	s := lo.RandomString(1000, lo.LettersCharset)
	for i := 0; i < b.N; i++ {
		_ = lo.RuneLength(s)
	}
}

func BenchmarkPascalCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.PascalCase("some_long_variable_name")
	}
}

func BenchmarkCamelCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.CamelCase("some_long_variable_name")
	}
}

func BenchmarkKebabCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.KebabCase("someLongVariableName")
	}
}

func BenchmarkSnakeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.SnakeCase("someLongVariableName")
	}
}

func BenchmarkWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.Words("someLongVariableName")
	}
}

func BenchmarkCapitalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.Capitalize("hello world")
	}
}

func BenchmarkEllipsis(b *testing.B) {
	s := lo.RandomString(200, lo.LettersCharset)
	for i := 0; i < b.N; i++ {
		_ = lo.Ellipsis(s, 50)
	}
}

// ---------------------------------------------------------------------------
// type_manipulation.go
// ---------------------------------------------------------------------------

func BenchmarkToPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.ToPtr(42)
	}
}

func BenchmarkFromPtr(b *testing.B) {
	p := lo.ToPtr(42)
	for i := 0; i < b.N; i++ {
		_ = lo.FromPtr(p)
	}
}

func BenchmarkFromPtrOr(b *testing.B) {
	var p *int
	for i := 0; i < b.N; i++ {
		_ = lo.FromPtrOr(p, 99)
	}
}

func BenchmarkCoreToSlicePtr(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.ToSlicePtr(ints)
			}
		})
	}
}

func BenchmarkToAnySlice(b *testing.B) {
	for _, n := range coreLengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.ToAnySlice(ints)
			}
		})
	}
}

func BenchmarkFromAnySlice(b *testing.B) {
	for _, n := range coreLengths {
		anys := lo.ToAnySlice(genSliceInt(n))
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.FromAnySlice[int](anys)
			}
		})
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.IsEmpty(0)
	}
}

func BenchmarkIsNotEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.IsNotEmpty(42)
	}
}

func BenchmarkCoalesce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = lo.Coalesce(0, 0, 0, 42, 99)
	}
}
