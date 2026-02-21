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

// @TODO: also apply to UniqValues.
func BenchmarkUniqKeys(b *testing.B) {
	m := []map[int64]int64{
		mapGenerator(100000),
		mapGenerator(100000),
		mapGenerator(100000),
	}

	// allocate just in time + ordered
	b.Run("lo.UniqKeys.jit-alloc", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			seen := make(map[int64]struct{})
			var result []int64

			for i := range m {
				for k := range m[i] {
					if _, exists := seen[k]; exists {
						continue
					}
					seen[k] = struct{}{}
					result = append(result, k) //nolint:staticcheck
				}
			}
		}
	})

	// preallocate + unordered
	b.Run("lo.UniqKeys.preallocate", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			size := 0
			for i := range m {
				size += len(m[i])
			}
			seen := make(map[int64]struct{}, size)

			for i := range m {
				for k := range m[i] {
					seen[k] = struct{}{}
				}
			}

			result := make([]int64, 0, len(seen))

			for k := range seen {
				result = append(result, k) //nolint:staticcheck
			}
		}
	})
}
