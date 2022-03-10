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

func TryCatch(callback func() error, catch func()) {
	if !Try(callback) {
		catch()
	}
}

func TryCatchWithErrorValue(callback func() error, catch func(any)) {
	if err, ok := TryWithErrorValue(callback); !ok {
		catch(err)
	}
}
