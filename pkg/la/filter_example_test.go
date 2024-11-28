package la

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
)

func ExampleFilter() {
	result := slices.Collect(Filter(slices.Values([]int64{1, 2, 3, 4}), func(nbr int64) bool {
		return nbr%2 == 0
	}))

	fmt.Printf("%v", result)
	// Output: [2 4]
}

func ExampleFilter2() {
	result := CollectMap(Filter2(maps.All(map[string]int{"foo": 1, "bar": 2, "baz": 3}), func(key string, value int) bool {
		return value%2 == 1
	}))

	fmt.Printf("%v %v %v", len(result), result["foo"], result["baz"])
	// Output: 2 1 3
}

func ExampleFilterByKeys() {
	result := CollectMap(FilterByKeys(maps.All(map[string]int{"foo": 1, "bar": 2, "baz": 3}), []string{"foo", "baz"}))

	fmt.Printf("%v %v %v", len(result), result["foo"], result["baz"])
	// Output: 2 1 3
}

func ExampleFilterByValues() {
	result := CollectMap(FilterByValues(maps.All(map[string]int{"foo": 1, "bar": 2, "baz": 3}), []int{1, 3}))

	fmt.Printf("%v %v %v", len(result), result["foo"], result["baz"])
	// Output: 2 1 3
}

func ExampleReject() {
	list := slices.Values([]int{0, 1, 2, 3, 4, 5})

	result := Collect(Reject(list, func(x int) bool {
		return x%2 == 0
	}))

	fmt.Printf("%v", result)
	// Output: [1 3 5]
}

func ExampleReject2() {
	kv := maps.All(map[string]int{"foo": 1, "bar": 2, "baz": 3})

	result := CollectMap(Reject2(kv, func(key string, value int) bool {
		return value%2 == 1
	}))

	fmt.Printf("%v", result)
	// Output: map[bar:2]
}

func ExampleRejectByKeys() {
	kv := maps.All(map[string]int{"foo": 1, "bar": 2, "baz": 3})

	result := CollectMap(RejectByKeys(kv, []string{"foo", "baz"}))

	fmt.Printf("%v", result)
	// Output: map[bar:2]
}

func ExampleRejectByValues() {
	kv := maps.All(map[string]int{"foo": 1, "bar": 2, "baz": 3})

	result := CollectMap(RejectByValues(kv, []int{1, 3}))

	fmt.Printf("%v", result)
	// Output: map[bar:2]
}

func ExampleRejectMap() {
	list := slices.Values([]int64{1, 2, 3, 4})

	result := Collect(RejectMap(list, func(nbr int64) (string, bool) {
		return strconv.FormatInt(nbr*2, 10), nbr%2 != 0
	}))

	fmt.Printf("%v", result)
	// Output: [4 8]
}

func ExampleUniq() {
	list := slices.Values([]int{1, 2, 2, 1})

	result := Collect(Uniq(list))

	fmt.Printf("%v", result)
	// Output: [1 2]
}

func ExampleUniqBy() {
	list := slices.Values([]int{0, 1, 2, 3, 4, 5})

	result := Collect(UniqBy(list, func(i int) int {
		return i % 3
	}))

	fmt.Printf("%v", result)
	// Output: [0 1 2]
}

func ExampleReplace() {
	list := slices.Values([]int{0, 1, 0, 1, 2, 3, 0})

	result := Collect(Replace(list, 0, 42, 1))
	fmt.Printf("%v\n", result)

	result = Collect(Replace(list, -1, 42, 1))
	fmt.Printf("%v\n", result)

	result = Collect(Replace(list, 0, 42, 2))
	fmt.Printf("%v\n", result)

	result = Collect(Replace(list, 0, 42, -1))
	fmt.Printf("%v\n", result)

	// Output:
	// [42 1 0 1 2 3 0]
	// [0 1 0 1 2 3 0]
	// [42 1 42 1 2 3 0]
	// [42 1 42 1 2 3 42]
}

func ExampleReplace2() {
	list := Enumerate(slices.Values([]int{0, 1, 0, 1, 2, 3, 0}))

	keys, values := KeyValues(Replace2(list, 0, 42, 1))
	fmt.Printf("%v %v\n", Collect(keys), Collect(values))

	keys, values = KeyValues(Replace2(list, -1, 42, 1))
	fmt.Printf("%v %v\n", Collect(keys), Collect(values))

	keys, values = KeyValues(Replace2(list, 0, 42, 2))
	fmt.Printf("%v %v\n", Collect(keys), Collect(values))

	keys, values = KeyValues(Replace2(list, 0, 42, -1))
	fmt.Printf("%v %v\n", Collect(keys), Collect(values))

	// Output:
	// [0 1 2 3 4 5 6] [42 1 0 1 2 3 0]
	// [0 1 2 3 4 5 6] [0 1 0 1 2 3 0]
	// [0 1 2 3 4 5 6] [42 1 42 1 2 3 0]
	// [0 1 2 3 4 5 6] [42 1 42 1 2 3 42]
}

func ExampleCompact() {
	list := slices.Values([]string{"", "foo", "", "bar", ""})

	result := Collect(Compact(list))

	fmt.Printf("%v", result)

	// Output: [foo bar]
}

func ExampleCompact2() {
	list := Enumerate(slices.Values([]string{"", "foo", "", "bar", ""}))

	keys, values := KeyValues(Compact2(list))

	fmt.Printf("%v %v", Collect(keys), Collect(values))

	// Output: [1 3] [foo bar]
}
