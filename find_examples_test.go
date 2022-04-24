package lo

import (
	"fmt"
	"math/rand"
)

func ExampleIndexOf_found() {
	result := IndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	fmt.Println(result)
	// Output: 2
}

func ExampleIndexOf_notFound() {
	result := IndexOf([]int{0, 1, 2, 1, 2, 3}, 6)
	fmt.Println(result)
	// Output: -1
}

func ExampleLastIndexOf_found() {
	result := LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	fmt.Println(result)
	// Output: 4
}

func ExampleLastIndexOf_notFound() {
	result := LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 6)
	fmt.Println(result)
	// Output: -1
}

func ExampleFind_found() {
	str, ok := Find([]string{"a", "b", "c", "d"}, func(i string) bool {
		return i == "b"
	})
	fmt.Printf("%q, %t\n", str, ok)
	// Output: "b", true
}

func ExampleFind_notFound() {
	str, ok := Find([]string{"a", "b", "c", "d"}, func(i string) bool {
		return i == "z"
	})
	fmt.Printf("%q, %t\n", str, ok)
	// Output: "", false
}

func ExampleFindIndexOf_found() {
	str, index, ok := FindIndexOf([]string{"a", "b", "a", "b"}, func(i string) bool {
		return i == "b"
	})
	fmt.Printf("%q, %d, %t\n", str, index, ok)
	// Output: "b", 1, true
}

func ExampleFindIndexOf_notFound() {
	str, index, ok := FindIndexOf([]string{"a", "b", "c", "d"}, func(i string) bool {
		return i == "z"
	})
	fmt.Printf("%q, %d, %t\n", str, index, ok)
	// Output: "", -1, false
}

func ExampleFindLastIndexOf_found() {
	str, index, ok := FindLastIndexOf[string]([]string{"a", "b", "a", "b"}, func(i string) bool {
		return i == "b"
	})
	fmt.Printf("%q, %d, %t\n", str, index, ok)
	// Output: "b", 3, true
}

func ExampleFindLastIndexOf_notFound() {
	str, index, ok := FindLastIndexOf[string]([]string{"a", "b", "c", "d"}, func(i string) bool {
		return i == "z"
	})
	fmt.Printf("%q, %d, %t\n", str, index, ok)
	// Output: "", -1, false
}

func ExampleFindOrElse_found() {
	result := FindOrElse([]string{"a", "b", "c", "d"}, "fallback", func(i string) bool {
		return i == "c"
	})
	fmt.Println(result)
	// Output: c
}

func ExampleFindOrElse_notFound() {
	result := FindOrElse([]string{"a", "b", "c", "d"}, "fallback", func(i string) bool {
		return i == "z"
	})
	fmt.Println(result)
	// Output: fallback
}

func ExampleMin() {
	result := Min([]int{1, 2, 3})
	fmt.Println(result)
	// Output: 1
}

// This example demonstrates that the type's "default value" is returned for empty collections
func ExampleMin_emptyCollection_int() {
	result := Min([]int{})
	fmt.Println(result)
	// Output: 0
}

// This example demonstrates that the type's "default value" is returned for empty collections
func ExampleMin_emptyCollection_string() {
	result := Min([]string{})
	fmt.Printf("%q\n", result)
	// Output: ""
}

func ExampleMinBy() {
	result := MinBy([]string{"string1", "str2", "st3", "st4", "str5"}, func(left string, right string) bool {
		return len(left) > 3 && len(left) < len(right)
	})
	fmt.Printf("%q\n", result)
	// Output: "str2"
}

func ExampleMinBy_emptyCollection() {
	result := MinBy([]string{}, func(left string, right string) bool {
		return len(left) < len(right)
	})
	fmt.Printf("%q\n", result)
	// Output: ""
}

func ExampleMax() {
	result := Max([]int{1, 5, 3})
	fmt.Println(result)
	// Output: 5
}

// This example demonstrates that the type's "default value" is returned for empty collections
func ExampleMax_emptyCollection_int() {
	result := Max([]int{})
	fmt.Println(result)
	// Output: 0
}

// This example demonstrates that the type's "default value" is returned for empty collections
func ExampleMax_emptyCollection_string() {
	result := Max([]string{})
	fmt.Printf("%q\n", result)
	// Output: ""
}

func ExampleMaxBy() {
	result := MaxBy([]string{"st1", "str2", "st3", "string4", "str5"}, func(item string, max string) bool {
		return len(item) > len(max)
	})
	fmt.Printf("%q\n", result)
	// Output: "string4"
}

func ExampleMaxBy_emptyCollection() {
	result := MaxBy([]string{}, func(item string, max string) bool {
		return len(item) > len(max)
	})
	fmt.Printf("%q\n", result)
	// Output: ""
}

func ExampleLast() {
	result, err := Last([]int{1, 2, 3})
	fmt.Printf("%d, %v", result, err)
	// Output: 3, <nil>
}

func ExampleLast_empty() {
	result, err := Last([]int{})
	fmt.Printf("%d, %q", result, err)
	// Output: 0, "last: cannot extract the last element of an empty slice"
}

func ExampleNth() {
	result, err := Nth([]string{"a", "b", "c", "d", "e"}, 2)
	fmt.Printf("%q, %v", result, err)
	// Output: "c", <nil>
}

func ExampleNth_negative() {
	result, err := Nth([]string{"a", "b", "c", "d", "e"}, -2)
	fmt.Printf("%q, %v", result, err)
	// Output: "d", <nil>
}

func ExampleNth_outOfBounds() {
	result, err := Nth([]string{"a", "b", "c", "d"}, 10)
	fmt.Printf("%q, %q", result, err)
	// Output: "", "nth: 10 out of slice bounds"
}

// This example demonstrates getting a sample from a collection, using your own random source
func ExampleSourceSample() {
	src := rand.NewSource(1)
	result := SourceSample(src, []string{"a", "b", "c", "d", "e"})
	fmt.Printf("%q", result)
	// Output: "b"
}

// This example demonstrates getting a sample from a collection, using your own randomizer
func ExampleIntnSample() {
	r := rand.New(rand.NewSource(1))
	result := IntnSample(r.Intn, []string{"a", "b", "c", "d", "e"})
	fmt.Printf("%q", result)
	// Output: "b"
}

// This example demonstrates getting a sample from a collection, using your own randomizer
func ExampleIntnSample_notVeryRandom() {
	alwaysReturn2 := func(in int) int {
		return 2
	}
	result := IntnSample(alwaysReturn2, []string{"a", "b", "c", "d", "e"})
	fmt.Printf("%q", result)
	// Output: "c"
}

// This example demonstrates getting a sample from a collection, implicitly using package rand
func ExampleSample() {
	rand.Seed(1)
	result := Sample([]string{"a", "b", "c", "d", "e"})
	fmt.Printf("%q", result)
	// Output: "b"
}

// This example demonstrates getting a series of samples from a collection, implicitly using package rand
func ExampleSamples() {
	rand.Seed(1)
	result := Samples([]string{"a", "b", "c"}, 2)
	fmt.Printf("%v", result)
	// Output: [c b]
}
