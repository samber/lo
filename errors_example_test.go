package lo

import (
	"errors"
	"fmt"
)

func ExampleValidate() {
	i := 42

	err1 := Validate(i < 0, "expected %d < 0", i)
	err2 := Validate(i > 0, "expected %d > 0", i)

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
	Must(42, nil)

	// won't panic
	cb := func() (int, error) {
		return 42, nil
	}
	Must(cb())

	// will panic
	Must(42, errors.New("my error"))

	// will panic with error message
	Must(42, errors.New("world"), "hello")
}

func ExampleMust0() {
	defer func() {
		_ = recover()
	}()

	// won't panic
	Must0(nil)

	// will panic
	Must0(errors.New("my error"))

	// will panic with error message
	Must0(errors.New("world"), "hello")
}

func ExampleMust1() {
	defer func() {
		_ = recover()
	}()

	// won't panic
	Must1(42, nil)

	// won't panic
	cb := func() (int, error) {
		return 42, nil
	}
	Must1(cb())

	// will panic
	Must1(42, errors.New("my error"))

	// will panic with error message
	Must1(42, errors.New("world"), "hello")
}

func ExampleMust2() {
	defer func() {
		_ = recover()
	}()

	// won't panic
	Must2(42, "hello", nil)

	// will panic
	Must2(42, "hello", errors.New("my error"))

	// will panic with error message
	Must2(42, "hello", errors.New("world"), "hello")
}

func ExampleMust3() {
	defer func() {
		_ = recover()
	}()

	// won't panic
	Must3(42, "hello", 4.2, nil)

	// will panic
	Must3(42, "hello", 4.2, errors.New("my error"))

	// will panic with error message
	Must3(42, "hello", 4.2, errors.New("world"), "hello")
}

func ExampleMust4() {
	defer func() {
		_ = recover()
	}()

	// won't panic
	Must4(42, "hello", 4.2, true, nil)

	// will panic
	Must4(42, "hello", 4.2, true, errors.New("my error"))

	// will panic with error message
	Must4(42, "hello", 4.2, true, errors.New("world"), "hello")
}

func ExampleMust5() {
	defer func() {
		_ = recover()
	}()

	// won't panic
	Must5(42, "hello", 4.2, true, foo{}, nil)

	// will panic
	Must5(42, "hello", 4.2, true, foo{}, errors.New("my error"))

	// will panic with error message
	Must5(42, "hello", 4.2, true, foo{}, errors.New("world"), "hello")
}

func ExampleMust6() {
	defer func() {
		_ = recover()
	}()

	// won't panic
	Must5(42, "hello", 4.2, true, foo{}, "foobar", nil)

	// will panic
	Must5(42, "hello", 4.2, true, foo{}, "foobar", errors.New("my error"))

	// will panic with error message
	Must5(42, "hello", 4.2, true, foo{}, "foobar", errors.New("world"), "hello")
}

func ExampleTry() {
	ok1 := Try(func() error {
		return nil
	})
	ok2 := Try(func() error {
		return errors.New("my error")
	})
	ok3 := Try(func() error {
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
	ok1 := Try1(func() error {
		return nil
	})
	ok2 := Try1(func() error {
		return errors.New("my error")
	})
	ok3 := Try1(func() error {
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
	ok1 := Try2(func() (int, error) {
		return 42, nil
	})
	ok2 := Try2(func() (int, error) {
		return 42, errors.New("my error")
	})
	ok3 := Try2(func() (int, error) {
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
	ok1 := Try3(func() (int, string, error) {
		return 42, "foobar", nil
	})
	ok2 := Try3(func() (int, string, error) {
		return 42, "foobar", errors.New("my error")
	})
	ok3 := Try3(func() (int, string, error) {
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
	ok1 := Try4(func() (int, string, float64, error) {
		return 42, "foobar", 4.2, nil
	})
	ok2 := Try4(func() (int, string, float64, error) {
		return 42, "foobar", 4.2, errors.New("my error")
	})
	ok3 := Try4(func() (int, string, float64, error) {
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
	ok1 := Try5(func() (int, string, float64, bool, error) {
		return 42, "foobar", 4.2, true, nil
	})
	ok2 := Try5(func() (int, string, float64, bool, error) {
		return 42, "foobar", 4.2, true, errors.New("my error")
	})
	ok3 := Try5(func() (int, string, float64, bool, error) {
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
	ok1 := Try6(func() (int, string, float64, bool, foo, error) {
		return 42, "foobar", 4.2, true, foo{}, nil
	})
	ok2 := Try6(func() (int, string, float64, bool, foo, error) {
		return 42, "foobar", 4.2, true, foo{}, errors.New("my error")
	})
	ok3 := Try6(func() (int, string, float64, bool, foo, error) {
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
	value1, ok1 := TryOr(func() (int, error) {
		return 42, nil
	}, 21)
	value2, ok2 := TryOr(func() (int, error) {
		return 42, errors.New("my error")
	}, 21)
	value3, ok3 := TryOr(func() (int, error) {
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
	value1, ok1 := TryOr1(func() (int, error) {
		return 42, nil
	}, 21)
	value2, ok2 := TryOr1(func() (int, error) {
		return 42, errors.New("my error")
	}, 21)
	value3, ok3 := TryOr1(func() (int, error) {
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
	value1, value2, ok3 := TryOr2(func() (int, string, error) {
		panic("my error")
	}, 21, "hello")

	fmt.Printf("%v %v %v\n", value1, value2, ok3)
	// Output: 21 hello false
}

func ExampleTryOr3() {
	value1, value2, value3, ok3 := TryOr3(func() (int, string, bool, error) {
		panic("my error")
	}, 21, "hello", false)

	fmt.Printf("%v %v %v %v\n", value1, value2, value3, ok3)
	// Output: 21 hello false false
}

func ExampleTryOr4() {
	value1, value2, value3, value4, ok3 := TryOr4(func() (int, string, bool, foo, error) {
		panic("my error")
	}, 21, "hello", false, foo{bar: "bar"})

	fmt.Printf("%v %v %v %v %v\n", value1, value2, value3, value4, ok3)
	// Output: 21 hello false {bar} false
}

func ExampleTryOr5() {
	value1, value2, value3, value4, value5, ok3 := TryOr5(func() (int, string, bool, foo, float64, error) {
		panic("my error")
	}, 21, "hello", false, foo{bar: "bar"}, 4.2)

	fmt.Printf("%v %v %v %v %v %v\n", value1, value2, value3, value4, value5, ok3)
	// Output: 21 hello false {bar} 4.2 false
}

func ExampleTryOr6() {
	value1, value2, value3, value4, value5, value6, ok3 := TryOr6(func() (int, string, bool, foo, float64, string, error) {
		panic("my error")
	}, 21, "hello", false, foo{bar: "bar"}, 4.2, "world")

	fmt.Printf("%v %v %v %v %v %v %v\n", value1, value2, value3, value4, value5, value6, ok3)
	// Output: 21 hello false {bar} 4.2 world false
}

func ExampleTryWithErrorValue() {
	err1, ok1 := TryWithErrorValue(func() error {
		return nil
	})
	err2, ok2 := TryWithErrorValue(func() error {
		return errors.New("my error")
	})
	err3, ok3 := TryWithErrorValue(func() error {
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
	TryCatchWithErrorValue(
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

	if rateLimitErr, ok := ErrorsAs[*myError](err); ok {
		fmt.Printf("is type myError, err: %s", rateLimitErr.Error())
	} else {
		fmt.Printf("is not type myError")
	}

	// Output: is type myError, err: my error
}

func ExampleAssert() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	age := 20

	// won't panic
	Assert(age >= 18)

	// won't panic
	Assert(age >= 18, "age must be at least 18")

	// will panic
	Assert(age < 18)

	// will panic
	Assert(age < 18, "age must be less than 18")

	// Output: assertion failed
}

func ExampleAssertf() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	age := 20

	// won't panic
	Assertf(age >= 18, "age must be at least 18, got %d", age)

	// will panic
	Assertf(age < 18, "age must be less than 18, got %d", age)

	// Output: assertion failed: age must be less than 18, got 20
}
