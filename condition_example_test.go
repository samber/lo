package lo_test

import (
	"fmt"

	"github.com/samber/lo"
)

func ExampleTernary() {
	result := lo.Ternary(true, "a", "b")

	fmt.Printf("%v", result)
	// Output: a
}

func ExampleTernaryF() {
	result := lo.TernaryF(true, func() string { return "a" }, func() string { return "b" })

	fmt.Printf("%v", result)
	// Output: a
}

func ExampleIf() {
	result1 := lo.If(true, 1).
		ElseIf(false, 2).
		Else(3)

	result2 := lo.If(false, 1).
		ElseIf(true, 2).
		Else(3)

	result3 := lo.If(false, 1).
		ElseIf(false, 2).
		Else(3)

	result4 := lo.IfF(true, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result5 := lo.IfF(false, func() int { return 1 }).
		ElseIfF(true, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result6 := lo.IfF(false, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	fmt.Printf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func ExampleIfF() {
	result1 := lo.If(true, 1).
		ElseIf(false, 2).
		Else(3)

	result2 := lo.If(false, 1).
		ElseIf(true, 2).
		Else(3)

	result3 := lo.If(false, 1).
		ElseIf(false, 2).
		Else(3)

	result4 := lo.IfF(true, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result5 := lo.IfF(false, func() int { return 1 }).
		ElseIfF(true, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result6 := lo.IfF(false, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	fmt.Printf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func ExampleifElse_ElseIf() {
	result1 := lo.If(true, 1).
		ElseIf(false, 2).
		Else(3)

	result2 := lo.If(false, 1).
		ElseIf(true, 2).
		Else(3)

	result3 := lo.If(false, 1).
		ElseIf(false, 2).
		Else(3)

	result4 := lo.IfF(true, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result5 := lo.IfF(false, func() int { return 1 }).
		ElseIfF(true, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result6 := lo.IfF(false, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	fmt.Printf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func ExampleifElse_ElseIfF() {
	result1 := lo.If(true, 1).
		ElseIf(false, 2).
		Else(3)

	result2 := lo.If(false, 1).
		ElseIf(true, 2).
		Else(3)

	result3 := lo.If(false, 1).
		ElseIf(false, 2).
		Else(3)

	result4 := lo.IfF(true, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result5 := lo.IfF(false, func() int { return 1 }).
		ElseIfF(true, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result6 := lo.IfF(false, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	fmt.Printf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func ExampleifElse_Else() {
	result1 := lo.If(true, 1).
		ElseIf(false, 2).
		Else(3)

	result2 := lo.If(false, 1).
		ElseIf(true, 2).
		Else(3)

	result3 := lo.If(false, 1).
		ElseIf(false, 2).
		Else(3)

	result4 := lo.IfF(true, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result5 := lo.IfF(false, func() int { return 1 }).
		ElseIfF(true, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result6 := lo.IfF(false, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	fmt.Printf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func ExampleifElse_ElseF() {
	result1 := lo.If(true, 1).
		ElseIf(false, 2).
		Else(3)

	result2 := lo.If(false, 1).
		ElseIf(true, 2).
		Else(3)

	result3 := lo.If(false, 1).
		ElseIf(false, 2).
		Else(3)

	result4 := lo.IfF(true, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result5 := lo.IfF(false, func() int { return 1 }).
		ElseIfF(true, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result6 := lo.IfF(false, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	fmt.Printf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func ExampleSwitch() {
	result1 := lo.Switch[int, string](1).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result2 := lo.Switch[int, string](2).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result3 := lo.Switch[int, string](42).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result4 := lo.Switch[int, string](1).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result5 := lo.Switch[int, string](2).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result6 := lo.Switch[int, string](42).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	fmt.Printf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func ExampleswitchCase_Case() {
	result1 := lo.Switch[int, string](1).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result2 := lo.Switch[int, string](2).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result3 := lo.Switch[int, string](42).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result4 := lo.Switch[int, string](1).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result5 := lo.Switch[int, string](2).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result6 := lo.Switch[int, string](42).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	fmt.Printf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func ExampleswitchCase_CaseF() {
	result1 := lo.Switch[int, string](1).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result2 := lo.Switch[int, string](2).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result3 := lo.Switch[int, string](42).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result4 := lo.Switch[int, string](1).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result5 := lo.Switch[int, string](2).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result6 := lo.Switch[int, string](42).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	fmt.Printf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func ExampleswitchCase_Default() {
	result1 := lo.Switch[int, string](1).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result2 := lo.Switch[int, string](2).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result3 := lo.Switch[int, string](42).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result4 := lo.Switch[int, string](1).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result5 := lo.Switch[int, string](2).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result6 := lo.Switch[int, string](42).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	fmt.Printf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func ExampleswitchCase_DefaultF() {
	result1 := lo.Switch[int, string](1).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result2 := lo.Switch[int, string](2).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result3 := lo.Switch[int, string](42).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result4 := lo.Switch[int, string](1).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result5 := lo.Switch[int, string](2).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result6 := lo.Switch[int, string](42).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
	fmt.Printf("%v\n", result4)
	fmt.Printf("%v\n", result5)
	fmt.Printf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}
