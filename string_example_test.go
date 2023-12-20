package lo_test

import (
	"fmt"
	"math"

	"github.com/samber/lo"
)

func ExampleSubstring() {
	result1 := lo.Substring("hello", 2, 3)
	result2 := lo.Substring("hello", -4, 3)
	result3 := lo.Substring("hello", -2, math.MaxUint)
	result4 := lo.Substring("🏠🐶🐱", 0, 2)
	result5 := lo.Substring("你好，世界", 0, 3)

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	// Output:
	// llo
	// ell
	// lo
	// 🏠🐶
	// 你好，
}

func ExampleChunkString() {
	result1 := lo.ChunkString("123456", 2)
	result2 := lo.ChunkString("1234567", 2)
	result3 := lo.ChunkString("", 2)
	result4 := lo.ChunkString("1", 2)

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	// Output:
	// [12 34 56]
	// [12 34 56 7]
	// []
	// [1]
}

func ExampleRuneLength() {
	result1, chars1 := lo.RuneLength("hellô"), len("hellô")
	result2, chars2 := lo.RuneLength("🤘"), len("🤘")

	fmt.Printf("%v %v\n", result1, chars1)
	fmt.Printf("%v %v\n", result2, chars2)
	// Output:
	// 5 6
	// 1 4
}
