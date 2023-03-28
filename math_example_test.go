package lo_test

import (
	"fmt"

	"github.com/samber/lo"
)

func ExampleRange() {
	result1 := lo.Range(4)
	result2 := lo.Range(-4)
	result3 := lo.RangeFrom(1, 5)
	result4 := lo.RangeFrom(1.0, 5)
	result5 := lo.RangeWithSteps(0, 20, 5)
	result6 := lo.RangeWithSteps[float32](-1.0, -4.0, -1.0)
	result7 := lo.RangeWithSteps(1, 4, -1)
	result8 := lo.Range(0)

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
	result1 := lo.Clamp(0, -10, 10)
	result2 := lo.Clamp(-42, -10, 10)
	result3 := lo.Clamp(42, -10, 10)

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

	sum := lo.Sum(list)

	fmt.Printf("%v", sum)
	// Output: 15
}

func ExampleSumBy() {
	list := []string{"foo", "bar"}

	result := lo.SumBy(list, func(item string) int {
		return len(item)
	})

	fmt.Printf("%v", result)
	// Output: 6
}
