package lo

import "errors"

// Recover0 executes a callback function and recovers from any panic that occurs,
// converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It does not return any values.
//
// Returns:
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise nil.
func Recover0(callback func()) (err error) {
	return Recover0Typed[error](callback, true)
}

// Recover0Error executes a callback function that returns an error, and recovers
// from any panic that occurs, converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns an error.
//
// Returns:
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise the error returned by the callback.
func Recover0Error(callback func() error) (err error) {
	return Recover0ErrorTyped[error](callback, true)
}

// Recover1 executes a callback function that returns a single value and recovers from any panic that occurs,
// converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns a single value of type A.
//
// Returns:
//   - r1: The value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise nil.
func Recover1[A any](callback func() A) (r1 A, err error) {
	return Recover1Typed[A, error](callback, true)
}

// Recover1Error executes a callback function that returns a single value and an error,
// and recovers from any panic that occurs, converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns a single value of type A and an error.
//
// Returns:
//   - r1: The value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise the error returned by the callback.
func Recover1Error[A any](callback func() (A, error)) (r1 A, err error) {
	return Recover1ErrorTyped[A, error](callback, true)
}

// Recover2 executes a callback function that returns two values and recovers from
// any panic that occurs, converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns two values of types A and B.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise nil.
func Recover2[A, B any](callback func() (A, B)) (r1 A, r2 B, err error) {
	return Recover2Typed[A, B, error](callback, true)
}

// Recover2Error executes a callback function that returns two values and an error,
// and recovers from any panic that occurs, converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns two values of types A and B, and an error.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise the error returned by the callback.
func Recover2Error[A, B any](callback func() (A, B, error)) (r1 A, r2 B, err error) {
	return Recover2ErrorTyped[A, B, error](callback, true)
}

// Recover3 executes a callback function that returns three values and recovers from any panic that occurs,
// converting it to an error if possible. It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns three values of types A, B, and C.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise nil.
func Recover3[A, B, C any](callback func() (A, B, C)) (r1 A, r2 B, r3 C, err error) {
	return Recover3Typed[A, B, C, error](callback, true)
}

// Recover3Error executes a callback function that returns three values and an error,
// and recovers from any panic that occurs, converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns three values of types A, B, and C, and an error.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise the error returned by the callback.
func Recover3Error[A, B, C any](callback func() (A, B, C, error)) (r1 A, r2 B, r3 C, err error) {
	return Recover3ErrorTyped[A, B, C, error](callback, true)
}

// Recover4 executes a callback function that returns four values and recovers from any panic that occurs,
// converting it to an error if possible. It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns four values of types A, B, C, and D.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - r4: The fourth value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise nil.
func Recover4[A, B, C, D any](callback func() (A, B, C, D)) (r1 A, r2 B, r3 C, r4 D, err error) {
	return Recover4Typed[A, B, C, D, error](callback, true)
}

// Recover4Error executes a callback function that returns four values and an error,
// and recovers from any panic that occurs, converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns four values of types A, B, C, and D, and an error.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - r4: The fourth value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise the error returned by the callback.
func Recover4Error[A, B, C, D any](callback func() (A, B, C, D, error)) (r1 A, r2 B, r3 C, r4 D, err error) {
	return Recover4ErrorTyped[A, B, C, D, error](callback, true)
}

// Recover5 executes a callback function that returns five values and recovers from any panic that occurs,
// converting it to an error if possible. It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns five values of types A, B, C, D, and E.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - r4: The fourth value returned by the callback function.
//   - r5: The fifth value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise nil.
func Recover5[A, B, C, D, E any](callback func() (A, B, C, D, E)) (r1 A, r2 B, r3 C, r4 D, r5 E, err error) {
	return Recover5Typed[A, B, C, D, E, error](callback, true)
}

// Recover5Error executes a callback function that returns five values and an error,
// and recovers from any panic that occurs, converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns five values of types A, B, C, D, and E, and an error.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - r4: The fourth value returned by the callback function.
//   - r5: The fifth value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise the error returned by the callback.
func Recover5Error[A, B, C, D, E any](callback func() (A, B, C, D, E, error)) (r1 A, r2 B, r3 C, r4 D, r5 E, err error) {
	return Recover5ErrorTyped[A, B, C, D, E, error](callback, true)
}

// Recover6 executes a callback function that returns six values and recovers from any panic that occurs,
// converting it to an error if possible. It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns six values of types A, B, C, D, E, and F.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - r4: The fourth value returned by the callback function.
//   - r5: The fifth value returned by the callback function.
//   - r6: The sixth value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise nil.
func Recover6[A, B, C, D, E, F any](callback func() (A, B, C, D, E, F)) (r1 A, r2 B, r3 C, r4 D, r5 E, r6 F, err error) {
	return Recover6Typed[A, B, C, D, E, F, error](callback, true)
}

// Recover6Error executes a callback function that returns six values and an error,
// and recovers from any panic that occurs, converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns six values of types A, B, C, D, E, F, and an error.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - r4: The fourth value returned by the callback function.
//   - r5: The fifth value returned by the callback function.
//   - r6: The sixth value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise the error returned by the callback.
func Recover6Error[A, B, C, D, E, F any](callback func() (A, B, C, D, E, F, error)) (r1 A, r2 B, r3 C, r4 D, r5 E, r6 F, err error) {
	return Recover6ErrorTyped[A, B, C, D, E, F, error](callback, true)
}

// TypedRecover is a generic function that recovers from a panic and assigns the recovered
// error to the provided error pointer. It can optionally catch string panics and convert
// them to errors.
//
// Parameters:
//   - err: A pointer to an error where the recovered error will be stored.
//   - catchString: A variadic boolean parameter. If provided and true, string panics will
//     be caught and converted to errors.
//
// The function does not return any values. Instead, it modifies the error pointer passed
// as an argument to store the recovered error.
//
// It is intended to be used as deferred function in a function that needs to recover from
// panics.
//
// Example:
//
//	func myFunction() {
//		defer TypedRecover0[error](func() { fmt.Println("Recovered from panic") })
//		panic("This is a panic")
//	}
func TypedRecover[R error](err *error, catchString ...bool) {
	catchStr := len(catchString) > 0 && catchString[0]
	if r := recover(); r != nil {
		if e, ok := r.(R); ok {
			*err = e
		} else if e, ok := r.(string); catchStr && ok {
			*err = errors.New(e)
		} else {
			panic(r)
		}
	}
}

// Recover0Typed executes a callback function and recovers from any panic that occurs,
// converting it to an error if possible. It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It does not return any value.
//   - catchString: An optional boolean flag indicating whether to catch string panics and convert them to errors.
//
// Type Parameters:
//   - R: The type of error to recover from. It should be an error type.
//
// Returns:
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise nil.
func Recover0Typed[R error](callback func(), catchString ...bool) (err error) {
	defer TypedRecover[R](&err, catchString...)
	callback()
	return
}

// Recover0ErrorTyped executes a callback function that returns an error and recovers from any panic that occurs,
// converting it to an error if possible. It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns an error.
//   - catchString: An optional boolean flag indicating whether to catch string panics and convert them to errors.
//
// Type Parameters:
//   - R: The type of error to recover from. It should be an error type.
//
// Returns:
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise the error returned by the callback.
func Recover0ErrorTyped[R error](callback func() error, catchString ...bool) (err error) {
	defer TypedRecover[R](&err, catchString...)
	err = callback()
	return
}

// Recover1Typed executes a callback function that returns a single value and recovers from any panic that occurs,
// converting it to an error if possible. It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns a single value of type A.
//   - catchString: An optional boolean flag indicating whether to catch string panics and convert them to errors.
//
// Type Parameters:
//   - A: The type of the value returned by the callback function.
//   - R: The type of error to recover from. It should be an error type.
//
// Returns:
//   - r1: The value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise nil.
func Recover1Typed[A any, R error](callback func() A, catchString ...bool) (r1 A, err error) {
	defer TypedRecover[R](&err, catchString...)
	r1 = callback()
	return
}

// Recover1ErrorTyped executes a callback function that returns a single value and an error,
// and recovers from any panic that occurs, converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns a single value of type A and an error.
//   - catchString: An optional boolean flag indicating whether to catch string panics and convert them to errors.
//
// Type Parameters:
//   - A: The type of the value returned by the callback function.
//   - R: The type of error to recover from. It should be an error type.
//
// Returns:
//   - r1: The value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise the error returned by the callback.
func Recover1ErrorTyped[A any, R error](callback func() (A, error), catchString ...bool) (r1 A, err error) {
	defer TypedRecover[R](&err, catchString...)
	r1, err = callback()
	return
}

// Recover2Typed executes a callback function that returns two values and recovers from any panic that occurs,
// converting it to an error if possible. It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns two values of types A and B.
//   - catchString: An optional boolean flag indicating whether to catch string panics and convert them to errors.
//
// Type Parameters:
//   - A: The type of the first value returned by the callback function.
//   - B: The type of the second value returned by the callback function.
//   - R: The type of error to recover from. It should be an error type.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise nil.
func Recover2Typed[A, B any, R error](callback func() (A, B), catchString ...bool) (r1 A, r2 B, err error) {
	defer TypedRecover[R](&err, catchString...)
	r1, r2 = callback()
	return
}

// Recover2ErrorTyped executes a callback function that returns two values and an error,
// and recovers from any panic that occurs, converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns two values of types A and B, and an error.
//   - catchString: An optional boolean flag indicating whether to catch string panics and convert them to errors.
//
// Type Parameters:
//   - A: The type of the first value returned by the callback function.
//   - B: The type of the second value returned by the callback function.
//   - R: The type of error to recover from. It should be an error type.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise the error returned by the callback.
func Recover2ErrorTyped[A, B any, R error](callback func() (A, B, error), catchString ...bool) (r1 A, r2 B, err error) {
	defer TypedRecover[R](&err, catchString...)
	r1, r2, err = callback()
	return
}

// Recover3Typed executes a callback function that returns three values and recovers from any panic that occurs,
// converting it to an error if possible. It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns three values of types A, B, and C.
//   - catchString: An optional boolean flag indicating whether to catch string panics and convert them to errors.
//
// Type Parameters:
//   - A: The type of the first value returned by the callback function.
//   - B: The type of the second value returned by the callback function.
//   - C: The type of the third value returned by the callback function.
//   - R: The type of error to recover from. It should be an error type.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise nil.
func Recover3Typed[A, B, C any, R error](callback func() (A, B, C), catchString ...bool) (r1 A, r2 B, r3 C, err error) {
	defer TypedRecover[R](&err, catchString...)
	r1, r2, r3 = callback()
	return
}

// Recover3ErrorTyped executes a callback function that returns three values and an error,
// and recovers from any panic that occurs, converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns three values of types A, B, and C, and an error.
//   - catchString: An optional boolean flag indicating whether to catch string panics and convert them to errors.
//
// Type Parameters:
//   - A: The type of the first value returned by the callback function.
//   - B: The type of the second value returned by the callback function.
//   - C: The type of the third value returned by the callback function.
//   - R: The type of error to recover from. It should be an error type.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise the error returned by the callback.
func Recover3ErrorTyped[A, B, C any, R error](callback func() (A, B, C, error), catchString ...bool) (r1 A, r2 B, r3 C, err error) {
	defer TypedRecover[R](&err, catchString...)
	r1, r2, r3, err = callback()
	return
}

// Recover4Typed executes a callback function that returns four values and recovers from any panic that occurs,
// converting it to an error if possible. It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns four values of types A, B, C, and D.
//   - catchString: An optional boolean flag indicating whether to catch string panics and convert them to errors.
//
// Type Parameters:
//   - A: The type of the first value returned by the callback function.
//   - B: The type of the second value returned by the callback function.
//   - C: The type of the third value returned by the callback function.
//   - D: The type of the fourth value returned by the callback function.
//   - R: The type of error to recover from. It should be an error type.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - r4: The fourth value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise nil.
func Recover4Typed[A, B, C, D any, R error](callback func() (A, B, C, D), catchString ...bool) (r1 A, r2 B, r3 C, r4 D, err error) {
	defer TypedRecover[R](&err, catchString...)
	r1, r2, r3, r4 = callback()
	return
}

// Recover4ErrorTyped executes a callback function that returns four values and an error,
// and recovers from any panic that occurs, converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns four values of types A, B, C, and D, and an error.
//   - catchString: An optional boolean flag indicating whether to catch string panics and convert them to errors.
//
// Type Parameters:
//   - A: The type of the first value returned by the callback function.
//   - B: The type of the second value returned by the callback function.
//   - C: The type of the third value returned by the callback function.
//   - D: The type of the fourth value returned by the callback function.
//   - R: The type of error to recover from. It should be an error type.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - r4: The fourth value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise the error returned by the callback.
func Recover4ErrorTyped[A, B, C, D any, R error](callback func() (A, B, C, D, error), catchString ...bool) (r1 A, r2 B, r3 C, r4 D, err error) {
	defer TypedRecover[R](&err, catchString...)
	r1, r2, r3, r4, err = callback()
	return
}

// Recover5Typed executes a callback function that returns five values and recovers from any panic that occurs,
// converting it to an error if possible. It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns five values of types A, B, C, D, and E.
//   - catchString: An optional boolean flag indicating whether to catch string panics and convert them to errors.
//
// Type Parameters:
//   - A: The type of the first value returned by the callback function.
//   - B: The type of the second value returned by the callback function.
//   - C: The type of the third value returned by the callback function.
//   - D: The type of the fourth value returned by the callback function.
//   - E: The type of the fifth value returned by the callback function.
//   - R: The type of error to recover from. It should be an error type.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - r4: The fourth value returned by the callback function.
//   - r5: The fifth value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise nil.
func Recover5Typed[A, B, C, D, E any, R error](callback func() (A, B, C, D, E), catchString ...bool) (r1 A, r2 B, r3 C, r4 D, r5 E, err error) {
	defer TypedRecover[R](&err, catchString...)
	r1, r2, r3, r4, r5 = callback()
	return
}

// Recover5ErrorTyped executes a callback function that returns five values and an error,
// and recovers from any panic that occurs, converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns five values of types A, B, C, D, and E, and an error.
//   - catchString: An optional boolean flag indicating whether to catch string panics and convert them to errors.
//
// Type Parameters:
//   - A: The type of the first value returned by the callback function.
//   - B: The type of the second value returned by the callback function.
//   - C: The type of the third value returned by the callback function.
//   - D: The type of the fourth value returned by the callback function.
//   - E: The type of the fifth value returned by the callback function.
//   - R: The type of error to recover from. It should be an error type.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - r4: The fourth value returned by the callback function.
//   - r5: The fifth value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise the error returned by the callback.
func Recover5ErrorTyped[A, B, C, D, E any, R error](callback func() (A, B, C, D, E, error), catchString ...bool) (r1 A, r2 B, r3 C, r4 D, r5 E, err error) {
	defer TypedRecover[R](&err, catchString...)
	r1, r2, r3, r4, r5, err = callback()
	return
}

// Recover6Typed executes a callback function that returns six values and recovers from any panic that occurs,
// converting it to an error if possible. It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns six values of types A, B, C, D, E, and F.
//   - catchString: An optional boolean flag indicating whether to catch string panics and convert them to errors.
//
// Type Parameters:
//   - A: The type of the first value returned by the callback function.
//   - B: The type of the second value returned by the callback function.
//   - C: The type of the third value returned by the callback function.
//   - D: The type of the fourth value returned by the callback function.
//   - E: The type of the fifth value returned by the callback function.
//   - F: The type of the sixth value returned by the callback function.
//   - R: The type of error to recover from. It should be an error type.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - r4: The fourth value returned by the callback function.
//   - r5: The fifth value returned by the callback function.
//   - r6: The sixth value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise nil.
func Recover6Typed[A, B, C, D, E, F any, R error](callback func() (A, B, C, D, E, F), catchString ...bool) (r1 A, r2 B, r3 C, r4 D, r5 E, r6 F, err error) {
	defer TypedRecover[R](&err, catchString...)
	r1, r2, r3, r4, r5, r6 = callback()
	return
}

// Recover6ErrorTyped executes a callback function that returns six values and an error,
// and recovers from any panic that occurs, converting it to an error if possible.
// It specifically recovers from panics of type error or string.
//
// Parameters:
//   - callback: A function to be executed. It returns six values of types A, B, C, D, E, F, and an error.
//   - catchString: An optional boolean flag indicating whether to catch string panics and convert them to errors.
//
// Type Parameters:
//   - A: The type of the first value returned by the callback function.
//   - B: The type of the second value returned by the callback function.
//   - C: The type of the third value returned by the callback function.
//   - D: The type of the fourth value returned by the callback function.
//   - E: The type of the fifth value returned by the callback function.
//   - F: The type of the sixth value returned by the callback function.
//   - R: The type of error to recover from. It should be an error type.
//
// Returns:
//   - r1: The first value returned by the callback function.
//   - r2: The second value returned by the callback function.
//   - r3: The third value returned by the callback function.
//   - r4: The fourth value returned by the callback function.
//   - r5: The fifth value returned by the callback function.
//   - r6: The sixth value returned by the callback function.
//   - err: An error if a panic of type error or string occurred and was recovered, otherwise the error returned by the callback.
func Recover6ErrorTyped[A, B, C, D, E, F any, R error](callback func() (A, B, C, D, E, F, error), catchString ...bool) (r1 A, r2 B, r3 C, r4 D, r5 E, r6 F, err error) {
	defer TypedRecover[R](&err, catchString...)
	r1, r2, r3, r4, r5, r6, err = callback()
	return
}
