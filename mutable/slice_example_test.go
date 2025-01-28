package mutable

import "fmt"

func ExampleShuffle() {
	list := []int{0, 1, 2, 3, 4, 5}

	Shuffle(list)

	fmt.Printf("%v", list)
}

func ExampleReverse() {
	list := []int{0, 1, 2, 3, 4, 5}

	Reverse(list)

	fmt.Printf("%v", list)
	// Output: [5 4 3 2 1 0]
}

// Fill fills elements of array with `initial` value.
// Play: https://go.dev/play/p/VwR34GzqEub
func Fill[T any, Slice ~[]T](collection Slice, initial T) {
	for i := range collection {
		collection[i] = initial
	}
}
