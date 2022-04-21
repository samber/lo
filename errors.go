package lo

// Must is a helper that wraps a call to a function returning a value and an error
// and panics if the error is non-nil.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}

	return val
}

// Must0 has the same behavior than Must, but callback returns no variable.
func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

// Must1 is an alias to Must
func Must1[T any](val T, err error) T {
	return Must(val, err)
}

// Must2 has the same behavior than Must, but callback returns 2 variables.
func Must2[T1 any, T2 any](val1 T1, val2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}

	return val1, val2
}

// Must3 has the same behavior than Must, but callback returns 3 variables.
func Must3[T1 any, T2 any, T3 any](val1 T1, val2 T2, val3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err)
	}

	return val1, val2, val3
}

// Must4 has the same behavior than Must, but callback returns 4 variables.
func Must4[T1 any, T2 any, T3 any, T4 any](val1 T1, val2 T2, val3 T3, val4 T4, err error) (T1, T2, T3, T4) {
	if err != nil {
		panic(err)
	}

	return val1, val2, val3, val4
}

// Must5 has the same behavior than Must, but callback returns 5 variables.
func Must5[T1 any, T2 any, T3 any, T4 any, T5 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5, err error) (T1, T2, T3, T4, T5) {
	if err != nil {
		panic(err)
	}

	return val1, val2, val3, val4, val5
}

// Must6 has the same behavior than Must, but callback returns 6 variables.
func Must6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5, val6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	if err != nil {
		panic(err)
	}

	return val1, val2, val3, val4, val5, val6
}

// Try calls the function and return false in case of error.
func Try(callback func() error) (ok bool) {
	ok = true

	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()

	err := callback()
	if err != nil {
		ok = false
	}

	return
}

// Try0 has the same behavior than Try, but callback returns no variable.
func Try0[T any](callback func()) bool {
	return Try(func() error {
		callback()
		return nil
	})
}

// Try1 is an alias to Try.
func Try1[T any](callback func() error) bool {
	return Try(callback)
}

// Try2 has the same behavior than Try, but callback returns 2 variables.
func Try2[T any](callback func() (T, error)) bool {
	return Try(func() error {
		_, err := callback()
		return err
	})
}

// Try3 has the same behavior than Try, but callback returns 3 variables.
func Try3[T, R any](callback func() (T, R, error)) bool {
	return Try(func() error {
		_, _, err := callback()
		return err
	})
}

// Try4 has the same behavior than Try, but callback returns 4 variables.
func Try4[T, R, S any](callback func() (T, R, S, error)) bool {
	return Try(func() error {
		_, _, _, err := callback()
		return err
	})
}

// Try5 has the same behavior than Try, but callback returns 5 variables.
func Try5[T, R, S, Q any](callback func() (T, R, S, Q, error)) bool {
	return Try(func() error {
		_, _, _, _, err := callback()
		return err
	})
}

// Try6 has the same behavior than Try, but callback returns 6 variables.
func Try6[T, R, S, Q, U any](callback func() (T, R, S, Q, U, error)) bool {
	return Try(func() error {
		_, _, _, _, _, err := callback()
		return err
	})
}

// TryWithErrorValue has the same behavior than Try, but also returns value passed to panic.
func TryWithErrorValue(callback func() error) (errorValue any, ok bool) {
	ok = true

	defer func() {
		if r := recover(); r != nil {
			ok = false
			errorValue = r
		}
	}()

	err := callback()
	if err != nil {
		ok = false
		errorValue = err
	}

	return
}

// TryCatch has the same behavior than Try, but calls the catch function in case of error.
func TryCatch(callback func() error, catch func()) {
	if !Try(callback) {
		catch()
	}
}

// TryCatchWithErrorValue has the same behavior than TryWithErrorValue, but calls the catch function in case of error.
func TryCatchWithErrorValue(callback func() error, catch func(any)) {
	if err, ok := TryWithErrorValue(callback); !ok {
		catch(err)
	}
}

func DelayError[T any](target *error) func(v T, err error) T {
	a := func(v T, err error) T {
		if *target == nil && err != nil {
			*target = err
		}
		return v
	}
	return a
}