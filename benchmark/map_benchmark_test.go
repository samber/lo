package benchmark

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"github.com/thoas/go-funk"
)

func sliceGenerator(size uint) []int64 {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	result := make([]int64, size)

	for i := uint(0); i < size; i++ {
		result[i] = r.Int63()
	}

	return result
}

func mapGenerator(size uint) map[int64]int64 {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	result := make(map[int64]int64, size)

	for i := uint(0); i < size; i++ {
		result[int64(i)] = r.Int63()
	}

	return result
}

func BenchmarkMap(b *testing.B) {
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

func BenchmarkUniqKeys(b *testing.B) {
	m := []map[int64]int64{
		mapGenerator(1000),
		mapGenerator(1000),
		mapGenerator(1000),
	}
	b.Run("lo.UniqKeys", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = lo.UniqKeys(m...)
		}
	})
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
