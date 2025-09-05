package lo

import (
	"fmt"
	"math"
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

func ExampleUniqMap() {
	type User struct {
		Name string
		Age  int
	}
	users := []User{{Name: "Alex", Age: 10}, {Name: "Alex", Age: 12}, {Name: "Bob", Age: 11}, {Name: "Alice", Age: 20}}

	result := UniqMap(users, func(u User, index int) string {
		return u.Name
	})

	fmt.Printf("%v", result)
	// Output: [Alex Bob Alice]
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

func ExampleForEachWhile() {
	list := []int64{1, 2, -math.MaxInt, 4}

	ForEachWhile(list, func(x int64, _ int) bool {
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

	fmt.Printf("%v\n", result[0])
	fmt.Printf("%v\n", result[1])
	fmt.Printf("%v\n", result[2])
	// Output:
	// [0 3]
	// [1 4]
	// [2 5]
}

func ExampleGroupByMap() {
	list := []int{0, 1, 2, 3, 4, 5}

	result := GroupByMap(list, func(i int) (int, int) {
		return i % 3, i * 2
	})

	fmt.Printf("%v\n", result[0])
	fmt.Printf("%v\n", result[1])
	fmt.Printf("%v\n", result[2])
	// Output:
	// [0 6]
	// [2 8]
	// [4 10]
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

func ExampleInterleave() {
	list1 := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	list2 := [][]int{{1}, {2, 5, 8}, {3, 6}, {4, 7, 9, 10}}

	result1 := Interleave(list1...)
	result2 := Interleave(list2...)

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	// Output:
	// [1 2 3 4 5 6 7 8 9]
	// [1 2 3 4 5 6 7 8 9 10]
}

func ExampleShuffle() {
	list := []int{0, 1, 2, 3, 4, 5}

	result := Shuffle(list)

	fmt.Printf("%v", result)
}

func ExampleReverse() {
	list := []int{0, 1, 2, 3, 4, 5}

	result := Reverse(list)

	fmt.Printf("%v", result)
	// Output: [5 4 3 2 1 0]
}

func ExampleFill() {
	list := []foo{{"a"}, {"a"}}

	result := Fill(list, foo{"b"})

	fmt.Printf("%v", result)
	// Output: [{b} {b}]
}

func ExampleRepeat() {
	result := Repeat(2, foo{"a"})

	fmt.Printf("%v", result)
	// Output: [{a} {a}]
}

func ExampleRepeatBy() {
	result := RepeatBy(5, func(i int) string {
		return strconv.FormatInt(int64(math.Pow(float64(i), 2)), 10)
	})

	fmt.Printf("%v", result)
	// Output: [0 1 4 9 16]
}

func ExampleKeyBy() {
	list := []string{"a", "aa", "aaa"}

	result := KeyBy(list, func(str string) int {
		return len(str)
	})

	fmt.Printf("%v", result)
	// Output: map[1:a 2:aa 3:aaa]
}

func ExampleSliceToMap() {
	list := []string{"a", "aa", "aaa"}

	result := SliceToMap(list, func(str string) (string, int) {
		return str, len(str)
	})

	fmt.Printf("%v", result)
	// Output: map[a:1 aa:2 aaa:3]
}

func ExampleFilterSliceToMap() {
	list := []string{"a", "aa", "aaa"}

	result := FilterSliceToMap(list, func(str string) (string, int, bool) {
		return str, len(str), len(str) > 1
	})

	fmt.Printf("%v", result)
	// Output: map[aa:2 aaa:3]
}

func ExampleKeyify() {
	list := []string{"a", "a", "b", "b", "d"}

	set := Keyify(list)
	_, ok1 := set["a"]
	_, ok2 := set["c"]
	fmt.Printf("%v\n", ok1)
	fmt.Printf("%v\n", ok2)
	fmt.Printf("%v\n", set)

	// Output:
	// true
	// false
	// map[a:{} b:{} d:{}]
}

func ExampleDrop() {
	list := []int{0, 1, 2, 3, 4, 5}

	result := Drop(list, 2)

	fmt.Printf("%v", result)
	// Output: [2 3 4 5]
}

func ExampleDropRight() {
	list := []int{0, 1, 2, 3, 4, 5}

	result := DropRight(list, 2)

	fmt.Printf("%v", result)
	// Output: [0 1 2 3]
}

func ExampleDropWhile() {
	list := []int{0, 1, 2, 3, 4, 5}

	result := DropWhile(list, func(val int) bool {
		return val < 2
	})

	fmt.Printf("%v", result)
	// Output: [2 3 4 5]
}

func ExampleDropRightWhile() {
	list := []int{0, 1, 2, 3, 4, 5}

	result := DropRightWhile(list, func(val int) bool {
		return val > 2
	})

	fmt.Printf("%v", result)
	// Output: [0 1 2]
}

func ExampleDropByIndex() {
	list := []int{0, 1, 2, 3, 4, 5}

	result := DropByIndex(list, 2)

	fmt.Printf("%v", result)
	// Output: [0 1 3 4 5]
}

func ExampleReject() {
	list := []int{0, 1, 2, 3, 4, 5}

	result := Reject(list, func(x int, _ int) bool {
		return x%2 == 0
	})

	fmt.Printf("%v", result)
	// Output: [1 3 5]
}

func ExampleCount() {
	list := []int{0, 1, 2, 3, 4, 5, 0, 1, 2, 3}

	result := Count(list, 2)

	fmt.Printf("%v", result)
	// Output: 2
}

func ExampleCountBy() {
	list := []int{0, 1, 2, 3, 4, 5, 0, 1, 2, 3}

	result := CountBy(list, func(i int) bool {
		return i < 4
	})

	fmt.Printf("%v", result)
	// Output: 8
}

func ExampleCountValues() {
	result1 := CountValues([]int{})
	result2 := CountValues([]int{1, 2})
	result3 := CountValues([]int{1, 2, 2})
	result4 := CountValues([]string{"foo", "bar", ""})
	result5 := CountValues([]string{"foo", "bar", "bar"})

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	// Output:
	// map[]
	// map[1:1 2:1]
	// map[1:1 2:2]
	// map[:1 bar:1 foo:1]
	// map[bar:2 foo:1]
}

func ExampleCountValuesBy() {
	isEven := func(v int) bool {
		return v%2 == 0
	}

	result1 := CountValuesBy([]int{}, isEven)
	result2 := CountValuesBy([]int{1, 2}, isEven)
	result3 := CountValuesBy([]int{1, 2, 2}, isEven)

	length := func(v string) int {
		return len(v)
	}

	result4 := CountValuesBy([]string{"foo", "bar", ""}, length)
	result5 := CountValuesBy([]string{"foo", "bar", "bar"}, length)

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	// Output:
	// map[]
	// map[false:1 true:1]
	// map[false:1 true:2]
	// map[0:1 3:2]
	// map[3:3]
}

func ExampleSubset() {
	list := []int{0, 1, 2, 3, 4, 5}

	result := Subset(list, 2, 3)

	fmt.Printf("%v", result)
	// Output: [2 3 4]
}

func ExampleSlice() {
	list := []int{0, 1, 2, 3, 4, 5}

	result := Slice(list, 1, 4)
	fmt.Printf("%v\n", result)

	result = Slice(list, 4, 1)
	fmt.Printf("%v\n", result)

	result = Slice(list, 4, 5)
	fmt.Printf("%v\n", result)

	// Output:
	// [1 2 3]
	// []
	// [4]
}

func ExampleReplace() {
	list := []int{0, 1, 0, 1, 2, 3, 0}

	result := Replace(list, 0, 42, 1)
	fmt.Printf("%v\n", result)

	result = Replace(list, -1, 42, 1)
	fmt.Printf("%v\n", result)

	result = Replace(list, 0, 42, 2)
	fmt.Printf("%v\n", result)

	result = Replace(list, 0, 42, -1)
	fmt.Printf("%v\n", result)

	// Output:
	// [42 1 0 1 2 3 0]
	// [0 1 0 1 2 3 0]
	// [42 1 42 1 2 3 0]
	// [42 1 42 1 2 3 42]
}

func ExampleReplaceAll() {
	list := []string{"", "foo", "", "bar", ""}

	result := Compact(list)

	fmt.Printf("%v", result)

	// Output: [foo bar]
}

func ExampleClone() {
	list := []int{1, 2, 3, 4, 5}

	result := Clone(list)

	fmt.Printf("%v", result)
	// Output: [1 2 3 4 5]
}

func ExampleIsSorted() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := IsSorted(list)

	fmt.Printf("%v", result)

	// Output: true
}

func ExampleIsSortedByKey() {
	list := []string{"a", "bb", "ccc"}

	result := IsSortedByKey(list, func(s string) int {
		return len(s)
	})

	fmt.Printf("%v", result)

	// Output: true
}
