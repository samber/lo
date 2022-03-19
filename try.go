package lo

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

func Try2[T any](callback func() (T, error)) bool {
	return Try(func() error {
		_, err := callback()
		return err
	})
}

func Try3[T, R any](callback func() (T, R, error)) bool {
	return Try(func() error {
		_, _, err := callback()
		return err
	})
}

func Try4[T, R, S any](callback func() (T, R, S, error)) bool {
	return Try(func() error {
		_, _, _, err := callback()
		return err
	})
}

func Try5[T, R, S, Q any](callback func() (T, R, S, Q, error)) bool {
	return Try(func() error {
		_, _, _, _, err := callback()
		return err
	})
}

func Try6[T, R, S, Q, U any](callback func() (T, R, S, Q, U, error)) bool {
	return Try(func() error {
		_, _, _, _, _, err := callback()
		return err
	})
}

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

//TODO: Overloads for TryWithErrorValue

func TryCatch(callback func() error, catch func()) {
	if !Try(callback) {
		catch()
	}
}

//TODO: Overloads for TryCatch

func TryCatchWithErrorValue(callback func() error, catch func(any)) {
	if err, ok := TryWithErrorValue(callback); !ok {
		catch(err)
	}
}

//TODO: Overloads for TryCatchWithErrorValue
