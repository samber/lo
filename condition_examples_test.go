package lo

import (
	"fmt"
)

func ExampleTernary_true() {
	result := Ternary[string](true, "a", "b")
	fmt.Println(result)
	// Output: a
}

func ExampleTernary_false() {
	result := Ternary[string](false, "a", "b")
	fmt.Println(result)
	// Output: b
}

func ExampleIf() {
	result := If[int](true, 1).
		ElseIf(false, 2).
		Else(3)
	fmt.Println(result)
	// Output: 1
}

func ExampleIf_elseIf() {
	result := If[int](false, 1).
		ElseIf(true, 2).
		Else(3)
	fmt.Println(result)
	// Output: 2
}

func ExampleIf_elseIf_else() {
	result := If[int](false, 1).
		ElseIf(false, 2).
		Else(3)
	fmt.Println(result)
	// Output: 3
}

func ExampleIf_callbacks() {
	result := IfF[int](true, func() int {
		return 1
	}).
		ElseIfF(false, func() int {
			return 2
		}).
		ElseF(func() int {
			return 3
		})
	fmt.Println(result)
	// Output: 1
}

func ExampleIf_mixed() {
	result := IfF[int](false, func() int {
		return 1
	}).
		Else(42)
	fmt.Println(result)
	// Output: 42
}

func ExampleSwitch_firstCase() {
	result := Switch[int, string](1).
		Case(1, "first case").
		Case(2, "second").
		Default("nothing else matched")
	fmt.Println(result)
	// Output: first case
}

func ExampleSwitch_second() {
	result := Switch[int, string](2).
		Case(1, "first case").
		Case(2, "second").
		Default("nothing else matched")
	fmt.Println(result)
	// Output: second
}

func ExampleSwitch_default() {
	result := Switch[int, string](42).
		Case(1, "first case").
		Case(2, "second").
		Default("nothing else matched")
	fmt.Println(result)
	// Output: nothing else matched
}

func ExampleSwitch_callbacks() {
	result := Switch[int, string](2).
		CaseF(1, func() string {
			return "case 1"
		}).
		CaseF(2, func() string {
			return "case 2"
		}).
		DefaultF(func() string {
			return "case 3"
		})
	fmt.Println(result)
	// Output: case 2
}

func ExampleSwitch_mixed() {
	result := Switch[int, string](9999).
		CaseF(1, func() string {
			return "1"
		}).
		Default("42")
	fmt.Println(result)
	// Output: 42
}
