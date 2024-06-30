//go:build goexperiment.rangefunc

package lo

import (
	"fmt"
	"strconv"
)

func ExampleFilterI() {
	iter := ToIterator(1, 2, 3, 4)

	result := FilterI(iter, func(nbr int, index int) bool {
		return nbr%2 == 0
	})

	fmt.Printf("%v", result.Slice())
	// Output: [2 4]
}

func ExampleMapI() {
	iter := ToIterator(1, 2, 3, 4)

	result := MapI(iter, func(nbr int, index int) string {
		return strconv.FormatInt(int64(nbr)*2, 10)
	})

	fmt.Printf("%v", result.Slice())
	// Output: [2 4 6 8]
}

func ExampleFilterMapI() {
	iter := ToIterator(1, 2, 3, 4)

	result := FilterMapI(iter, func(nbr int, index int) (string, bool) {
		return strconv.FormatInt(int64(nbr)*2, 10), nbr%2 == 0
	})

	fmt.Printf("%v", result.Slice())
	// Output: [4 8]
}
