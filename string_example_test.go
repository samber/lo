package lo

import (
	"fmt"
	"math"
)

func ExampleSubstring() {
	result1 := Substring("hello", 2, 3)
	result2 := Substring("hello", -4, 3)
	result3 := Substring("hello", -2, math.MaxUint)

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	// Output:
	// llo
	// ell
	// lo
}

func ExampleChunkString() {
	result1 := ChunkString("123456", 2)
	result2 := ChunkString("1234567", 2)
	result3 := ChunkString("", 2)
	result4 := ChunkString("1", 2)

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
	result1, chars1 := RuneLength("hellô"), len("hellô")
	result2, chars2 := RuneLength("🤘"), len("🤘")

	fmt.Printf("%v %v\n", result1, chars1)
	fmt.Printf("%v %v\n", result2, chars2)
	// Output:
	// 5 6
	// 1 4
}
