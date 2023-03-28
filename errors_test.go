package lo_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	is := assert.New(t)

	slice := []string{"a"}
	result1 := lo.Validate(len(slice) == 0, "Slice should be empty but contains %v", slice)

	slice = []string{}
	result2 := lo.Validate(len(slice) == 0, "Slice should be empty but contains %v", slice)

	is.Error(result1)
	is.NoError(result2)
}

func TestMust(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal("foo", lo.Must("foo", nil))
	is.PanicsWithValue("something went wrong", func() {
		lo.Must("", errors.New("something went wrong"))
	})
	is.PanicsWithValue("operation shouldn't fail: something went wrong", func() {
		lo.Must("", errors.New("something went wrong"), "operation shouldn't fail")
	})
	is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
		lo.Must("", errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
	})

	is.Equal(1, lo.Must(1, true))
	is.PanicsWithValue("not ok", func() {
		lo.Must(1, false)
	})
	is.PanicsWithValue("operation shouldn't fail", func() {
		lo.Must(1, false, "operation shouldn't fail")
	})
	is.PanicsWithValue("operation shouldn't fail with foo", func() {
		lo.Must(1, false, "operation shouldn't fail with %s", "foo")
	})

	cb := func() error {
		return assert.AnError
	}
	is.PanicsWithValue("operation should fail: assert.AnError general error for testing", func() {
		lo.Must0(cb(), "operation should fail")
	})

	is.PanicsWithValue("must: invalid err type 'int', should either be a bool or an error", func() {
		lo.Must0(0)
	})
	is.PanicsWithValue("must: invalid err type 'string', should either be a bool or an error", func() {
		lo.Must0("error")
	})
}

func TestMustX(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	{
		is.PanicsWithValue("something went wrong", func() {
			lo.Must0(errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			lo.Must0(errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
		is.NotPanics(func() {
			lo.Must0(nil)
		})
	}

	{
		val1 := lo.Must1(1, nil)
		is.Equal(1, val1)
		is.PanicsWithValue("something went wrong", func() {
			lo.Must1(1, errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			lo.Must1(1, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2 := lo.Must2(1, 2, nil)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.PanicsWithValue("something went wrong", func() {
			lo.Must2(1, 2, errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			lo.Must2(1, 2, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3 := lo.Must3(1, 2, 3, nil)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.PanicsWithValue("something went wrong", func() {
			lo.Must3(1, 2, 3, errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			lo.Must3(1, 2, 3, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3, val4 := lo.Must4(1, 2, 3, 4, nil)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Equal(4, val4)
		is.PanicsWithValue("something went wrong", func() {
			lo.Must4(1, 2, 3, 4, errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			lo.Must4(1, 2, 3, 4, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3, val4, val5 := lo.Must5(1, 2, 3, 4, 5, nil)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Equal(4, val4)
		is.Equal(5, val5)
		is.PanicsWithValue("something went wrong", func() {
			lo.Must5(1, 2, 3, 4, 5, errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			lo.Must5(1, 2, 3, 4, 5, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3, val4, val5, val6 := lo.Must6(1, 2, 3, 4, 5, 6, nil)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Equal(4, val4)
		is.Equal(5, val5)
		is.Equal(6, val6)
		is.PanicsWithValue("something went wrong", func() {
			lo.Must6(1, 2, 3, 4, 5, 6, errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			lo.Must6(1, 2, 3, 4, 5, 6, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
	}

	{
		is.PanicsWithValue("not ok", func() {
			lo.Must0(false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			lo.Must0(false, "operation shouldn't fail with %s", "foo")
		})
		is.NotPanics(func() {
			lo.Must0(true)
		})
	}

	{
		val1 := lo.Must1(1, true)
		is.Equal(1, val1)
		is.PanicsWithValue("not ok", func() {
			lo.Must1(1, false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			lo.Must1(1, false, "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2 := lo.Must2(1, 2, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.PanicsWithValue("not ok", func() {
			lo.Must2(1, 2, false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			lo.Must2(1, 2, false, "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3 := lo.Must3(1, 2, 3, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.PanicsWithValue("not ok", func() {
			lo.Must3(1, 2, 3, false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			lo.Must3(1, 2, 3, false, "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3, val4 := lo.Must4(1, 2, 3, 4, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Equal(4, val4)
		is.PanicsWithValue("not ok", func() {
			lo.Must4(1, 2, 3, 4, false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			lo.Must4(1, 2, 3, 4, false, "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3, val4, val5 := lo.Must5(1, 2, 3, 4, 5, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Equal(4, val4)
		is.Equal(5, val5)
		is.PanicsWithValue("not ok", func() {
			lo.Must5(1, 2, 3, 4, 5, false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			lo.Must5(1, 2, 3, 4, 5, false, "operation shouldn't fail with %s", "foo")
		})
	}

	{
		val1, val2, val3, val4, val5, val6 := lo.Must6(1, 2, 3, 4, 5, 6, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Equal(4, val4)
		is.Equal(5, val5)
		is.Equal(6, val6)
		is.PanicsWithValue("not ok", func() {
			lo.Must6(1, 2, 3, 4, 5, 6, false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			lo.Must6(1, 2, 3, 4, 5, 6, false, "operation shouldn't fail with %s", "foo")
		})
	}
}

func TestTry(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.False(lo.Try(func() error {
		panic("error")
	}))
	is.True(lo.Try(func() error {
		return nil
	}))
	is.False(lo.Try(func() error {
		return fmt.Errorf("fail")
	}))
}

func TestTryX(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.True(lo.Try1(func() error {
		return nil
	}))

	is.True(lo.Try2(func() (string, error) {
		return "", nil
	}))

	is.True(lo.Try3(func() (string, string, error) {
		return "", "", nil
	}))

	is.True(lo.Try4(func() (string, string, string, error) {
		return "", "", "", nil
	}))

	is.True(lo.Try5(func() (string, string, string, string, error) {
		return "", "", "", "", nil
	}))

	is.True(lo.Try6(func() (string, string, string, string, string, error) {
		return "", "", "", "", "", nil
	}))

	is.False(lo.Try1(func() error {
		panic("error")
	}))

	is.False(lo.Try2(func() (string, error) {
		panic("error")
	}))

	is.False(lo.Try3(func() (string, string, error) {
		panic("error")
	}))

	is.False(lo.Try4(func() (string, string, string, error) {
		panic("error")
	}))

	is.False(lo.Try5(func() (string, string, string, string, error) {
		panic("error")
	}))

	is.False(lo.Try6(func() (string, string, string, string, string, error) {
		panic("error")
	}))

	is.False(lo.Try1(func() error {
		return errors.New("foo")
	}))

	is.False(lo.Try2(func() (string, error) {
		return "", errors.New("foo")
	}))

	is.False(lo.Try3(func() (string, string, error) {
		return "", "", errors.New("foo")
	}))

	is.False(lo.Try4(func() (string, string, string, error) {
		return "", "", "", errors.New("foo")
	}))

	is.False(lo.Try5(func() (string, string, string, string, error) {
		return "", "", "", "", errors.New("foo")
	}))

	is.False(lo.Try6(func() (string, string, string, string, string, error) {
		return "", "", "", "", "", errors.New("foo")
	}))
}

func TestTryOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	a1, ok1 := lo.TryOr(func() (int, error) { panic("error") }, 42)
	a2, ok2 := lo.TryOr(func() (int, error) { return 21, assert.AnError }, 42)
	a3, ok3 := lo.TryOr(func() (int, error) { return 21, nil }, 42)

	is.Equal(42, a1)
	is.False(ok1)

	is.Equal(42, a2)
	is.False(ok2)

	is.Equal(21, a3)
	is.True(ok3)
}

func TestTryOrX(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	{
		a1, ok1 := lo.TryOr1(func() (int, error) { panic("error") }, 42)
		a2, ok2 := lo.TryOr1(func() (int, error) { return 21, assert.AnError }, 42)
		a3, ok3 := lo.TryOr1(func() (int, error) { return 21, nil }, 42)

		is.Equal(42, a1)
		is.False(ok1)

		is.Equal(42, a2)
		is.False(ok2)

		is.Equal(21, a3)
		is.True(ok3)
	}

	{
		a1, b1, ok1 := lo.TryOr2(func() (int, string, error) { panic("error") }, 42, "hello")
		a2, b2, ok2 := lo.TryOr2(func() (int, string, error) { return 21, "world", assert.AnError }, 42, "hello")
		a3, b3, ok3 := lo.TryOr2(func() (int, string, error) { return 21, "world", nil }, 42, "hello")

		is.Equal(42, a1)
		is.Equal("hello", b1)
		is.False(ok1)

		is.Equal(42, a2)
		is.Equal("hello", b2)
		is.False(ok2)

		is.Equal(21, a3)
		is.Equal("world", b3)
		is.True(ok3)
	}

	{
		a1, b1, c1, ok1 := lo.TryOr3(func() (int, string, bool, error) { panic("error") }, 42, "hello", false)
		a2, b2, c2, ok2 := lo.TryOr3(func() (int, string, bool, error) { return 21, "world", true, assert.AnError }, 42, "hello", false)
		a3, b3, c3, ok3 := lo.TryOr3(func() (int, string, bool, error) { return 21, "world", true, nil }, 42, "hello", false)

		is.Equal(42, a1)
		is.Equal("hello", b1)
		is.Equal(false, c1)
		is.False(ok1)

		is.Equal(42, a2)
		is.Equal("hello", b2)
		is.Equal(false, c2)
		is.False(ok2)

		is.Equal(21, a3)
		is.Equal("world", b3)
		is.Equal(true, c3)
		is.True(ok3)
	}

	{
		a1, b1, c1, d1, ok1 := lo.TryOr4(func() (int, string, bool, int, error) { panic("error") }, 42, "hello", false, 42)
		a2, b2, c2, d2, ok2 := lo.TryOr4(func() (int, string, bool, int, error) { return 21, "world", true, 21, assert.AnError }, 42, "hello", false, 42)
		a3, b3, c3, d3, ok3 := lo.TryOr4(func() (int, string, bool, int, error) { return 21, "world", true, 21, nil }, 42, "hello", false, 42)

		is.Equal(42, a1)
		is.Equal("hello", b1)
		is.Equal(false, c1)
		is.Equal(42, d1)
		is.False(ok1)

		is.Equal(42, a2)
		is.Equal("hello", b2)
		is.Equal(false, c2)
		is.Equal(42, d2)
		is.False(ok2)

		is.Equal(21, a3)
		is.Equal("world", b3)
		is.Equal(true, c3)
		is.Equal(21, d3)
		is.True(ok3)
	}

	{
		a1, b1, c1, d1, e1, ok1 := lo.TryOr5(func() (int, string, bool, int, int, error) { panic("error") }, 42, "hello", false, 42, 42)
		a2, b2, c2, d2, e2, ok2 := lo.TryOr5(func() (int, string, bool, int, int, error) { return 21, "world", true, 21, 21, assert.AnError }, 42, "hello", false, 42, 42)
		a3, b3, c3, d3, e3, ok3 := lo.TryOr5(func() (int, string, bool, int, int, error) { return 21, "world", true, 21, 21, nil }, 42, "hello", false, 42, 42)

		is.Equal(42, a1)
		is.Equal("hello", b1)
		is.Equal(false, c1)
		is.Equal(42, d1)
		is.Equal(42, e1)
		is.False(ok1)

		is.Equal(42, a2)
		is.Equal("hello", b2)
		is.Equal(false, c2)
		is.Equal(42, d2)
		is.Equal(42, e2)
		is.False(ok2)

		is.Equal(21, a3)
		is.Equal("world", b3)
		is.Equal(true, c3)
		is.Equal(21, d3)
		is.Equal(21, e3)
		is.True(ok3)
	}

	{
		a1, b1, c1, d1, e1, f1, ok1 := lo.TryOr6(func() (int, string, bool, int, int, int, error) { panic("error") }, 42, "hello", false, 42, 42, 42)
		a2, b2, c2, d2, e2, f2, ok2 := lo.TryOr6(func() (int, string, bool, int, int, int, error) { return 21, "world", true, 21, 21, 21, assert.AnError }, 42, "hello", false, 42, 42, 42)
		a3, b3, c3, d3, e3, f3, ok3 := lo.TryOr6(func() (int, string, bool, int, int, int, error) { return 21, "world", true, 21, 21, 21, nil }, 42, "hello", false, 42, 42, 42)

		is.Equal(42, a1)
		is.Equal("hello", b1)
		is.Equal(false, c1)
		is.Equal(42, d1)
		is.Equal(42, e1)
		is.Equal(42, f1)
		is.False(ok1)

		is.Equal(42, a2)
		is.Equal("hello", b2)
		is.Equal(false, c2)
		is.Equal(42, d2)
		is.Equal(42, e2)
		is.Equal(42, f2)
		is.False(ok2)

		is.Equal(21, a3)
		is.Equal("world", b3)
		is.Equal(true, c3)
		is.Equal(21, d3)
		is.Equal(21, e3)
		is.Equal(21, f3)
		is.True(ok3)
	}
}

func TestTryWithErrorValue(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	err, ok := lo.TryWithErrorValue(func() error {
		// getting error in case of panic, using recover function
		panic("error")
	})
	is.False(ok)
	is.Equal("error", err)

	err, ok = lo.TryWithErrorValue(func() error {
		return errors.New("foo")
	})
	is.False(ok)
	is.EqualError(err.(error), "foo")

	err, ok = lo.TryWithErrorValue(func() error {
		return nil
	})
	is.True(ok)
	is.Equal(nil, err)
}

func TestTryCatch(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	caught := false
	lo.TryCatch(func() error {
		panic("error")
	}, func() {
		// error was caught
		caught = true
	})
	is.True(caught)

	caught = false
	lo.TryCatch(func() error {
		return nil
	}, func() {
		// no error to be caught
		caught = true
	})
	is.False(caught)
}

func TestTryCatchWithErrorValue(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	caught := false
	lo.TryCatchWithErrorValue(func() error {
		panic("error")
	}, func(val any) {
		// error was caught
		caught = val == "error"
	})
	is.True(caught)

	caught = false
	lo.TryCatchWithErrorValue(func() error {
		return nil
	}, func(val any) {
		// no error to be caught
		caught = true
	})
	is.False(caught)
}

type internalError struct {
	foobar string
}

func (e *internalError) Error() string {
	return "internal error"
}

func TestErrorsAs(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	err, ok := lo.ErrorsAs[*internalError](fmt.Errorf("hello world"))
	is.False(ok)
	is.Nil(nil, err)

	err, ok = lo.ErrorsAs[*internalError](&internalError{foobar: "foobar"})
	is.True(ok)
	is.Equal(&internalError{foobar: "foobar"}, err)

	err, ok = lo.ErrorsAs[*internalError](nil)
	is.False(ok)
	is.Nil(nil, err)
}
