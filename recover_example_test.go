package lo

import (
	"errors"
	"fmt"
)

func ExampleRecover0() {
	err := Recover0(func() {
		panic("something went wrong")
	})
	fmt.Println(err)
	// Output: something went wrong
}

func ExampleRecover0Error() {
	err := Recover0Error(func() error {
		return errors.New("regular error")
	})
	fmt.Println(err)
	// Output: regular error
}

func ExampleRecover1() {
	result, err := Recover1(func() int {
		panic("calculation error")
		return 42
	})
	fmt.Println(result, err)
	// Output: 0 calculation error
}

func ExampleRecover1Error() {
	result, err := Recover1Error(func() (int, error) {
		return 0, errors.New("division by zero")
	})
	fmt.Println(result, err)
	// Output: 0 division by zero
}

func ExampleRecover2() {
	x, y, err := Recover2(func() (int, string) {
		panic("unexpected input")
		return 10, "hello"
	})
	fmt.Println(x, y, err)
	// Output: 0  unexpected input
}

func ExampleRecover2Error() {
	x, y, err := Recover2Error(func() (int, string, error) {
		return 0, "", errors.New("invalid operation")
	})
	fmt.Println(x, y, err)
	// Output: 0  invalid operation
}

func ExampleRecover3() {
	a, b, c, err := Recover3(func() (int, string, bool) {
		panic("critical failure")
		return 1, "test", true
	})
	fmt.Println(a, b, c, err)
	// Output: 0  false critical failure
}

func ExampleRecover3Error() {
	a, b, c, err := Recover3Error(func() (int, string, bool, error) {
		return 0, "", false, errors.New("operation failed")
	})
	fmt.Println(a, b, c, err)
	// Output: 0  false operation failed
}

func ExampleRecover4() {
	w, x, y, z, err := Recover4(func() (int, string, bool, float64) {
		panic("system crash")
		return 1, "test", true, 3.14
	})
	fmt.Println(w, x, y, z, err)
	// Output: 0  false 0 system crash
}

func ExampleRecover4Error() {
	w, x, y, z, err := Recover4Error(func() (int, string, bool, float64, error) {
		return 0, "", false, 0.0, errors.New("calculation error")
	})
	fmt.Println(w, x, y, z, err)
	// Output: 0  false 0 calculation error
}

func ExampleRecover5() {
	v, w, x, y, z, err := Recover5(func() (int, string, bool, float64, []int) {
		panic("out of memory")
		return 1, "test", true, 3.14, []int{1, 2, 3}
	})
	fmt.Println(v, w, x, y, z, err)
	// Output: 0  false 0 [] out of memory
}

func ExampleRecover5Error() {
	v, w, x, y, z, err := Recover5Error(func() (int, string, bool, float64, []int, error) {
		return 0, "", false, 0.0, nil, errors.New("invalid input")
	})
	fmt.Println(v, w, x, y, z, err)
	// Output: 0  false 0 [] invalid input
}

func ExampleRecover6() {
	u, v, w, x, y, z, err := Recover6(func() (int, string, bool, float64, []int, map[string]int) {
		panic("stack overflow")
		return 1, "test", true, 3.14, []int{1, 2, 3}, map[string]int{"a": 1}
	})
	fmt.Println(u, v, w, x, y, z, err)
	// Output: 0  false 0 [] map[] stack overflow
}

func ExampleRecover6Error() {
	u, v, w, x, y, z, err := Recover6Error(func() (int, string, bool, float64, []int, map[string]int, error) {
		return 0, "", false, 0.0, nil, nil, errors.New("operation timeout")
	})
	fmt.Println(u, v, w, x, y, z, err)
	// Output: 0  false 0 [] map[] operation timeout
}

func ExampleTypedRecover() {
	nestedMyErrorPanic := func() (err error) {
		defer TypedRecover[myError](&err, false)
		panic(myError{})
	}
	err := nestedMyErrorPanic()

	nestedStringPanic := func() (err error) {
		defer TypedRecover[error](&err, true)
		panic("test string panic")
	}
	err2 := nestedStringPanic()

	fmt.Println(err, err2)
	// Output: my error test string panic
}

func ExampleRecover0Typed() {
	err := Recover0Typed[myError](func() {
		panic(myError{})
	})
	fmt.Println(err)
	// Output: my error
}

func ExampleRecover0ErrorTyped() {
	err := Recover0ErrorTyped[myError](func() error {
		panic("critical failure")
	}, true)
	fmt.Println(err)
	// Output: critical failure
}

func ExampleRecover1Typed() {
	result, err := Recover1Typed[string, error](func() string {
		panic("processing error")
		return "success"
	}, true)
	fmt.Println(result, err)
	// Output:  processing error
}

func ExampleRecover1ErrorTyped() {
	result, err := Recover1ErrorTyped[string, error](func() (string, error) {
		panic("processing error")
		return "success", nil
	}, true)
	fmt.Println(result, err)
	// Output:  processing error
}

func ExampleRecover2Typed() {
	a, b, err := Recover2Typed[int, string, error](func() (int, string) {
		panic("invalid input")
		return 10, "hello"
	}, true)
	fmt.Println(a, b, err)
	// Output: 0  invalid input
}

func ExampleRecover2ErrorTyped() {
	x, y, err := Recover2ErrorTyped[bool, float64, error](func() (bool, float64, error) {
		panic("unexpected condition")
		return true, 3.14, nil
	}, true)
	fmt.Println(x, y, err)
	// Output: false 0 unexpected condition
}

func ExampleRecover3Typed() {
	a, b, c, err := Recover3Typed[int, string, bool, error](func() (int, string, bool) {
		panic("data corruption")
		return 42, "test", true
	}, true)
	fmt.Println(a, b, c, err)
	// Output: 0  false data corruption
}

func ExampleRecover3ErrorTyped() {
	x, y, z, err := Recover3ErrorTyped[int, string, bool, error](func() (int, string, bool, error) {
		panic("calculation error")
		return 1, "two", true, nil
	}, true)
	fmt.Println(x, y, z, err)
	// Output: 0  false calculation error
}

func ExampleRecover4Typed() {
	a, b, c, d, err := Recover4Typed[int, string, bool, float64, error](func() (int, string, bool, float64) {
		panic("unexpected state")
		return 1, "test", true, 3.14
	}, true)
	fmt.Println(a, b, c, d, err)
	// Output: 0  false 0 unexpected state
}

func ExampleRecover4ErrorTyped() {
	w, x, y, z, err := Recover4ErrorTyped[int, string, bool, float64, error](func() (int, string, bool, float64, error) {
		panic("division by zero")
		return 1, "two", true, 3.14, nil
	}, true)
	fmt.Println(w, x, y, z, err)
	// Output: 0  false 0 division by zero
}

func ExampleRecover5Typed() {
	a, b, c, d, e, err := Recover5Typed[int, string, bool, float64, []int, error](func() (int, string, bool, float64, []int) {
		panic("memory allocation failure")
		return 1, "test", true, 3.14, []int{1, 2, 3}
	}, true)
	fmt.Println(a, b, c, d, e, err)
	// Output: 0  false 0 [] memory allocation failure
}

func ExampleRecover5ErrorTyped() {
	v, w, x, y, z, err := Recover5ErrorTyped[int, string, bool, float64, []int, error](func() (int, string, bool, float64, []int, error) {
		panic("stack overflow")
		return 1, "two", true, 3.14, []int{1, 2, 3}, nil
	}, true)
	fmt.Println(v, w, x, y, z, err)
	// Output: 0  false 0 [] stack overflow
}

func ExampleRecover6Typed() {
	a, b, c, d, e, f, err := Recover6Typed[int, string, bool, float64, []int, map[string]int, error](func() (int, string, bool, float64, []int, map[string]int) {
		panic("network failure")
		return 1, "test", true, 3.14, []int{1, 2, 3}, map[string]int{"a": 1}
	}, true)
	fmt.Println(a, b, c, d, e, f, err)
	// Output: 0  false 0 [] map[] network failure
}

func ExampleRecover6ErrorTyped() {
	u, v, w, x, y, z, err := Recover6ErrorTyped[int, string, bool, float64, []int, map[string]int, error](func() (int, string, bool, float64, []int, map[string]int, error) {
		panic("database connection lost")
		return 1, "two", true, 3.14, []int{1, 2, 3}, map[string]int{"a": 1}, nil
	}, true)
	fmt.Println(u, v, w, x, y, z, err)
	// Output: 0  false 0 [] map[] database connection lost
}
