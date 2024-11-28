package la

import (
	"fmt"
	"math"
	"strconv"
)

func ExampleTimes() {
	result := Collect(Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	}))

	fmt.Printf("%v", result)
	// Output: [0 1 2]
}

func ExampleTimes2() {
	keys, values := KeyValues(Times2(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	}))

	fmt.Printf("%v %v", Collect(keys), Collect(values))
	// Output: [0 1 2] [0 1 2]
}

func ExampleRepeat() {
	result := Collect(Repeat(2, foo{"a"}))

	fmt.Printf("%v", result)
	// Output: [{a} {a}]
}

func ExampleRepeatBy() {
	result := Collect(RepeatBy(5, func(i int) string {
		return strconv.FormatInt(int64(math.Pow(float64(i), 2)), 10)
	}))

	fmt.Printf("%v", result)
	// Output: [0 1 4 9 16]
}
