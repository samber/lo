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

func ExampleProduct() {
	list := []int{1, 2, 3, 4, 5}

	result := Product(list)

	fmt.Printf("%v", result)
	// Output: 120
}

func ExampleProductBy() {
	list := []string{"foo", "bar"}

	result := ProductBy(list, func(item string) int {
		return len(item)
	})

	fmt.Printf("%v", result)
	// Output: 9
}

func ExampleMean() {
	list := []int{1, 2, 3, 4, 5}

	result := Mean(list)

	fmt.Printf("%v", result)
	// Output: 3
}

func ExampleMeanBy() {
	list := []string{"foo", "bar"}

	result := MeanBy(list, func(item string) int {
		return len(item)
	})

	fmt.Printf("%v", result)
	// Output: 3
}
