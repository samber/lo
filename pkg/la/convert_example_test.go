package la

import (
	"fmt"
	"github.com/samber/lo"
	"maps"
	"slices"
	"sort"
	"strings"
)

func ExampleTuples() {
	result := Collect(Tuples(maps.All(
		map[string]int{"foo": 1, "bar": 2, "baz": 3},
	)))

	sort.Slice(result, func(i, j int) bool {
		return strings.Compare(result[i].A, result[j].A) < 0
	})

	fmt.Printf("%v", result)
	// Output: [{bar 2} {baz 3} {foo 1}]
}

func ExampleFromTuples() {
	keys, values := KeyValues(FromTuples([]lo.Tuple2[string, int]{
		{"foo", 1}, {"bar", 2}, {"baz", 3},
	}))

	sortedKeys := Collect(keys)
	sortedValues := Collect(values)

	slices.Sort(sortedKeys)
	slices.Sort(sortedValues)

	fmt.Printf("%v %v", sortedKeys, sortedValues)
	// Output: [bar baz foo] [1 2 3]
}

func ExampleCollectMap() {
	mapOut := CollectMap(maps.All(map[string]int{"foo": 1, "bar": 2, "baz": 3}))

	fmt.Printf("%v %v %v", mapOut["foo"], mapOut["bar"], mapOut["baz"])
	// Output: 1 2 3
}

func ExampleCollectToMap() {
	mapOut := map[string]int{"foo": 3, "bar": 1, "baz": 2, "qux": 4}

	mapOut = CollectToMap(
		maps.All(map[string]int{"foo": 1, "bar": 2, "baz": 3}),
		mapOut,
	)

	fmt.Printf("%v %v %v %v", mapOut["foo"], mapOut["bar"], mapOut["baz"], mapOut["qux"])
	// Output: 1 2 3 4
}

func ExampleCollect() {
	sl := Collect(
		slices.Values([]int{1, 2, 3, 4, 5}),
		WithSliceCapacity(100),
	)

	fmt.Printf("%v %v", cap(sl), sl)
	// Output: 100 [1 2 3 4 5]
}

func ExampleCollectTo() {
	sl := []int{9, 10}
	sl = CollectTo(slices.Values([]int{1, 2, 3, 4, 5}), sl)

	fmt.Printf("%v %v", cap(sl), sl)
	// Output: 8 [9 10 1 2 3 4 5]
}

func ExampleSeq2ToSeq() {
	kv := maps.All(map[int]int64{1: 1, 2: 2, 3: 3, 4: 4})

	result := Collect(Seq2ToSeq(kv, func(k int, v int64) string {
		return fmt.Sprintf("%d_%d", k, v)
	}))

	slices.Sort(result)

	fmt.Printf("%v", result)
	// Output: [1_1 2_2 3_3 4_4]
}

func ExampleEntries() {
	result := Collect(Entries(maps.All(map[string]int{"foo": 1, "bar": 2, "baz": 3})))

	sort.Slice(result, func(i, j int) bool {
		return strings.Compare(result[i].Key, result[j].Key) < 0
	})
	fmt.Printf("%v", result)
	// Output: [{bar 2} {baz 3} {foo 1}]
}

func ExampleFromEntries() {
	result := CollectMap(FromEntries(slices.Values([]lo.Entry[string, int]{
		{Key: "foo", Value: 1},
		{Key: "bar", Value: 2},
		{Key: "baz", Value: 3},
	})))

	fmt.Printf("%v %v %v %v", len(result), result["foo"], result["bar"], result["baz"])
	// Output: 3 1 2 3
}

func ExampleKeyBy() {
	list := slices.Values([]string{"a", "aa", "aaa"})

	result := CollectMap(KeyBy(list, func(str string) int {
		return len(str)
	}))

	fmt.Printf("%v", result)
	// Output: map[1:a 2:aa 3:aaa]
}

func ExampleAssociate() {
	list := slices.Values([]string{"a", "aa", "aaa"})

	result := CollectMap(Associate(list, func(str string) (string, int) {
		return str, len(str)
	}))

	fmt.Printf("%v", result)
	// Output: map[a:1 aa:2 aaa:3]
}
