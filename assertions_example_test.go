package lo

import "fmt"

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
