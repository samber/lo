package la

import (
	"fmt"
	"github.com/samber/lo"
	"math"
	"slices"
)

func ExampleReduce() {
	result := Reduce(slices.Values([]int64{1, 2, 3, 4}), func(agg int64, item int64) int64 {
		return agg + item
	}, 0)

	fmt.Printf("%v", result)
	// Output: 10
}

func ExampleReduce2() {
	result := Reduce2(Enumerate(slices.Values([]int64{1, 2, 3, 4})), func(agg int64, k int, item int64) int64 {
		if k%2 == 0 {
			return agg + item
		}

		return agg
	}, 0)

	fmt.Printf("%v", result)
	// Output: 4
}

func ExampleForEach() {
	ForEach(slices.Values([]int64{1, 2, 3, 4}), func(x int64) {
		fmt.Println(x)
	})

	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleForEach2() {
	ForEach2(Enumerate(slices.Values([]int64{1, 2, 3, 4})), func(k int, x int64) {
		fmt.Printf("%d %d\n", k, x)
	})

	// Output:
	// 0 1
	// 1 2
	// 2 3
	// 3 4
}

func ExampleForEachWhile() {
	ForEachWhile(slices.Values([]int64{1, 2, -math.MaxInt, 4}), func(x int64) bool {
		if x < 0 {
			return false
		}
		fmt.Println(x)
		return true
	})

	// Output:
	// 1
	// 2
}

func ExampleForEachWhile2() {
	ForEachWhile2(Enumerate(slices.Values([]int64{1, 2, -math.MaxInt, 4})), func(k int, x int64) bool {
		if x < 0 {
			return false
		}
		fmt.Printf("%d %d\n", k, x)
		return true
	})

	// Output:
	// 0 1
	// 1 2
}

func ExampleKeyValues() {
	kv := FromTuples([]lo.Tuple2[string, int]{
		{"foo", 1},
		{"bar", 2},
		{"baz", 3},
	})

	keys, values := KeyValues(kv)

	fmt.Printf("%v %v", Collect(keys), Collect(values))
	// Output: [foo bar baz] [1 2 3]
}
