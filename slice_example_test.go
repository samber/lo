package lo

import (
	"fmt"
	"strconv"
)

func ExampleFilter() {
	list := []int64{1, 2, 3, 4}

	result := Filter(list, func(nbr int64, index int) bool {
		return nbr%2 == 0
	})

	fmt.Printf("%v", result)
	// Output: [2 4]
}

func ExampleMap() {
	list := []int64{1, 2, 3, 4}

	result := Map(list, func(nbr int64, index int) string {
		return strconv.FormatInt(nbr*2, 10)
	})

	fmt.Printf("%v", result)
	// Output: [2 4 6 8]
}

func ExampleFilterMap() {
	list := []int64{1, 2, 3, 4}

	result := FilterMap(list, func(nbr int64, index int) (string, bool) {
		return strconv.FormatInt(nbr*2, 10), nbr%2 == 0
	})

	fmt.Printf("%v", result)
	// Output: [4 8]
}

func ExampleFlatMap() {
	list := []int64{1, 2, 3, 4}

	result := FlatMap(list, func(nbr int64, index int) []string {
		return []string{
			strconv.FormatInt(nbr, 10), // base 10
			strconv.FormatInt(nbr, 2),  // base 2
		}
	})

	fmt.Printf("%v", result)
	// Output: [1 1 2 10 3 11 4 100]
}

func ExampleReduce() {
	list := []int64{1, 2, 3, 4}

	result := Reduce(list, func(agg int64, item int64, index int) int64 {
		return agg + item
	}, 0)

	fmt.Printf("%v", result)
	// Output: 10
}

func ExampleReduceRight() {
	list := [][]int{{0, 1}, {2, 3}, {4, 5}}

	result := ReduceRight(list, func(agg []int, item []int, index int) []int {
		return append(agg, item...)
	}, []int{})

	fmt.Printf("%v", result)
	// Output: [4 5 2 3 0 1]
}

func ExampleForEach() {
	list := []int64{1, 2, 3, 4}

	ForEach(list, func(x int64, _ int) {
		fmt.Println(x)
	})

	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleTimes() {
	result := Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})

	fmt.Printf("%v", result)
	// Output: [0 1 2]
}

func ExampleUniq() {
	list := []int{1, 2, 2, 1}

	result := Uniq(list)

	fmt.Printf("%v", result)
	// Output: [1 2]
}

func ExampleUniqBy() {
	list := []int{0, 1, 2, 3, 4, 5}

	result := UniqBy(list, func(i int) int {
		return i % 3
	})

	fmt.Printf("%v", result)
	// Output: [0 1 2]
}

func ExampleGroupBy() {
	list := []int{0, 1, 2, 3, 4, 5}

	result := GroupBy(list, func(i int) int {
		return i % 3
	})

	for _, item := range result {
		fmt.Printf("%v\n", item)
	}
	// Output:
	// [0 3]
	// [1 4]
	// [2 5]
}

func ExampleChunk() {
	list := []int{0, 1, 2, 3, 4}

	result := Chunk(list, 2)

	for _, item := range result {
		fmt.Printf("%v\n", item)
	}
	// Output:
	// [0 1]
	// [2 3]
	// [4]
}

func ExamplePartitionBy() {
	list := []int{-2, -1, 0, 1, 2, 3, 4}

	result := PartitionBy(list, func(x int) string {
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	})

	for _, item := range result {
		fmt.Printf("%v\n", item)
	}
	// Output:
	// [-2 -1]
	// [0 2 4]
	// [1 3]
}

func ExampleFlatten() {
	list := [][]int{{0, 1, 2}, {3, 4, 5}}

	result := Flatten(list)

	fmt.Printf("%v", result)
	// Output: [0 1 2 3 4 5]
}
