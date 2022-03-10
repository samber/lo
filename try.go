package lo

func Try(callback func() error) (err bool) {
	err = true

	defer func() {
		if r := recover(); r != nil {
			err = false
		}
	}()

	if callback() != nil {
		err = false
	}

	return
}

func TryWithErrorValue(callback func() error) (err bool, errorValue any) {
	err = true

	defer func() {
		if r := recover(); r != nil {
			err = false
			errorValue = r
		}
	}()

	if errorValue = callback(); errorValue != nil {
		err = false
	}

	return
}

func TryCatch(callback func() error, catch func()) {
	if !Try(callback) {
		catch()
	}
}

func TryCatchWithErrorValue(callback func() error, catch func(any)) {
	if err, val := TryWithErrorValue(callback); !err {
		catch(val)
	}
}
