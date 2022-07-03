package lo

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	is := assert.New(t)

	is.Equal("foo", Must("foo", nil))
	is.PanicsWithValue("something went wrong", func() {
		Must("", errors.New("something went wrong"))
	})
	is.PanicsWithValue("operation shouldn't fail: something went wrong", func() {
		Must("", errors.New("something went wrong"), "operation shouldn't fail")
	})
	is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
		Must("", errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
	})

	is.Equal(1, Must(1, true))
	is.PanicsWithValue("not ok", func() {
		Must(1, false)
	})
	is.PanicsWithValue("operation shouldn't fail", func() {
		Must(1, false, "operation shouldn't fail")
	})
	is.PanicsWithValue("operation shouldn't fail with foo", func() {
		Must(1, false, "operation shouldn't fail with %s", "foo")
	})
}

func TestMustX(t *testing.T) {
	is := assert.New(t)

	{
		is.PanicsWithValue("something went wrong", func() {
			Must0(errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			Must0(errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
		is.NotPanics(func() {
			Must0(nil)
		})
	}

	{
		val1 := Must1(1, nil)
		is.Equal(1, val1)
		is.PanicsWithValue("something went wrong", func() {
			Must1(1, errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			Must1(1, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2 := Must2(1, 2, nil)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.PanicsWithValue("something went wrong", func() {
			Must2(1, 2, errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			Must2(1, 2, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3 := Must3(1, 2, 3, nil)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.PanicsWithValue("something went wrong", func() {
			Must3(1, 2, 3, errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			Must3(1, 2, 3, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3, val4 := Must4(1, 2, 3, 4, nil)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Equal(4, val4)
		is.PanicsWithValue("something went wrong", func() {
			Must4(1, 2, 3, 4, errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			Must4(1, 2, 3, 4, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3, val4, val5 := Must5(1, 2, 3, 4, 5, nil)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Equal(4, val4)
		is.Equal(5, val5)
		is.PanicsWithValue("something went wrong", func() {
			Must5(1, 2, 3, 4, 5, errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			Must5(1, 2, 3, 4, 5, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3, val4, val5, val6 := Must6(1, 2, 3, 4, 5, 6, nil)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Equal(4, val4)
		is.Equal(5, val5)
		is.Equal(6, val6)
		is.PanicsWithValue("something went wrong", func() {
			Must6(1, 2, 3, 4, 5, 6, errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			Must6(1, 2, 3, 4, 5, 6, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
	}

	{
		is.PanicsWithValue("not ok", func() {
			Must0(false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			Must0(false, "operation shouldn't fail with %s", "foo")
		})
		is.NotPanics(func() {
			Must0(true)
		})
	}

	{
		val1 := Must1(1, true)
		is.Equal(1, val1)
		is.PanicsWithValue("not ok", func() {
			Must1(1, false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			Must1(1, false, "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2 := Must2(1, 2, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.PanicsWithValue("not ok", func() {
			Must2(1, 2, false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			Must2(1, 2, false, "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3 := Must3(1, 2, 3, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.PanicsWithValue("not ok", func() {
			Must3(1, 2, 3, false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			Must3(1, 2, 3, false, "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3, val4 := Must4(1, 2, 3, 4, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Equal(4, val4)
		is.PanicsWithValue("not ok", func() {
			Must4(1, 2, 3, 4, false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			Must4(1, 2, 3, 4, false, "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3, val4, val5 := Must5(1, 2, 3, 4, 5, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Equal(4, val4)
		is.Equal(5, val5)
		is.PanicsWithValue("not ok", func() {
			Must5(1, 2, 3, 4, 5, false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			Must5(1, 2, 3, 4, 5, false, "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3, val4, val5, val6 := Must6(1, 2, 3, 4, 5, 6, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Equal(4, val4)
		is.Equal(5, val5)
		is.Equal(6, val6)
		is.PanicsWithValue("not ok", func() {
			Must6(1, 2, 3, 4, 5, 6, false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			Must6(1, 2, 3, 4, 5, 6, false, "operation shouldn't fail with %s", "foo")
		})
	}
}

func TestTry(t *testing.T) {
	is := assert.New(t)

	is.False(Try(func() error {
		panic("error")
	}))
	is.True(Try(func() error {
		return nil
	}))
	is.False(Try(func() error {
		return fmt.Errorf("fail")
	}))
}

func TestTryX(t *testing.T) {
	is := assert.New(t)

	is.True(Try2(func() (string, error) {
		return "", nil
	}))

	is.True(Try3(func() (string, string, error) {
		return "", "", nil
	}))

	is.True(Try4(func() (string, string, string, error) {
		return "", "", "", nil
	}))

	is.True(Try5(func() (string, string, string, string, error) {
		return "", "", "", "", nil
	}))

	is.True(Try6(func() (string, string, string, string, string, error) {
		return "", "", "", "", "", nil
	}))

	is.False(Try2(func() (string, error) {
		panic("error")
	}))

	is.False(Try3(func() (string, string, error) {
		panic("error")
	}))

	is.False(Try4(func() (string, string, string, error) {
		panic("error")
	}))

	is.False(Try5(func() (string, string, string, string, error) {
		panic("error")
	}))

	is.False(Try6(func() (string, string, string, string, string, error) {
		panic("error")
	}))

	is.False(Try2(func() (string, error) {
		return "", errors.New("foo")
	}))

	is.False(Try3(func() (string, string, error) {
		return "", "", errors.New("foo")
	}))

	is.False(Try4(func() (string, string, string, error) {
		return "", "", "", errors.New("foo")
	}))

	is.False(Try5(func() (string, string, string, string, error) {
		return "", "", "", "", errors.New("foo")
	}))

	is.False(Try6(func() (string, string, string, string, string, error) {
		return "", "", "", "", "", errors.New("foo")
	}))
}

func TestTryWithErrorValue(t *testing.T) {
	is := assert.New(t)

	err, ok := TryWithErrorValue(func() error {
		panic("error")
	})
	is.False(ok)
	is.Equal("error", err)

	err, ok = TryWithErrorValue(func() error {
		return nil
	})
	is.True(ok)
	is.Equal(nil, err)
}

func TestTryCatch(t *testing.T) {
	is := assert.New(t)

	caught := false
	TryCatch(func() error {
		panic("error")
	}, func() {
		//error was caught
		caught = true
	})
	is.True(caught)

	caught = false
	TryCatch(func() error {
		return nil
	}, func() {
		//no error to be caught
		caught = true
	})
	is.False(caught)
}

func TestTryCatchWithErrorValue(t *testing.T) {
	is := assert.New(t)

	caught := false
	TryCatchWithErrorValue(func() error {
		panic("error")
	}, func(val any) {
		//error was caught
		caught = val == "error"
	})
	is.True(caught)

	caught = false
	TryCatchWithErrorValue(func() error {
		return nil
	}, func(val any) {
		//no error to be caught
		caught = true
	})
	is.False(caught)
}
