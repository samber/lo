package lo

import (
	"fmt"
)

func ExampleT2() {
	result := T2("hello", 2)
	fmt.Printf("%v %v", result.A, result.B)
	// Output: hello 2
}

func ExampleT3() {
	result := T3("hello", 2, true)
	fmt.Printf("%v %v %v", result.A, result.B, result.C)
	// Output: hello 2 true
}

func ExampleT4() {
	result := T4("hello", 2, true, foo{bar: "bar"})
	fmt.Printf("%v %v %v %v", result.A, result.B, result.C, result.D)
	// Output: hello 2 true {bar}
}

func ExampleT5() {
	result := T5("hello", 2, true, foo{bar: "bar"}, 4.2)
	fmt.Printf("%v %v %v %v %v", result.A, result.B, result.C, result.D, result.E)
	// Output: hello 2 true {bar} 4.2
}

func ExampleT6() {
	result := T6("hello", 2, true, foo{bar: "bar"}, 4.2, "plop")
	fmt.Printf("%v %v %v %v %v %v", result.A, result.B, result.C, result.D, result.E, result.F)
	// Output: hello 2 true {bar} 4.2 plop
}

func ExampleT7() {
	result := T7("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false)
	fmt.Printf("%v %v %v %v %v %v %v", result.A, result.B, result.C, result.D, result.E, result.F, result.G)
	// Output: hello 2 true {bar} 4.2 plop false
}

func ExampleT8() {
	result := T8("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false, 42)
	fmt.Printf("%v %v %v %v %v %v %v %v", result.A, result.B, result.C, result.D, result.E, result.F, result.G, result.H)
	// Output: hello 2 true {bar} 4.2 plop false 42
}

func ExampleT9() {
	result := T9("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false, 42, "hello world")
	fmt.Printf("%v %v %v %v %v %v %v %v %v", result.A, result.B, result.C, result.D, result.E, result.F, result.G, result.H, result.I)
	// Output: hello 2 true {bar} 4.2 plop false 42 hello world
}

func ExampleUnpack2() {
	a, b := Unpack2(T2("hello", 2))
	fmt.Printf("%v %v", a, b)
	// Output: hello 2
}

func ExampleUnpack3() {
	a, b, c := Unpack3(T3("hello", 2, true))
	fmt.Printf("%v %v %v", a, b, c)
	// Output: hello 2 true
}

func ExampleUnpack4() {
	a, b, c, d := Unpack4(T4("hello", 2, true, foo{bar: "bar"}))
	fmt.Printf("%v %v %v %v", a, b, c, d)
	// Output: hello 2 true {bar}
}

func ExampleUnpack5() {
	a, b, c, d, e := Unpack5(T5("hello", 2, true, foo{bar: "bar"}, 4.2))
	fmt.Printf("%v %v %v %v %v", a, b, c, d, e)
	// Output: hello 2 true {bar} 4.2
}

func ExampleUnpack6() {
	a, b, c, d, e, f := Unpack6(T6("hello", 2, true, foo{bar: "bar"}, 4.2, "plop"))
	fmt.Printf("%v %v %v %v %v %v", a, b, c, d, e, f)
	// Output: hello 2 true {bar} 4.2 plop
}

func ExampleUnpack7() {
	a, b, c, d, e, f, g := Unpack7(T7("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false))
	fmt.Printf("%v %v %v %v %v %v %v", a, b, c, d, e, f, g)
	// Output: hello 2 true {bar} 4.2 plop false
}

func ExampleUnpack8() {
	a, b, c, d, e, f, g, h := Unpack8(T8("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false, 42))
	fmt.Printf("%v %v %v %v %v %v %v %v", a, b, c, d, e, f, g, h)
	// Output: hello 2 true {bar} 4.2 plop false 42
}

func ExampleUnpack9() {
	a, b, c, d, e, f, g, h, i := Unpack9(T9("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false, 42, "hello world"))
	fmt.Printf("%v %v %v %v %v %v %v %v %v", a, b, c, d, e, f, g, h, i)
	// Output: hello 2 true {bar} 4.2 plop false 42 hello world
}

func ExampleZip2() {
	result := Zip2([]string{"hello"}, []int{2})
	fmt.Printf("%v", result)
	// Output: [{hello 2}]
}

func ExampleZip3() {
	result := Zip3([]string{"hello"}, []int{2}, []bool{true})
	fmt.Printf("%v", result)
	// Output: [{hello 2 true}]
}

func ExampleZip4() {
	result := Zip4([]string{"hello"}, []int{2}, []bool{true}, []foo{{bar: "bar"}})
	fmt.Printf("%v", result)
	// Output: [{hello 2 true {bar}}]
}

func ExampleZip5() {
	result := Zip5([]string{"hello"}, []int{2}, []bool{true}, []foo{{bar: "bar"}}, []float64{4.2})
	fmt.Printf("%v", result)
	// Output: [{hello 2 true {bar} 4.2}]
}

func ExampleZip6() {
	result := Zip6([]string{"hello"}, []int{2}, []bool{true}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"})
	fmt.Printf("%v", result)
	// Output: [{hello 2 true {bar} 4.2 plop}]
}

func ExampleZip7() {
	result := Zip7([]string{"hello"}, []int{2}, []bool{true}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"}, []bool{false})
	fmt.Printf("%v", result)
	// Output: [{hello 2 true {bar} 4.2 plop false}]
}

func ExampleZip8() {
	result := Zip8([]string{"hello"}, []int{2}, []bool{true}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"}, []bool{false}, []int{42})
	fmt.Printf("%v", result)
	// Output: [{hello 2 true {bar} 4.2 plop false 42}]
}

func ExampleZip9() {
	result := Zip9([]string{"hello"}, []int{2}, []bool{true}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"}, []bool{false}, []int{42}, []string{"hello world"})
	fmt.Printf("%v", result)
	// Output: [{hello 2 true {bar} 4.2 plop false 42 hello world}]
}

func ExampleUnzip2() {
	a, b := Unzip2([]Tuple2[string, int]{T2("hello", 2)})
	fmt.Printf("%v %v", a, b)
	// Output: [hello] [2]
}

func ExampleUnzip3() {
	a, b, c := Unzip3([]Tuple3[string, int, bool]{T3("hello", 2, true)})
	fmt.Printf("%v %v %v", a, b, c)
	// Output: [hello] [2] [true]
}

func ExampleUnzip4() {
	a, b, c, d := Unzip4([]Tuple4[string, int, bool, foo]{T4("hello", 2, true, foo{bar: "bar"})})
	fmt.Printf("%v %v %v %v", a, b, c, d)
	// Output: [hello] [2] [true] [{bar}]
}

func ExampleUnzip5() {
	a, b, c, d, e := Unzip5([]Tuple5[string, int, bool, foo, float64]{T5("hello", 2, true, foo{bar: "bar"}, 4.2)})
	fmt.Printf("%v %v %v %v %v", a, b, c, d, e)
	// Output: [hello] [2] [true] [{bar}] [4.2]
}

func ExampleUnzip6() {
	a, b, c, d, e, f := Unzip6([]Tuple6[string, int, bool, foo, float64, string]{T6("hello", 2, true, foo{bar: "bar"}, 4.2, "plop")})
	fmt.Printf("%v %v %v %v %v %v", a, b, c, d, e, f)
	// Output: [hello] [2] [true] [{bar}] [4.2] [plop]
}

func ExampleUnzip7() {
	a, b, c, d, e, f, g := Unzip7([]Tuple7[string, int, bool, foo, float64, string, bool]{T7("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false)})
	fmt.Printf("%v %v %v %v %v %v %v", a, b, c, d, e, f, g)
	// Output: [hello] [2] [true] [{bar}] [4.2] [plop] [false]
}

func ExampleUnzip8() {
	a, b, c, d, e, f, g, h := Unzip8([]Tuple8[string, int, bool, foo, float64, string, bool, int]{T8("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false, 42)})
	fmt.Printf("%v %v %v %v %v %v %v %v", a, b, c, d, e, f, g, h)
	// Output: [hello] [2] [true] [{bar}] [4.2] [plop] [false] [42]
}

func ExampleUnzip9() {
	a, b, c, d, e, f, g, h, i := Unzip9([]Tuple9[string, int, bool, foo, float64, string, bool, int, string]{T9("hello", 2, true, foo{bar: "bar"}, 4.2, "plop", false, 42, "hello world")})
	fmt.Printf("%v %v %v %v %v %v %v %v %v", a, b, c, d, e, f, g, h, i)
	// Output: [hello] [2] [true] [{bar}] [4.2] [plop] [false] [42] [hello world]
}

func ExampleCrossJoin2() {
	result := CrossJoin2([]string{"a", "b"}, []int{1, 2, 3, 4})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// {a 1}
	// {a 2}
	// {a 3}
	// {a 4}
	// {b 1}
	// {b 2}
	// {b 3}
	// {b 4}
}

func ExampleCrossJoin3() {
	result := CrossJoin3([]string{"a", "b"}, []int{1, 2, 3, 4}, []bool{true, false})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// {a 1 true}
	// {a 1 false}
	// {a 2 true}
	// {a 2 false}
	// {a 3 true}
	// {a 3 false}
	// {a 4 true}
	// {a 4 false}
	// {b 1 true}
	// {b 1 false}
	// {b 2 true}
	// {b 2 false}
	// {b 3 true}
	// {b 3 false}
	// {b 4 true}
	// {b 4 false}
}

func ExampleCrossJoin4() {
	result := CrossJoin4([]string{"a", "b"}, []int{1, 2, 3, 4}, []bool{true, false}, []foo{{bar: "bar"}})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// {a 1 true {bar}}
	// {a 1 false {bar}}
	// {a 2 true {bar}}
	// {a 2 false {bar}}
	// {a 3 true {bar}}
	// {a 3 false {bar}}
	// {a 4 true {bar}}
	// {a 4 false {bar}}
	// {b 1 true {bar}}
	// {b 1 false {bar}}
	// {b 2 true {bar}}
	// {b 2 false {bar}}
	// {b 3 true {bar}}
	// {b 3 false {bar}}
	// {b 4 true {bar}}
	// {b 4 false {bar}}
}

func ExampleCrossJoin5() {
	result := CrossJoin5([]string{"a", "b"}, []int{1, 2, 3, 4}, []bool{true, false}, []foo{{bar: "bar"}}, []float64{4.2})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// {a 1 true {bar} 4.2}
	// {a 1 false {bar} 4.2}
	// {a 2 true {bar} 4.2}
	// {a 2 false {bar} 4.2}
	// {a 3 true {bar} 4.2}
	// {a 3 false {bar} 4.2}
	// {a 4 true {bar} 4.2}
	// {a 4 false {bar} 4.2}
	// {b 1 true {bar} 4.2}
	// {b 1 false {bar} 4.2}
	// {b 2 true {bar} 4.2}
	// {b 2 false {bar} 4.2}
	// {b 3 true {bar} 4.2}
	// {b 3 false {bar} 4.2}
	// {b 4 true {bar} 4.2}
	// {b 4 false {bar} 4.2}
}

func ExampleCrossJoin6() {
	result := CrossJoin6([]string{"a", "b"}, []int{1, 2, 3, 4}, []bool{true, false}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// {a 1 true {bar} 4.2 plop}
	// {a 1 false {bar} 4.2 plop}
	// {a 2 true {bar} 4.2 plop}
	// {a 2 false {bar} 4.2 plop}
	// {a 3 true {bar} 4.2 plop}
	// {a 3 false {bar} 4.2 plop}
	// {a 4 true {bar} 4.2 plop}
	// {a 4 false {bar} 4.2 plop}
	// {b 1 true {bar} 4.2 plop}
	// {b 1 false {bar} 4.2 plop}
	// {b 2 true {bar} 4.2 plop}
	// {b 2 false {bar} 4.2 plop}
	// {b 3 true {bar} 4.2 plop}
	// {b 3 false {bar} 4.2 plop}
	// {b 4 true {bar} 4.2 plop}
	// {b 4 false {bar} 4.2 plop}
}

func ExampleCrossJoin7() {
	result := CrossJoin7([]string{"a", "b"}, []int{1, 2, 3, 4}, []bool{true, false}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"}, []bool{false})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// {a 1 true {bar} 4.2 plop false}
	// {a 1 false {bar} 4.2 plop false}
	// {a 2 true {bar} 4.2 plop false}
	// {a 2 false {bar} 4.2 plop false}
	// {a 3 true {bar} 4.2 plop false}
	// {a 3 false {bar} 4.2 plop false}
	// {a 4 true {bar} 4.2 plop false}
	// {a 4 false {bar} 4.2 plop false}
	// {b 1 true {bar} 4.2 plop false}
	// {b 1 false {bar} 4.2 plop false}
	// {b 2 true {bar} 4.2 plop false}
	// {b 2 false {bar} 4.2 plop false}
	// {b 3 true {bar} 4.2 plop false}
	// {b 3 false {bar} 4.2 plop false}
	// {b 4 true {bar} 4.2 plop false}
	// {b 4 false {bar} 4.2 plop false}
}

func ExampleCrossJoin8() {
	result := CrossJoin8([]string{"a", "b"}, []int{1, 2, 3, 4}, []bool{true, false}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"}, []bool{false}, []int{42})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// {a 1 true {bar} 4.2 plop false 42}
	// {a 1 false {bar} 4.2 plop false 42}
	// {a 2 true {bar} 4.2 plop false 42}
	// {a 2 false {bar} 4.2 plop false 42}
	// {a 3 true {bar} 4.2 plop false 42}
	// {a 3 false {bar} 4.2 plop false 42}
	// {a 4 true {bar} 4.2 plop false 42}
	// {a 4 false {bar} 4.2 plop false 42}
	// {b 1 true {bar} 4.2 plop false 42}
	// {b 1 false {bar} 4.2 plop false 42}
	// {b 2 true {bar} 4.2 plop false 42}
	// {b 2 false {bar} 4.2 plop false 42}
	// {b 3 true {bar} 4.2 plop false 42}
	// {b 3 false {bar} 4.2 plop false 42}
	// {b 4 true {bar} 4.2 plop false 42}
	// {b 4 false {bar} 4.2 plop false 42}
}

func ExampleCrossJoin9() {
	result := CrossJoin9([]string{"a", "b"}, []int{1, 2, 3, 4}, []bool{true, false}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"}, []bool{false}, []int{42}, []string{"hello world"})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// {a 1 true {bar} 4.2 plop false 42 hello world}
	// {a 1 false {bar} 4.2 plop false 42 hello world}
	// {a 2 true {bar} 4.2 plop false 42 hello world}
	// {a 2 false {bar} 4.2 plop false 42 hello world}
	// {a 3 true {bar} 4.2 plop false 42 hello world}
	// {a 3 false {bar} 4.2 plop false 42 hello world}
	// {a 4 true {bar} 4.2 plop false 42 hello world}
	// {a 4 false {bar} 4.2 plop false 42 hello world}
	// {b 1 true {bar} 4.2 plop false 42 hello world}
	// {b 1 false {bar} 4.2 plop false 42 hello world}
	// {b 2 true {bar} 4.2 plop false 42 hello world}
	// {b 2 false {bar} 4.2 plop false 42 hello world}
	// {b 3 true {bar} 4.2 plop false 42 hello world}
	// {b 3 false {bar} 4.2 plop false 42 hello world}
	// {b 4 true {bar} 4.2 plop false 42 hello world}
	// {b 4 false {bar} 4.2 plop false 42 hello world}
}

func ExampleCrossJoinBy2() {
	result := CrossJoinBy2([]string{"a", "b"}, []int{1, 2, 3, 4}, func(a string, b int) string {
		return fmt.Sprintf("%v-%v", a, b)
	})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// a-1
	// a-2
	// a-3
	// a-4
	// b-1
	// b-2
	// b-3
	// b-4
}

func ExampleCrossJoinBy3() {
	result := CrossJoinBy3([]string{"a", "b"}, []int{1, 2, 3, 4}, []bool{true, false}, func(a string, b int, c bool) string {
		return fmt.Sprintf("%v-%v-%v", a, b, c)
	})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// a-1-true
	// a-1-false
	// a-2-true
	// a-2-false
	// a-3-true
	// a-3-false
	// a-4-true
	// a-4-false
	// b-1-true
	// b-1-false
	// b-2-true
	// b-2-false
	// b-3-true
	// b-3-false
	// b-4-true
	// b-4-false
}

func ExampleCrossJoinBy4() {
	result := CrossJoinBy4([]string{"a", "b"}, []int{1, 2, 3, 4}, []bool{true, false}, []foo{{bar: "bar"}}, func(a string, b int, c bool, d foo) string {
		return fmt.Sprintf("%v-%v-%v-%v", a, b, c, d)
	})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// a-1-true-{bar}
	// a-1-false-{bar}
	// a-2-true-{bar}
	// a-2-false-{bar}
	// a-3-true-{bar}
	// a-3-false-{bar}
	// a-4-true-{bar}
	// a-4-false-{bar}
	// b-1-true-{bar}
	// b-1-false-{bar}
	// b-2-true-{bar}
	// b-2-false-{bar}
	// b-3-true-{bar}
	// b-3-false-{bar}
	// b-4-true-{bar}
	// b-4-false-{bar}
}

func ExampleCrossJoinBy5() {
	result := CrossJoinBy5([]string{"a", "b"}, []int{1, 2, 3, 4}, []bool{true, false}, []foo{{bar: "bar"}}, []float64{4.2}, func(a string, b int, c bool, d foo, e float64) string {
		return fmt.Sprintf("%v-%v-%v-%v-%v", a, b, c, d, e)
	})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// a-1-true-{bar}-4.2
	// a-1-false-{bar}-4.2
	// a-2-true-{bar}-4.2
	// a-2-false-{bar}-4.2
	// a-3-true-{bar}-4.2
	// a-3-false-{bar}-4.2
	// a-4-true-{bar}-4.2
	// a-4-false-{bar}-4.2
	// b-1-true-{bar}-4.2
	// b-1-false-{bar}-4.2
	// b-2-true-{bar}-4.2
	// b-2-false-{bar}-4.2
	// b-3-true-{bar}-4.2
	// b-3-false-{bar}-4.2
	// b-4-true-{bar}-4.2
	// b-4-false-{bar}-4.2
}

func ExampleCrossJoinBy6() {
	result := CrossJoinBy6([]string{"a", "b"}, []int{1, 2, 3, 4}, []bool{true, false}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"}, func(a string, b int, c bool, d foo, e float64, f string) string {
		return fmt.Sprintf("%v-%v-%v-%v-%v-%v", a, b, c, d, e, f)
	})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// a-1-true-{bar}-4.2-plop
	// a-1-false-{bar}-4.2-plop
	// a-2-true-{bar}-4.2-plop
	// a-2-false-{bar}-4.2-plop
	// a-3-true-{bar}-4.2-plop
	// a-3-false-{bar}-4.2-plop
	// a-4-true-{bar}-4.2-plop
	// a-4-false-{bar}-4.2-plop
	// b-1-true-{bar}-4.2-plop
	// b-1-false-{bar}-4.2-plop
	// b-2-true-{bar}-4.2-plop
	// b-2-false-{bar}-4.2-plop
	// b-3-true-{bar}-4.2-plop
	// b-3-false-{bar}-4.2-plop
	// b-4-true-{bar}-4.2-plop
	// b-4-false-{bar}-4.2-plop
}

func ExampleCrossJoinBy7() {
	result := CrossJoinBy7([]string{"a", "b"}, []int{1, 2, 3, 4}, []bool{true, false}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"}, []bool{false}, func(a string, b int, c bool, d foo, e float64, f string, g bool) string {
		return fmt.Sprintf("%v-%v-%v-%v-%v-%v-%v", a, b, c, d, e, f, g)
	})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// a-1-true-{bar}-4.2-plop-false
	// a-1-false-{bar}-4.2-plop-false
	// a-2-true-{bar}-4.2-plop-false
	// a-2-false-{bar}-4.2-plop-false
	// a-3-true-{bar}-4.2-plop-false
	// a-3-false-{bar}-4.2-plop-false
	// a-4-true-{bar}-4.2-plop-false
	// a-4-false-{bar}-4.2-plop-false
	// b-1-true-{bar}-4.2-plop-false
	// b-1-false-{bar}-4.2-plop-false
	// b-2-true-{bar}-4.2-plop-false
	// b-2-false-{bar}-4.2-plop-false
	// b-3-true-{bar}-4.2-plop-false
	// b-3-false-{bar}-4.2-plop-false
	// b-4-true-{bar}-4.2-plop-false
	// b-4-false-{bar}-4.2-plop-false
}

func ExampleCrossJoinBy8() {
	result := CrossJoinBy8([]string{"a", "b"}, []int{1, 2, 3, 4}, []bool{true, false}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"}, []bool{false}, []int{42}, func(a string, b int, c bool, d foo, e float64, f string, g bool, h int) string {
		return fmt.Sprintf("%v-%v-%v-%v-%v-%v-%v-%v", a, b, c, d, e, f, g, h)
	})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// a-1-true-{bar}-4.2-plop-false-42
	// a-1-false-{bar}-4.2-plop-false-42
	// a-2-true-{bar}-4.2-plop-false-42
	// a-2-false-{bar}-4.2-plop-false-42
	// a-3-true-{bar}-4.2-plop-false-42
	// a-3-false-{bar}-4.2-plop-false-42
	// a-4-true-{bar}-4.2-plop-false-42
	// a-4-false-{bar}-4.2-plop-false-42
	// b-1-true-{bar}-4.2-plop-false-42
	// b-1-false-{bar}-4.2-plop-false-42
	// b-2-true-{bar}-4.2-plop-false-42
	// b-2-false-{bar}-4.2-plop-false-42
	// b-3-true-{bar}-4.2-plop-false-42
	// b-3-false-{bar}-4.2-plop-false-42
	// b-4-true-{bar}-4.2-plop-false-42
	// b-4-false-{bar}-4.2-plop-false-42
}

func ExampleCrossJoinBy9() {
	result := CrossJoinBy9([]string{"a", "b"}, []int{1, 2, 3, 4}, []bool{true, false}, []foo{{bar: "bar"}}, []float64{4.2}, []string{"plop"}, []bool{false}, []int{42}, []string{"hello world"}, func(a string, b int, c bool, d foo, e float64, f string, g bool, h int, i string) string {
		return fmt.Sprintf("%v-%v-%v-%v-%v-%v-%v-%v-%v", a, b, c, d, e, f, g, h, i)
	})
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
	// Output:
	// a-1-true-{bar}-4.2-plop-false-42-hello world
	// a-1-false-{bar}-4.2-plop-false-42-hello world
	// a-2-true-{bar}-4.2-plop-false-42-hello world
	// a-2-false-{bar}-4.2-plop-false-42-hello world
	// a-3-true-{bar}-4.2-plop-false-42-hello world
	// a-3-false-{bar}-4.2-plop-false-42-hello world
	// a-4-true-{bar}-4.2-plop-false-42-hello world
	// a-4-false-{bar}-4.2-plop-false-42-hello world
	// b-1-true-{bar}-4.2-plop-false-42-hello world
	// b-1-false-{bar}-4.2-plop-false-42-hello world
	// b-2-true-{bar}-4.2-plop-false-42-hello world
	// b-2-false-{bar}-4.2-plop-false-42-hello world
	// b-3-true-{bar}-4.2-plop-false-42-hello world
	// b-3-false-{bar}-4.2-plop-false-42-hello world
	// b-4-true-{bar}-4.2-plop-false-42-hello world
	// b-4-false-{bar}-4.2-plop-false-42-hello world
}
