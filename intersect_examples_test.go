package lo

import (
	"fmt"
)

func ExampleContains() {
	ok := Contains([]int{0, 1, 2, 3, 4, 5}, 4)
	fmt.Printf("%t\n", ok)
	// Output: true
}

func ExampleContainsBy() {
	ok := ContainsBy([]int{0, 1, 2, 3, 4, 5}, func(x int) bool {
		return x == 4
	})
	fmt.Printf("%t\n", ok)
	// Output: true
}

func ExampleEvery() {
	ok := Every([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	fmt.Printf("%t\n", ok)
	// Output: true
}

func ExampleEvery_notAllFound() {
	ok := Every([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	fmt.Printf("%t\n", ok)
	// Output: false
}

func ExampleSome() {
	ok := Some([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	fmt.Printf("%t\n", ok)
	// Output: true
}

func ExampleSome_noIntersect() {
	ok := Some([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	fmt.Printf("%t\n", ok)
	// Output: false
}

func ExampleIntersect() {
	result := Intersect([]int{0, 1, 2, 3, 4, 5}, []int{2, 4, 6})
	fmt.Printf("%v\n", result)
	// Output: [2 4]
}

func ExampleDifference() {
	left, right := Difference([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 6})
	fmt.Printf("%v\n", left)
	fmt.Printf("%v\n", right)
	// Output:
	// [1 3 4 5]
	// [6]
}

func ExampleDifference_identical() {
	left, right := Difference([]int{0, 1, 2}, []int{0, 1, 2})
	fmt.Printf("%v\n", left)
	fmt.Printf("%v\n", right)
	// Output:
	// []
	// []
}

func ExampleUnion() {
	result := Union([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 10})
	fmt.Printf("%v\n", result)
	// Output [0 1 2 3 4 5 10]
}
