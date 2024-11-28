package la

import (
	"fmt"
	"iter"
	"maps"
	"slices"
	"strconv"
)

func ExampleMap() {
	result := Collect(Map(slices.Values([]int64{1, 2, 3, 4}), func(nbr int64) string {
		return strconv.FormatInt(nbr*2, 10)
	}))

	fmt.Printf("%v", result)
	// Output: [2 4 6 8]
}

func ExampleMap2() {
	result := CollectMap(Map2(maps.All(map[string]int{"foo": 1, "bar": 2}), func(k string, v int) (int, string) {
		return v, k
	}))

	fmt.Printf("%v\n", result)
	// Output: map[1:foo 2:bar]
}

func ExampleMapKeys() {
	result := CollectMap(MapKeys(maps.All(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}), func(_ int, k int) string {
		return strconv.FormatInt(int64(k), 10)
	}))

	fmt.Printf("%v %v %v %v %v", len(result), result["1"], result["2"], result["3"], result["4"])
	// Output: 4 1 2 3 4
}

func ExampleMapValues() {
	result := CollectMap(MapValues(maps.All(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}), func(v int, _ int) string {
		return strconv.FormatInt(int64(v), 10)
	}))

	fmt.Printf("%v %q %q %q %q", len(result), result[1], result[2], result[3], result[4])
	// Output: 4 "1" "2" "3" "4"
}

func ExampleFilterMap() {
	result := Collect(FilterMap(slices.Values([]int64{1, 2, 3, 4}), func(nbr int64) (string, bool) {
		return strconv.FormatInt(nbr*2, 10), nbr%2 == 0
	}))

	fmt.Printf("%v", result)
	// Output: [4 8]
}

func ExampleFilterMap2() {
	result := CollectMap(FilterMap2(maps.All(map[string]int64{"foo": 1, "bar": 2, "baz": 3}), func(k string, nbr int64) (string, bool) {
		return strconv.FormatInt(nbr*2, 10), nbr%2 == 0
	}))

	fmt.Printf("%v", result)
	// Output: map[bar:4]
}

func ExampleFlatMap() {
	result := Collect(FlatMap(slices.Values([]int64{1, 2, 3, 4}), func(nbr int64) iter.Seq[string] {
		return slices.Values([]string{
			strconv.FormatInt(nbr, 10), // base 10
			strconv.FormatInt(nbr, 2),  // base 2
		})
	}))

	fmt.Printf("%v", result)
	// Output: [1 1 2 10 3 11 4 100]
}

func ExampleFlatMapEnumerated() {
	keys, values := KeyValues(FlatMapEnumerated(Enumerate(slices.Values([]int64{1, 2, 3, 4})), func(_ int, nbr int64) iter.Seq[string] {
		return slices.Values([]string{
			strconv.FormatInt(nbr, 10), // base 10
			strconv.FormatInt(nbr, 2),  // base 2
		})
	}))

	fmt.Printf("%v-%v", Collect(keys), Collect(values))
	// Output: [0 1 2 3 4 5 6 7]-[1 1 2 10 3 11 4 100]
}

func ExampleChunk() {
	list := slices.Values([]int{0, 1, 2, 3, 4})

	result := Chunk(list, 2)

	for item := range result {
		fmt.Printf("%v\n", item)
	}
	// Output:
	// [0 1]
	// [2 3]
	// [4]
}

func ExampleChunk2() {
	list := Enumerate(slices.Values([]int{4, 3, 2, 3, 4}))

	result := Chunk2(list, 2)

	for item := range result {
		keys, values := KeyValues(item)

		fmt.Printf("%v %v\n", Collect(keys), Collect(values))
	}
	// Output:
	// [0 1] [4 3]
	// [2 3] [2 3]
	// [4] [4]
}

func ExampleFlatten() {
	list := slices.Values([]iter.Seq[int]{slices.Values([]int{0, 1, 2}), slices.Values([]int{3, 4, 5})})

	result := Flatten(list)

	fmt.Printf("%v", Collect(result))
	// Output: [0 1 2 3 4 5]
}
func ExampleFlattenSlice() {
	list := slices.Values([][]int{{0, 1, 2}, {3, 4, 5}})

	result := FlattenSlice(list)

	fmt.Printf("%v", Collect(result))
	// Output: [0 1 2 3 4 5]
}

func ExampleInterleave() {
	list1 := []iter.Seq[int]{
		slices.Values([]int{1, 4, 7}),
		slices.Values([]int{2, 5, 8}),
		slices.Values([]int{3, 6, 9}),
	}
	list2 := []iter.Seq[int]{
		slices.Values([]int{1}),
		slices.Values([]int{2, 5, 8}),
		slices.Values([]int{3, 6}),
		slices.Values([]int{4, 7, 9, 10}),
	}

	result1 := Collect(Interleave(list1...))
	result2 := Collect(Interleave(list2...))

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	// Output:
	// [1 2 3 4 5 6 7 8 9]
	// [1 2 3 4 5 6 7 8 9 10]
}

func ExampleInterleave2() {
	list1 := []iter.Seq2[int, int]{
		Enumerate(slices.Values([]int{1, 4, 7})),
		Enumerate(slices.Values([]int{2, 5, 8})),
		Enumerate(slices.Values([]int{3, 6, 9})),
	}
	list2 := []iter.Seq2[int, int]{
		Enumerate(slices.Values([]int{1})),
		Enumerate(slices.Values([]int{2, 5, 8})),
		Enumerate(slices.Values([]int{3, 6})),
		Enumerate(slices.Values([]int{4, 7, 9, 10})),
	}

	keys1, values1 := KeyValues(Interleave2(list1...))
	keys2, values2 := KeyValues(Interleave2(list2...))

	fmt.Printf("%v %v\n", Collect(keys1), Collect(values1))
	fmt.Printf("%v %v\n", Collect(keys2), Collect(values2))
	// Output:
	// [0 0 0 1 1 1 2 2 2] [1 2 3 4 5 6 7 8 9]
	// [0 0 0 0 1 1 1 2 2 3] [1 2 3 4 5 6 7 8 9 10]
}
