package benchmark

import (
	"math/rand"
	"strconv"
	"time"
)

var lengths = []int{10, 100, 1000}

func genSliceString(n int) []string {
	res := make([]string, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, strconv.Itoa(rand.Intn(100_000)))
	}
	return res
}

func genSliceInt(n int) []int {
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(100_000))
	}
	return res
}

type heavy = [100]int

func genSliceHeavy(n int) []heavy {
	result := make([]heavy, n)
	for i := range result {
		for j := range result[i] {
			result[i][j] = i + j
		}
	}
	return result
}

func genMap(n int) map[string]int {
	m := make(map[string]int, n)
	for i := 0; i < n; i++ {
		m[strconv.Itoa(i)] = i
	}
	return m
}

type clonableString struct {
	val string
}

func (c clonableString) Clone() clonableString {
	return clonableString{c.val}
}

// sliceGenerator creates a random int64 slice (used by comparison benchmarks).
func sliceGenerator(size uint) []int64 {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	result := make([]int64, size)

	for i := uint(0); i < size; i++ {
		result[i] = r.Int63()
	}

	return result
}

// mapGenerator creates a random int64 map (used by comparison benchmarks).
func mapGenerator(size uint) map[int64]int64 {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	result := make(map[int64]int64, size)

	for i := uint(0); i < size; i++ {
		result[int64(i)] = r.Int63()
	}

	return result
}
