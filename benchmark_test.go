package lo

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/thoas/go-funk"

	lop "github.com/samber/lo/parallel"
)

func sliceGenerator(size uint) []int64 {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	result := make([]int64, size)

	for i := uint(0); i < size; i++ {
		result[i] = r.Int63()
	}

	return result
}

func BenchmarkMap(b *testing.B) {
	arr := sliceGenerator(1000000)

	b.Run("lo.Map", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = Map(arr, func(x int64, i int) string {
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

func BenchmarkForEach(b *testing.B) {
	arr := Times(1000, func(index int) time.Location {
		return time.Location{}
	})

	b.Run("ForEach", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			ForEach(arr, func(item time.Location, index int) { _ = item })
		}
	})

	b.Run("oldForEach", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			iteratee := func(item time.Location, index int) { _ = item }
			for i, item := range arr {
				iteratee(item, i)
			}
		}
	})
}
