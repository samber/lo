package la

import (
	"fmt"
	"github.com/samber/lo"
	"maps"
	"slices"
	"sort"
)

func ExampleEnumerate() {
	keys, values := KeyValues(Enumerate(slices.Values([]int{1, 2, 3, 4, 5})))

	fmt.Printf("%v %v", Collect(keys), Collect(values))
	// Output: [0 1 2 3 4] [1 2 3 4 5]
}

func ExampleKeys() {
	kv := maps.All(map[string]int{"foo": 1, "bar": 2})
	kv2 := maps.All(map[string]int{"baz": 3})

	result := Collect(Keys(kv, kv2))
	sort.Strings(result)
	fmt.Printf("%v", result)
	// Output: [bar baz foo]
}

func ExampleUniqKeys() {
	kv := maps.All(map[string]int{"foo": 1, "bar": 2})
	kv2 := maps.All(map[string]int{"bar": 3})

	result := Collect(UniqKeys(kv, kv2))
	sort.Strings(result)
	fmt.Printf("%v", result)
	// Output: [bar foo]
}

func ExampleValues() {
	kv := maps.All(map[string]int{"foo": 1, "bar": 2})
	kv2 := maps.All(map[string]int{"baz": 3})

	result := Collect(Values(kv, kv2))

	sort.Ints(result)
	fmt.Printf("%v", result)
	// Output: [1 2 3]
}

func ExampleUniqValues() {
	kv := maps.All(map[string]int{"foo": 1, "bar": 2})
	kv2 := maps.All(map[string]int{"baz": 2})

	result := Collect(UniqValues(kv, kv2))

	sort.Ints(result)
	fmt.Printf("%v", result)
	// Output: [1 2]
}

func ExampleInvert() {
	kv := maps.All(map[string]int{"foo": 1, "bar": 2, "baz": 3})

	result := CollectMap(Invert(kv))

	fmt.Printf("%v %v %v %v", len(result), result[1], result[2], result[3])
	// Output: 3 foo bar baz
}

func ExampleJoin2() {
	it := Join2(
		maps.All(map[string]int{"a": 1, "b": 2}),
		maps.All(map[string]int{"b": 3, "c": 4}),
	)

	resultMap := CollectMap(it)
	resultTuples := slices.Collect(Tuples(it))

	fmt.Printf("%v %v %v %v %v", len(resultMap), resultMap["a"], resultMap["b"], resultMap["c"], resultTuples)
	// Output: 3 1 3 4 [{a 1} {b 2} {b 3} {c 4}]
}

func ExampleFollow() {
	v1, v2 := Enumerate(slices.Values([]int{1, 2, 3})), Enumerate(slices.Values([]int{6, 7, 8}))

	res := Follow(v1, v2)

	ks, vs := lo.Unzip2(Collect(Tuples(res)))

	fmt.Printf("%v %v", ks, vs)
	// Output: [0 1 2 3 4 5] [1 2 3 6 7 8]
}
