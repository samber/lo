package lo

import (
	"fmt"

	"github.com/kr/pretty"
)

func ExampleKeys() {
	result := Keys(map[string]int{"foo": 1})
	fmt.Printf("%v", result)
	// Output: [foo]
}

func ExampleValues() {
	result := Values(map[string]int{"foo": 1})
	fmt.Printf("%v", result)
	// Output: [1]
}

func ExamplePickBy() {
	result := PickBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(key string, value int) bool {
		return value%2 == 0
	})
	fmt.Printf("%v", result)
	// Output: map[bar:2]
}

func ExamplePickByKeys() {
	result := PickByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"bar"})
	fmt.Printf("%v", result)
	// map[bar: 2}
}

func ExamplePickByValues() {
	result := PickByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{2})
	fmt.Printf("%v", result)
	// map[bar: 2}
}

func ExampleOmitBy() {
	result := OmitBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(key string, value int) bool {
		return value%2 == 1
	})
	fmt.Printf("%v", result)
	// Output: map[bar:2]
}

func ExampleOmitByKeys() {
	result := OmitByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz"})
	fmt.Printf("%v", result)
	// Output: map[bar:2]
}

func ExampleOmitByValues() {
	result := OmitByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{1, 3})
	fmt.Printf("%v", result)
	// Output: map[bar:2]
}

func ExampleEntries() {
	result := Entries(map[string]int{"foo": 1})
	pretty.Println(result)
	// Output:
	// []lo.Entry[string,int]{
	//     {Key:"foo", Value:1},
	//}
}

func ExampleFromEntries() {
	result := FromEntries([]Entry[string, int]{
		{
			Key:   "foo",
			Value: 1,
		},
	})
	fmt.Printf("%v", result)
	// Output: map[foo:1]
}

func ExampleInvert() {
	result := Invert(map[string]int{"a": 1})
	fmt.Printf("%v", result)
	// Output: map[1:a]
}

func ExampleAssign() {
	result := Assign(map[string]int{"a": 1}, map[string]int{"a": 3})
	fmt.Printf("%v", result)
	// Output: map[a:3]
}

func ExampleMapKeys() {
	var numberToWord = map[int]string{
		1: "one",
	}

	result := MapKeys(map[int]string{1: "foo"}, func(_ string, k int) string {
		return numberToWord[k]
	})
	fmt.Printf("%v", result)
	// Output: map[one:foo]
}

func ExampleMapValues() {
	result := MapValues[int, string, int](
		map[int]string{1: "racecar"},
		func(v string, _ int) int {
			return len(v)
		},
	)
	fmt.Printf("%v", result)
	// Output: map[1:6]
}
