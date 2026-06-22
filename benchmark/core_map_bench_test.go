package benchmark

import (
	"strconv"
	"testing"

	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"github.com/thoas/go-funk"
)

func BenchmarkKeys(b *testing.B) {
	for _, n := range lengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Keys(m)
			}
		})
	}
}

func BenchmarkUniqKeys(b *testing.B) {
	for _, n := range lengths {
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
	for _, n := range lengths {
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
	for _, n := range lengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Values(m)
			}
		})
	}
}

func BenchmarkUniqValues(b *testing.B) {
	m := []map[int64]int64{
		mapGenerator(1000),
		mapGenerator(1000),
		mapGenerator(1000),
	}
	b.Run("lo.UniqValues", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = lo.UniqValues(m...)
		}
	})
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
	for _, n := range lengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.PickBy(m, func(_ string, v int) bool { return v%2 == 0 })
			}
		})
	}
}

func BenchmarkPickByKeys(b *testing.B) {
	for _, n := range lengths {
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
	for _, n := range lengths {
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
	for _, n := range lengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.OmitBy(m, func(_ string, v int) bool { return v%2 == 0 })
			}
		})
	}
}

func BenchmarkOmitByKeys(b *testing.B) {
	for _, n := range lengths {
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
	for _, n := range lengths {
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
	for _, n := range lengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Entries(m)
			}
		})
	}
}

func BenchmarkFromEntries(b *testing.B) {
	for _, n := range lengths {
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
	for _, n := range lengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Invert(m)
			}
		})
	}
}

func BenchmarkAssign(b *testing.B) {
	for _, n := range lengths {
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
	for _, n := range lengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.ChunkEntries(m, 5)
			}
		})
	}
}

func BenchmarkMapKeys(b *testing.B) {
	for _, n := range lengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.MapKeys(m, func(_ int, k string) string { return k + "_x" })
			}
		})
	}
}

func BenchmarkMapValues(b *testing.B) {
	for _, n := range lengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.MapValues(m, func(v int, _ string) int { return v * 2 })
			}
		})
	}
}

func BenchmarkMapEntries(b *testing.B) {
	for _, n := range lengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.MapEntries(m, func(k string, v int) (string, int) { return k, v * 2 })
			}
		})
	}
}

func BenchmarkMapToSlice(b *testing.B) {
	for _, n := range lengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.MapToSlice(m, func(k string, v int) string { return k + "=" + strconv.Itoa(v) })
			}
		})
	}
}

func BenchmarkFilterMapToSlice(b *testing.B) {
	for _, n := range lengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FilterMapToSlice(m, func(k string, v int) (string, bool) { return k, v%2 == 0 })
			}
		})
	}
}

func BenchmarkFilterKeys(b *testing.B) {
	m := mapGenerator(1000)
	b.Run("lo.FilterKeys", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = lo.FilterKeys(m, func(k, v int64) bool { return k%2 == 0 })
		}
	})
}

func BenchmarkFilterValues(b *testing.B) {
	m := mapGenerator(1000)
	b.Run("lo.FilterValues", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = lo.FilterValues(m, func(k, v int64) bool { return v%2 == 0 })
		}
	})
}

// ---------------------------------------------------------------------------
// Comparison benchmarks (lo vs lop vs go-funk vs manual loop)
// ---------------------------------------------------------------------------

func BenchmarkMapComparison(b *testing.B) {
	arr := sliceGenerator(1000000)

	b.Run("lo.Map", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = lo.Map(arr, func(x int64, i int) string {
				return strconv.FormatInt(x, 10)
			})
		}
	})

	b.Run("lop.Map", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = lop.Map(arr, func(x int64, i int) string {
				return strconv.FormatInt(x, 10)
			})
		}
	})

	b.Run("reflect", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = funk.Map(arr, func(x int64) string {
				return strconv.FormatInt(x, 10)
			})
		}
	})

	b.Run("for", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			results := make([]string, len(arr))

			for i, item := range arr {
				result := strconv.FormatInt(item, 10)
				results[i] = result
			}
		}
	})
}
