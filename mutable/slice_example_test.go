package mutable

import "fmt"

func ExampleMap() {
	list := []int{1, 2, 3, 4}

	Map(list, func(nbr int) int {
		return nbr * 2
	})

	fmt.Printf("%v", list)
	// Output: [2 4 6 8]
}

func ExampleMapI() {
	list := []int{1, 2, 3, 4}

	MapI(list, func(nbr int, index int) int {
		return nbr * index
	})

	fmt.Printf("%v", list)
	// Output: [0 2 6 12]
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
