package lo_test

import (
	"fmt"

	"github.com/samber/lo"
)

func ExampleValidate() {
	i := 42

	err1 := lo.Validate(i < 0, "expected %d < 0", i)
	err2 := lo.Validate(i > 0, "expected %d > 0", i)

	fmt.Printf("%v\n%v", err1, err2)
	// Output:
	// expected 42 < 0
	// <nil>
}

func ExampleMust() {
	defer func() {
		_ = recover()
	}()

	// won't panic
	lo.Must(42, nil)

	// won't panic
	cb := func() (int, error) {
		return 42, nil
	}
	lo.Must(cb())

	// will panic
	lo.Must(42, fmt.Errorf("my error"))

	// will panic with error message
	lo.Must(42, fmt.Errorf("world"), "hello")
}

func ExampleMust0() {
	defer func() {
		_ = recover()
	}()

	// won't panic
	lo.Must0(nil)

	// will panic
	lo.Must0(fmt.Errorf("my error"))

	// will panic with error message
	lo.Must0(fmt.Errorf("world"), "hello")
}

func ExampleMust1() {
	defer func() {
		_ = recover()
	}()

	// won't panic
	lo.Must1(42, nil)

	// won't panic
	cb := func() (int, error) {
		return 42, nil
	}
	lo.Must1(cb())

	// will panic
	lo.Must1(42, fmt.Errorf("my error"))

	// will panic with error message
	lo.Must1(42, fmt.Errorf("world"), "hello")
}

func ExampleMust2() {
	defer func() {
		_ = recover()
	}()

	// won't panic
	lo.Must2(42, "hello", nil)

	// will panic
	lo.Must2(42, "hello", fmt.Errorf("my error"))

	// will panic with error message
	lo.Must2(42, "hello", fmt.Errorf("world"), "hello")
}

func ExampleMust3() {
	defer func() {
		_ = recover()
	}()

	// won't panic
	lo.Must3(42, "hello", 4.2, nil)

	// will panic
	lo.Must3(42, "hello", 4.2, fmt.Errorf("my error"))

	// will panic with error message
	lo.Must3(42, "hello", 4.2, fmt.Errorf("world"), "hello")
}

func ExampleMust4() {
	defer func() {
		_ = recover()
	}()

	// won't panic
	lo.Must4(42, "hello", 4.2, true, nil)

	// will panic
	lo.Must4(42, "hello", 4.2, true, fmt.Errorf("my error"))

	// will panic with error message
	lo.Must4(42, "hello", 4.2, true, fmt.Errorf("world"), "hello")
}

func ExampleMust5() {
	defer func() {
		_ = recover()
	}()

	// won't panic
	lo.Must5(42, "hello", 4.2, true, foo{}, nil)

	// will panic
	lo.Must5(42, "hello", 4.2, true, foo{}, fmt.Errorf("my error"))

	// will panic with error message
	lo.Must5(42, "hello", 4.2, true, foo{}, fmt.Errorf("world"), "hello")
}

func ExampleMust6() {
	defer func() {
		_ = recover()
	}()

	// won't panic
	lo.Must5(42, "hello", 4.2, true, foo{}, "foobar", nil)

	// will panic
	lo.Must5(42, "hello", 4.2, true, foo{}, "foobar", fmt.Errorf("my error"))

	// will panic with error message
	lo.Must5(42, "hello", 4.2, true, foo{}, "foobar", fmt.Errorf("world"), "hello")
}

func ExampleTry() {
	ok1 := lo.Try(func() error {
		return nil
	})
	ok2 := lo.Try(func() error {
		return fmt.Errorf("my error")
	})
	ok3 := lo.Try(func() error {
		panic("my error")
	})

	fmt.Printf("%v\n", ok1)
	fmt.Printf("%v\n", ok2)
	fmt.Printf("%v\n", ok3)
	// Output:
	// true
	// false
	// false
}

func ExampleTry1() {
	ok1 := lo.Try1(func() error {
		return nil
	})
	ok2 := lo.Try1(func() error {
		return fmt.Errorf("my error")
	})
	ok3 := lo.Try1(func() error {
		panic("my error")
	})

	fmt.Printf("%v\n", ok1)
	fmt.Printf("%v\n", ok2)
	fmt.Printf("%v\n", ok3)
	// Output:
	// true
	// false
	// false
}

func ExampleTry2() {
	ok1 := lo.Try2(func() (int, error) {
		return 42, nil
	})
	ok2 := lo.Try2(func() (int, error) {
		return 42, fmt.Errorf("my error")
	})
	ok3 := lo.Try2(func() (int, error) {
		panic("my error")
	})

	fmt.Printf("%v\n", ok1)
	fmt.Printf("%v\n", ok2)
	fmt.Printf("%v\n", ok3)
	// Output:
	// true
	// false
	// false
}

func ExampleTry3() {
	ok1 := lo.Try3(func() (int, string, error) {
		return 42, "foobar", nil
	})
	ok2 := lo.Try3(func() (int, string, error) {
		return 42, "foobar", fmt.Errorf("my error")
	})
	ok3 := lo.Try3(func() (int, string, error) {
		panic("my error")
	})

	fmt.Printf("%v\n", ok1)
	fmt.Printf("%v\n", ok2)
	fmt.Printf("%v\n", ok3)
	// Output:
	// true
	// false
	// false
}

func ExampleTry4() {
	ok1 := lo.Try4(func() (int, string, float64, error) {
		return 42, "foobar", 4.2, nil
	})
	ok2 := lo.Try4(func() (int, string, float64, error) {
		return 42, "foobar", 4.2, fmt.Errorf("my error")
	})
	ok3 := lo.Try4(func() (int, string, float64, error) {
		panic("my error")
	})

	fmt.Printf("%v\n", ok1)
	fmt.Printf("%v\n", ok2)
	fmt.Printf("%v\n", ok3)
	// Output:
	// true
	// false
	// false
}

func ExampleTry5() {
	ok1 := lo.Try5(func() (int, string, float64, bool, error) {
		return 42, "foobar", 4.2, true, nil
	})
	ok2 := lo.Try5(func() (int, string, float64, bool, error) {
		return 42, "foobar", 4.2, true, fmt.Errorf("my error")
	})
	ok3 := lo.Try5(func() (int, string, float64, bool, error) {
		panic("my error")
	})

	fmt.Printf("%v\n", ok1)
	fmt.Printf("%v\n", ok2)
	fmt.Printf("%v\n", ok3)
	// Output:
	// true
	// false
	// false
}

func ExampleTry6() {
	ok1 := lo.Try6(func() (int, string, float64, bool, foo, error) {
		return 42, "foobar", 4.2, true, foo{}, nil
	})
	ok2 := lo.Try6(func() (int, string, float64, bool, foo, error) {
		return 42, "foobar", 4.2, true, foo{}, fmt.Errorf("my error")
	})
	ok3 := lo.Try6(func() (int, string, float64, bool, foo, error) {
		panic("my error")
	})

	fmt.Printf("%v\n", ok1)
	fmt.Printf("%v\n", ok2)
	fmt.Printf("%v\n", ok3)
	// Output:
	// true
	// false
	// false
}

func ExampleTryOr() {
	value1, ok1 := lo.TryOr(func() (int, error) {
		return 42, nil
	}, 21)
	value2, ok2 := lo.TryOr(func() (int, error) {
		return 42, fmt.Errorf("my error")
	}, 21)
	value3, ok3 := lo.TryOr(func() (int, error) {
		panic("my error")
	}, 21)

	fmt.Printf("%v %v\n", value1, ok1)
	fmt.Printf("%v %v\n", value2, ok2)
	fmt.Printf("%v %v\n", value3, ok3)
	// Output:
	// 42 true
	// 21 false
	// 21 false
}

func ExampleTryOr1() {
	value1, ok1 := lo.TryOr1(func() (int, error) {
		return 42, nil
	}, 21)
	value2, ok2 := lo.TryOr1(func() (int, error) {
		return 42, fmt.Errorf("my error")
	}, 21)
	value3, ok3 := lo.TryOr1(func() (int, error) {
		panic("my error")
	}, 21)

	fmt.Printf("%v %v\n", value1, ok1)
	fmt.Printf("%v %v\n", value2, ok2)
	fmt.Printf("%v %v\n", value3, ok3)
	// Output:
	// 42 true
	// 21 false
	// 21 false
}

func ExampleTryOr2() {
	value1, value2, ok3 := lo.TryOr2(func() (int, string, error) {
		panic("my error")
	}, 21, "hello")

	fmt.Printf("%v %v %v\n", value1, value2, ok3)
	// Output: 21 hello false
}

func ExampleTryOr3() {
	value1, value2, value3, ok3 := lo.TryOr3(func() (int, string, bool, error) {
		panic("my error")
	}, 21, "hello", false)

	fmt.Printf("%v %v %v %v\n", value1, value2, value3, ok3)
	// Output: 21 hello false false
}

func ExampleTryOr4() {
	value1, value2, value3, value4, ok3 := lo.TryOr4(func() (int, string, bool, foo, error) {
		panic("my error")
	}, 21, "hello", false, foo{bar: "bar"})

	fmt.Printf("%v %v %v %v %v\n", value1, value2, value3, value4, ok3)
	// Output: 21 hello false {bar} false
}

func ExampleTryOr5() {
	value1, value2, value3, value4, value5, ok3 := lo.TryOr5(func() (int, string, bool, foo, float64, error) {
		panic("my error")
	}, 21, "hello", false, foo{bar: "bar"}, 4.2)

	fmt.Printf("%v %v %v %v %v %v\n", value1, value2, value3, value4, value5, ok3)
	// Output: 21 hello false {bar} 4.2 false
}

func ExampleTryOr6() {
	value1, value2, value3, value4, value5, value6, ok3 := lo.TryOr6(func() (int, string, bool, foo, float64, string, error) {
		panic("my error")
	}, 21, "hello", false, foo{bar: "bar"}, 4.2, "world")

	fmt.Printf("%v %v %v %v %v %v %v\n", value1, value2, value3, value4, value5, value6, ok3)
	// Output: 21 hello false {bar} 4.2 world false
}

func ExampleTryWithErrorValue() {
	err1, ok1 := lo.TryWithErrorValue(func() error {
		return nil
	})
	err2, ok2 := lo.TryWithErrorValue(func() error {
		return fmt.Errorf("my error")
	})
	err3, ok3 := lo.TryWithErrorValue(func() error {
		panic("my error")
	})

	fmt.Printf("%v %v\n", err1, ok1)
	fmt.Printf("%v %v\n", err2, ok2)
	fmt.Printf("%v %v\n", err3, ok3)
	// Output:
	// <nil> true
	// my error false
	// my error false
}

func ExampleTryCatchWithErrorValue() {
	lo.TryCatchWithErrorValue(
		func() error {
			panic("trigger an error")
		},
		func(err any) {
			fmt.Printf("catch: %s", err)
		},
	)

	// Output: catch: trigger an error
}

type myError struct{}

func (e myError) Error() string {
	return "my error"
}

func ExampleErrorsAs() {
	doSomething := func() error {
		return &myError{}
	}

	err := doSomething()

	if rateLimitErr, ok := lo.ErrorsAs[*myError](err); ok {
		fmt.Printf("is type myError, err: %s", rateLimitErr.Error())
	} else {
		fmt.Printf("is not type myError")
	}

	// Output: is type myError, err: my error
}
