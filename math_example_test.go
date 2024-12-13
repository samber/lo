package lo

import (
	"fmt"
)

func ExampleRange() {
	result1 := Range(4)
	result2 := Range(-4)
	result3 := RangeFrom(1, 5)
	result4 := RangeFrom(1.0, 5)
	result5 := RangeWithSteps(0, 20, 5)
	result6 := RangeWithSteps[float32](-1.0, -4.0, -1.0)
	result7 := RangeWithSteps(1, 4, -1)
	result8 := Range(0)

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	fmt.Printf("%v\n", result6)
	fmt.Printf("%v\n", result7)
	fmt.Printf("%v\n", result8)
	// Output:
	// [0 1 2 3]
	// [0 -1 -2 -3]
	// [1 2 3 4 5]
	// [1 2 3 4 5]
	// [0 5 10 15]
	// [-1 -2 -3]
	// []
	// []
}

func ExampleClamp() {
	result1 := Clamp(0, -10, 10)
	result2 := Clamp(-42, -10, 10)
	result3 := Clamp(42, -10, 10)

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	// Output:
	// 0
	// -10
	// 10
}

func ExampleSum() {
	list := []int{1, 2, 3, 4, 5}

	sum := Sum(list)

	fmt.Printf("%v", sum)
	// Output: 15
}

func ExampleSumBy() {
	list := []string{"foo", "bar"}

	result := SumBy(list, func(item string) int {
		return len(item)
	})

	fmt.Printf("%v", result)
	// Output: 6
}

func ExampleMean() {
	list := []int{1, 2, 3, 4, 5}

	result := Mean(list)

	fmt.Printf("%v", result)
}

func ExampleMeanBy() {
	list := []string{"foo", "bar"}

	result := MeanBy(list, func(item string) int {
		return len(item)
	})

	fmt.Printf("%v", result)
}

func ExampleRound() {
	result1 := Round(1.23456)
	result2 := Round(1.23456, 2)
	result3 := Round(1.23456, 3)
	result4 := Round(1.23456, 7)
	result5 := Round(1.234999999999999, 15)
	result6 := Round(1.234999999999999, 7)
	result7 := Round(1.235, 14)

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	fmt.Printf("%v\n", result6)
	fmt.Printf("%v\n", result7)

	// Output:
	// 1.235
	// 1.23
	// 1.235
	// 1.23456
	// 1.234999999999999
	// 1.235
	// 1.235
}

func ExampleTruncate() {
	result1 := Truncate(1.23456)
	result2 := Truncate(1.23456, 2)
	result3 := Truncate(1.23456, 4)
	result4 := Truncate(1.23456, 7)
	result5 := Truncate(1.2349999999999999, 15)
	result6 := Truncate(1.2349999999999999, 7)

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	fmt.Printf("%v\n", result6)
	// Output:
	// 1.234
	// 1.23
	// 1.2345
	// 1.23456
	// 1.234999999999999
	// 1.2349999
}
