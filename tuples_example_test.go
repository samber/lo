package lo_test

import (
	"fmt"

	"github.com/samber/lo"
)

func ExampleT2() {
	result := lo.T2("hello", 2)
	fmt.Printf("%v %v", result.A, result.B)
	// Output: hello 2
}

func ExampleT3() {
	result := lo.T3("hello", 2, true)
	fmt.Printf("%v %v %v", result.A, result.B, result.C)
	// Output: hello 2 true
}

func ExampleT4() {
	result := lo.T4("hello", 2, true, foo{bar: "bar"})
	fmt.Printf("%v %v %v %v", result.A, result.B, result.C, result.D)
	// Output: hello 2 true {bar}
}

func ExampleT5() {
	result := lo.T5("hello", 2, true, foo{bar: "bar"}, 4.2)
	fmt.Printf("%v %v %v %v %v", result.A, result.B, result.C, result.D, result.E)
	// Output: hello 2 true {bar} 4.2
}

func ExampleT6() {
	result := lo.T6("hello", 2, true, foo{bar: "bar"}, 4.2, "plop")
	fmt.Printf("%v %v %v %v %v %v", result.A, result.B, result.C, result.D, result.E, result.F)
	// Output: hello 2 true {bar} 4.2 plop
}

func ExampleT7() {
	result := lo.T7("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false)
	fmt.Printf("%v %v %v %v %v %v %v", result.A, result.B, result.C, result.D, result.E, result.F, result.G)
	// Output: hello 2 true {bar} 4.2 plop false
}

func ExampleT8() {
	result := lo.T8("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false, 42)
	fmt.Printf("%v %v %v %v %v %v %v %v", result.A, result.B, result.C, result.D, result.E, result.F, result.G, result.H)
	// Output: hello 2 true {bar} 4.2 plop false 42
}

func ExampleT9() {
	result := lo.T9("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false, 42, "hello world")
	fmt.Printf("%v %v %v %v %v %v %v %v %v", result.A, result.B, result.C, result.D, result.E, result.F, result.G, result.H, result.I)
	// Output: hello 2 true {bar} 4.2 plop false 42 hello world
}

func ExampleUnpack2() {
	a, b := lo.Unpack2(lo.T2("hello", 2))
	fmt.Printf("%v %v", a, b)
	// Output: hello 2
}

func ExampleUnpack3() {
	a, b, c := lo.Unpack3(lo.T3("hello", 2, true))
	fmt.Printf("%v %v %v", a, b, c)
	// Output: hello 2 true
}

func ExampleUnpack4() {
	a, b, c, d := lo.Unpack4(lo.T4("hello", 2, true, foo{bar: "bar"}))
	fmt.Printf("%v %v %v %v", a, b, c, d)
	// Output: hello 2 true {bar}
}

func ExampleUnpack5() {
	a, b, c, d, e := lo.Unpack5(lo.T5("hello", 2, true, foo{bar: "bar"}, 4.2))
	fmt.Printf("%v %v %v %v %v", a, b, c, d, e)
	// Output: hello 2 true {bar} 4.2
}

func ExampleUnpack6() {
	a, b, c, d, e, f := lo.Unpack6(lo.T6("hello", 2, true, foo{bar: "bar"}, 4.2, "plop"))
	fmt.Printf("%v %v %v %v %v %v", a, b, c, d, e, f)
	// Output: hello 2 true {bar} 4.2 plop
}

func ExampleUnpack7() {
	a, b, c, d, e, f, g := lo.Unpack7(lo.T7("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false))
	fmt.Printf("%v %v %v %v %v %v %v", a, b, c, d, e, f, g)
	// Output: hello 2 true {bar} 4.2 plop false
}

func ExampleUnpack8() {
	a, b, c, d, e, f, g, h := lo.Unpack8(lo.T8("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false, 42))
	fmt.Printf("%v %v %v %v %v %v %v %v", a, b, c, d, e, f, g, h)
	// Output: hello 2 true {bar} 4.2 plop false 42
}

func ExampleUnpack9() {
	a, b, c, d, e, f, g, h, i := lo.Unpack9(lo.T9("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false, 42, "hello world"))
	fmt.Printf("%v %v %v %v %v %v %v %v %v", a, b, c, d, e, f, g, h, i)
	// Output: hello 2 true {bar} 4.2 plop false 42 hello world
}

func ExampleZip2() {
	result := lo.Zip2([]string{"hello"}, []int{2})
	fmt.Printf("%v", result)
	// Output: [{hello 2}]
}

func ExampleZip3() {
	result := lo.Zip3([]string{"hello"}, []int{2}, []bool{true})
	fmt.Printf("%v", result)
	// Output: [{hello 2 true}]
}

func ExampleZip4() {
	result := lo.Zip4([]string{"hello"}, []int{2}, []bool{true}, []foo{{bar: "bar"}})
	fmt.Printf("%v", result)
	// Output: [{hello 2 true {bar}}]
}

func ExampleZip5() {
	result := lo.Zip5([]string{"hello"}, []int{2}, []bool{true}, []foo{{bar: "bar"}}, []float64{4.2})
	fmt.Printf("%v", result)
	// Output: [{hello 2 true {bar} 4.2}]
}

func ExampleZip6() {
	result := lo.Zip6([]string{"hello"}, []int{2}, []bool{true}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"})
	fmt.Printf("%v", result)
	// Output: [{hello 2 true {bar} 4.2 plop}]
}

func ExampleZip7() {
	result := lo.Zip7([]string{"hello"}, []int{2}, []bool{true}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"}, []bool{false})
	fmt.Printf("%v", result)
	// Output: [{hello 2 true {bar} 4.2 plop false}]
}

func ExampleZip8() {
	result := lo.Zip8([]string{"hello"}, []int{2}, []bool{true}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"}, []bool{false}, []int{42})
	fmt.Printf("%v", result)
	// Output: [{hello 2 true {bar} 4.2 plop false 42}]
}

func ExampleZip9() {
	result := lo.Zip9([]string{"hello"}, []int{2}, []bool{true}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"}, []bool{false}, []int{42}, []string{"hello world"})
	fmt.Printf("%v", result)
	// Output: [{hello 2 true {bar} 4.2 plop false 42 hello world}]
}

func ExampleUnzip2() {
	a, b := lo.Unzip2([]lo.Tuple2[string, int]{lo.T2("hello", 2)})
	fmt.Printf("%v %v", a, b)
	// Output: [hello] [2]
}

func ExampleUnzip3() {
	a, b, c := lo.Unzip3([]lo.Tuple3[string, int, bool]{lo.T3("hello", 2, true)})
	fmt.Printf("%v %v %v", a, b, c)
	// Output: [hello] [2] [true]
}

func ExampleUnzip4() {
	a, b, c, d := lo.Unzip4([]lo.Tuple4[string, int, bool, foo]{lo.T4("hello", 2, true, foo{bar: "bar"})})
	fmt.Printf("%v %v %v %v", a, b, c, d)
	// Output: [hello] [2] [true] [{bar}]
}

func ExampleUnzip5() {
	a, b, c, d, e := lo.Unzip5([]lo.Tuple5[string, int, bool, foo, float64]{lo.T5("hello", 2, true, foo{bar: "bar"}, 4.2)})
	fmt.Printf("%v %v %v %v %v", a, b, c, d, e)
	// Output: [hello] [2] [true] [{bar}] [4.2]
}

func ExampleUnzip6() {
	a, b, c, d, e, f := lo.Unzip6([]lo.Tuple6[string, int, bool, foo, float64, string]{lo.T6("hello", 2, true, foo{bar: "bar"}, 4.2, "plop")})
	fmt.Printf("%v %v %v %v %v %v", a, b, c, d, e, f)
	// Output: [hello] [2] [true] [{bar}] [4.2] [plop]
}

func ExampleUnzip7() {
	a, b, c, d, e, f, g := lo.Unzip7([]lo.Tuple7[string, int, bool, foo, float64, string, bool]{lo.T7("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false)})
	fmt.Printf("%v %v %v %v %v %v %v", a, b, c, d, e, f, g)
	// Output: [hello] [2] [true] [{bar}] [4.2] [plop] [false]
}

func ExampleUnzip8() {
	a, b, c, d, e, f, g, h := lo.Unzip8([]lo.Tuple8[string, int, bool, foo, float64, string, bool, int]{lo.T8("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false, 42)})
	fmt.Printf("%v %v %v %v %v %v %v %v", a, b, c, d, e, f, g, h)
	// Output: [hello] [2] [true] [{bar}] [4.2] [plop] [false] [42]
}

func ExampleUnzip9() {
	a, b, c, d, e, f, g, h, i := lo.Unzip9([]lo.Tuple9[string, int, bool, foo, float64, string, bool, int, string]{lo.T9("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false, 42, "hello world")})
	fmt.Printf("%v %v %v %v %v %v %v %v %v", a, b, c, d, e, f, g, h, i)
	// Output: [hello] [2] [true] [{bar}] [4.2] [plop] [false] [42] [hello world]
}
