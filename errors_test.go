package lo

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"runtime/debug"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		slice     []string
		wantError bool
	}{
		{name: "non-empty slice returns error", slice: []string{"a"}, wantError: true},
		{name: "empty slice returns no error", slice: []string{}, wantError: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			result := Validate(len(tt.slice) == 0, "Slice should be empty but contains %v", tt.slice)
			if tt.wantError {
				is.Error(result)
			} else {
				is.NoError(result)
			}
		})
	}
}

func TestMust(t *testing.T) { //nolint:paralleltest
	// t.Parallel()

	t.Run("string value via error", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
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
	})

	t.Run("int value via bool", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

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
	})

	t.Run("Must0 error-type validation", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		cb := func() error {
			return assert.AnError
		}
		is.PanicsWithValue("operation should fail: assert.AnError general error for testing", func() {
			Must0(cb(), "operation should fail")
		})

		is.PanicsWithValue("must: invalid err type 'int', should either be a bool or an error", func() {
			Must0(0)
		})
		is.PanicsWithValue("must: invalid err type 'string', should either be a bool or an error", func() {
			Must0("error")
		})
	})
}

func TestMustX(t *testing.T) { //nolint:paralleltest
	// t.Parallel()

	t.Run("Must0 with error", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		is.PanicsWithValue("something went wrong", func() {
			Must0(errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			Must0(errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
		is.NotPanics(func() {
			Must0(nil)
		})
	})

	t.Run("Must1 with error", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		val1 := Must1(1, nil)
		is.Equal(1, val1)
		is.PanicsWithValue("something went wrong", func() {
			Must1(1, errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			Must1(1, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
	})

	t.Run("Must2 with error", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		val1, val2 := Must2(1, 2, nil)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.PanicsWithValue("something went wrong", func() {
			Must2(1, 2, errors.New("something went wrong"))
		})
		is.PanicsWithValue("operation shouldn't fail with foo: something went wrong", func() {
			Must2(1, 2, errors.New("something went wrong"), "operation shouldn't fail with %s", "foo")
		})
	})

	t.Run("Must3 with error", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

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
	})

	t.Run("Must4 with error", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

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
	})

	t.Run("Must5 with error", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

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
	})

	t.Run("Must6 with error", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

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
	})

	t.Run("Must0 with bool", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		is.PanicsWithValue("not ok", func() {
			Must0(false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			Must0(false, "operation shouldn't fail with %s", "foo")
		})
		is.NotPanics(func() {
			Must0(true)
		})
	})

	t.Run("Must1 with bool", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		val1 := Must1(1, true)
		is.Equal(1, val1)
		is.PanicsWithValue("not ok", func() {
			Must1(1, false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			Must1(1, false, "operation shouldn't fail with %s", "foo")
		})
	})

	t.Run("Must2 with bool", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		val1, val2 := Must2(1, 2, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.PanicsWithValue("not ok", func() {
			Must2(1, 2, false)
		})
		is.PanicsWithValue("operation shouldn't fail with foo", func() {
			Must2(1, 2, false, "operation shouldn't fail with %s", "foo")
		})
	})

	t.Run("Must3 with bool", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

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
	})

	t.Run("Must4 with bool", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

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
	})

	t.Run("Must5 with bool", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

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
	})

	t.Run("Must6 with bool", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

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
	})
}

func mustCheckerWithStack(err any, messageArgs ...any) {
	if err == nil {
		return
	}

	switch e := err.(type) {
	case bool:
		if !e {
			message := messageFromMsgAndArgs(messageArgs...)
			if message == "" {
				message = "not ok"
			}

			// panic(stackErrors.New(message))
			panic(errorsJoin(errors.New(message), errors.New(string(debug.Stack()))))
		}

	case error:
		message := messageFromMsgAndArgs(messageArgs...)
		if message != "" {
			// panic(stackErrors.Wrap(e, message))
			panic(errorsJoin(e, errors.New(message), errors.New(string(debug.Stack()))))
		}
		// panic(stackErrors.WithStack(e))
		panic(errorsJoin(e, errors.New(string(debug.Stack()))))

	default:
		// panic(stackErrors.New("must: invalid err type '" + reflect.TypeOf(err).Name() + "', should either be a bool or an error"))
		panic(errorsJoin(errors.New("must: invalid err type '"+reflect.TypeOf(err).Name()+"', should either be a bool or an error"),
			errors.New(string(debug.Stack()))))
	}
}

// errorsJoin: var errorsJoin = errors.Join // only go 1.20+, not in go 1.18 .
func errorsJoin(es ...error) joinErrors { return joinErrors(es) }

type joinErrors []error

func (es joinErrors) Is(target error) bool {
	for _, e := range es {
		if errors.Is(e, target) {
			return true
		}
	}
	return error(es) == target
}

func (es joinErrors) Error() string {
	sb := strings.Builder{}
	for _, e := range es {
		sb.WriteString(e.Error())
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (es joinErrors) As(t any) bool {
	for _, e := range es {
		if errors.As(e, t) {
			return true
		}
	}
	return false
}

func TestMust_userCustomHandler(t *testing.T) { //nolint:paralleltest
	oldMustChecker := MustChecker
	MustChecker = mustCheckerWithStack
	defer func() {
		MustChecker = oldMustChecker
	}()

	t.Run("wrap stack", func(t *testing.T) { //nolint:paralleltest
		err, ok := TryWithErrorValue(func() error {
			Must("foo", errors.New("wrap callstack"))
			return nil
		})
		assert.False(t, ok)
		fullErrStr := fmt.Sprintf("%+v", err)
		assert.Contains(t, fullErrStr, "/errors_test.go:", fullErrStr)
	})
	t.Run("wrap as", func(t *testing.T) { //nolint:paralleltest
		e, ok := TryWithErrorValue(func() error {
			Must("foo", errorsJoin(io.EOF, &url.Error{
				Op:  "test op",
				URL: "test url",
				Err: io.ErrUnexpectedEOF,
			}))
			return nil
		})
		assert.False(t, ok)
		err, ok := e.(error)
		assert.True(t, ok)
		errURL, ok := ErrorsAs[*url.Error](err)
		assert.True(t, ok)
		assert.NotNil(t, errURL)
		if errURL != nil {
			assert.Equal(t, "test url", errURL.URL)
			assert.Equal(t, "test op", errURL.Op)
			assert.ErrorIs(t, err, io.EOF)
			assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
		}
	})
}

func TestTry(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		fn       func() error
		expected bool
	}{
		{name: "panics", fn: func() error { panic("error") }, expected: false},
		{name: "returns nil", fn: func() error { return nil }, expected: true},
		{name: "returns error", fn: func() error { return errors.New("fail") }, expected: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			is.Equal(tt.expected, Try(tt.fn))
		})
	}
}

func TestTryX(t *testing.T) {
	t.Parallel()

	t.Run("Try1", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		is.True(Try1(func() error { return nil }))
		is.False(Try1(func() error { panic("error") }))
		is.False(Try1(func() error { return errors.New("foo") }))
	})

	t.Run("Try2", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		is.True(Try2(func() (string, error) { return "", nil }))
		is.False(Try2(func() (string, error) { panic("error") }))
		is.False(Try2(func() (string, error) { return "", errors.New("foo") }))
	})

	t.Run("Try3", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		is.True(Try3(func() (string, string, error) { return "", "", nil }))
		is.False(Try3(func() (string, string, error) { panic("error") }))
		is.False(Try3(func() (string, string, error) { return "", "", errors.New("foo") }))
	})

	t.Run("Try4", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		is.True(Try4(func() (string, string, string, error) { return "", "", "", nil }))
		is.False(Try4(func() (string, string, string, error) { panic("error") }))
		is.False(Try4(func() (string, string, string, error) { return "", "", "", errors.New("foo") }))
	})

	t.Run("Try5", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		is.True(Try5(func() (string, string, string, string, error) { return "", "", "", "", nil }))
		is.False(Try5(func() (string, string, string, string, error) { panic("error") }))
		is.False(Try5(func() (string, string, string, string, error) { return "", "", "", "", errors.New("foo") }))
	})

	t.Run("Try6", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		is.True(Try6(func() (string, string, string, string, string, error) { return "", "", "", "", "", nil }))
		is.False(Try6(func() (string, string, string, string, string, error) { panic("error") }))
		is.False(Try6(func() (string, string, string, string, string, error) { return "", "", "", "", "", errors.New("foo") }))
	})
}

func TestTryOr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		fn         func() (int, error)
		fallback   int
		expected   int
		expectedOk bool
	}{
		{name: "panics, returns fallback", fn: func() (int, error) { panic("error") }, fallback: 42, expected: 42, expectedOk: false},
		{name: "returns error, returns fallback", fn: func() (int, error) { return 21, assert.AnError }, fallback: 42, expected: 42, expectedOk: false},
		{name: "succeeds, returns value", fn: func() (int, error) { return 21, nil }, fallback: 42, expected: 21, expectedOk: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result, ok := TryOr(tt.fn, tt.fallback)
			is.Equal(tt.expected, result)
			is.Equal(tt.expectedOk, ok)
		})
	}
}

func TestTryOrX(t *testing.T) {
	t.Parallel()

	t.Run("TryOr1", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		a1, ok1 := TryOr1(func() (int, error) { panic("error") }, 42)
		a2, ok2 := TryOr1(func() (int, error) { return 21, assert.AnError }, 42)
		a3, ok3 := TryOr1(func() (int, error) { return 21, nil }, 42)

		is.Equal(42, a1)
		is.False(ok1)

		is.Equal(42, a2)
		is.False(ok2)

		is.Equal(21, a3)
		is.True(ok3)
	})

	t.Run("TryOr2", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		a1, b1, ok1 := TryOr2(func() (int, string, error) { panic("error") }, 42, "hello")
		a2, b2, ok2 := TryOr2(func() (int, string, error) { return 21, "world", assert.AnError }, 42, "hello")
		a3, b3, ok3 := TryOr2(func() (int, string, error) { return 21, "world", nil }, 42, "hello")

		is.Equal(42, a1)
		is.Equal("hello", b1)
		is.False(ok1)

		is.Equal(42, a2)
		is.Equal("hello", b2)
		is.False(ok2)

		is.Equal(21, a3)
		is.Equal("world", b3)
		is.True(ok3)
	})

	t.Run("TryOr3", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		a1, b1, c1, ok1 := TryOr3(func() (int, string, bool, error) { panic("error") }, 42, "hello", false)
		a2, b2, c2, ok2 := TryOr3(func() (int, string, bool, error) { return 21, "world", true, assert.AnError }, 42, "hello", false)
		a3, b3, c3, ok3 := TryOr3(func() (int, string, bool, error) { return 21, "world", true, nil }, 42, "hello", false)

		is.Equal(42, a1)
		is.Equal("hello", b1)
		is.False(c1)
		is.False(ok1)

		is.Equal(42, a2)
		is.Equal("hello", b2)
		is.False(c2)
		is.False(ok2)

		is.Equal(21, a3)
		is.Equal("world", b3)
		is.True(c3)
		is.True(ok3)
	})

	t.Run("TryOr4", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		a1, b1, c1, d1, ok1 := TryOr4(func() (int, string, bool, int, error) { panic("error") }, 42, "hello", false, 42)
		a2, b2, c2, d2, ok2 := TryOr4(func() (int, string, bool, int, error) { return 21, "world", true, 21, assert.AnError }, 42, "hello", false, 42)
		a3, b3, c3, d3, ok3 := TryOr4(func() (int, string, bool, int, error) { return 21, "world", true, 21, nil }, 42, "hello", false, 42)

		is.Equal(42, a1)
		is.Equal("hello", b1)
		is.False(c1)
		is.Equal(42, d1)
		is.False(ok1)

		is.Equal(42, a2)
		is.Equal("hello", b2)
		is.False(c2)
		is.Equal(42, d2)
		is.False(ok2)

		is.Equal(21, a3)
		is.Equal("world", b3)
		is.True(c3)
		is.Equal(21, d3)
		is.True(ok3)
	})

	t.Run("TryOr5", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		a1, b1, c1, d1, e1, ok1 := TryOr5(func() (int, string, bool, int, int, error) { panic("error") }, 42, "hello", false, 42, 42)
		a2, b2, c2, d2, e2, ok2 := TryOr5(func() (int, string, bool, int, int, error) { return 21, "world", true, 21, 21, assert.AnError }, 42, "hello", false, 42, 42)
		a3, b3, c3, d3, e3, ok3 := TryOr5(func() (int, string, bool, int, int, error) { return 21, "world", true, 21, 21, nil }, 42, "hello", false, 42, 42)

		is.Equal(42, a1)
		is.Equal("hello", b1)
		is.False(c1)
		is.Equal(42, d1)
		is.Equal(42, e1)
		is.False(ok1)

		is.Equal(42, a2)
		is.Equal("hello", b2)
		is.False(c2)
		is.Equal(42, d2)
		is.Equal(42, e2)
		is.False(ok2)

		is.Equal(21, a3)
		is.Equal("world", b3)
		is.True(c3)
		is.Equal(21, d3)
		is.Equal(21, e3)
		is.True(ok3)
	})

	t.Run("TryOr6", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		a1, b1, c1, d1, e1, f1, ok1 := TryOr6(func() (int, string, bool, int, int, int, error) { panic("error") }, 42, "hello", false, 42, 42, 42)
		a2, b2, c2, d2, e2, f2, ok2 := TryOr6(func() (int, string, bool, int, int, int, error) { return 21, "world", true, 21, 21, 21, assert.AnError }, 42, "hello", false, 42, 42, 42)
		a3, b3, c3, d3, e3, f3, ok3 := TryOr6(func() (int, string, bool, int, int, int, error) { return 21, "world", true, 21, 21, 21, nil }, 42, "hello", false, 42, 42, 42)

		is.Equal(42, a1)
		is.Equal("hello", b1)
		is.False(c1)
		is.Equal(42, d1)
		is.Equal(42, e1)
		is.Equal(42, f1)
		is.False(ok1)

		is.Equal(42, a2)
		is.Equal("hello", b2)
		is.False(c2)
		is.Equal(42, d2)
		is.Equal(42, e2)
		is.Equal(42, f2)
		is.False(ok2)

		is.Equal(21, a3)
		is.Equal("world", b3)
		is.True(c3)
		is.Equal(21, d3)
		is.Equal(21, e3)
		is.Equal(21, f3)
		is.True(ok3)
	})
}

func TestTryWithErrorValue(t *testing.T) {
	t.Parallel()

	t.Run("panics with string value", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		err, ok := TryWithErrorValue(func() error {
			// getting error in case of panic, using recover function
			panic("error")
		})
		is.False(ok)
		is.Equal("error", err)
	})

	t.Run("returns wrapped error", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		err, ok := TryWithErrorValue(func() error {
			return errors.New("foo")
		})
		is.False(ok)
		e, isError := err.(error)
		is.True(isError)
		is.EqualError(e, "foo")
	})

	t.Run("succeeds", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		err, ok := TryWithErrorValue(func() error {
			return nil
		})
		is.True(ok)
		is.Nil(err)
	})
}

func TestTryCatch(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		fn       func() error
		expected bool
	}{
		{name: "panics", fn: func() error { panic("error") }, expected: true},
		{name: "no panic", fn: func() error { return nil }, expected: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			caught := false
			TryCatch(tt.fn, func() {
				caught = true
			})
			is.Equal(tt.expected, caught)
		})
	}
}

func TestTryCatchWithErrorValue(t *testing.T) {
	t.Parallel()

	t.Run("panics with string value", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		caught := false
		TryCatchWithErrorValue(func() error {
			panic("error")
		}, func(val any) {
			// error was caught
			caught = val == "error"
		})
		is.True(caught)
	})

	t.Run("no panic, callback not invoked", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		caught := false
		TryCatchWithErrorValue(func() error {
			return nil
		}, func(val any) {
			// no error to be caught
			caught = true
		})
		is.False(caught)
	})
}

type internalError struct {
	foobar string
}

func (e *internalError) Error() string {
	return "internal error"
}

func TestErrorsAs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       error
		expectedErr *internalError
		expectedOk  bool
	}{
		{name: "wrong error type", input: errors.New("hello world"), expectedErr: nil, expectedOk: false},
		{name: "matching error type", input: &internalError{foobar: "foobar"}, expectedErr: &internalError{foobar: "foobar"}, expectedOk: true},
		{name: "nil error", input: nil, expectedErr: nil, expectedOk: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			err, ok := ErrorsAs[*internalError](tt.input)
			is.Equal(tt.expectedOk, ok)
			is.Equal(tt.expectedErr, err)
		})
	}
}

func TestAssert(t *testing.T) { //nolint:paralleltest
	// t.Parallel()

	t.Run("does not panic when true", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		is.NotPanics(func() {
			Assert(true)
		})
		is.NotPanics(func() {
			Assert(true, "user defined message")
		})
	})

	t.Run("panics when false", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		is.PanicsWithValue("assertion failed", func() {
			Assert(false)
		})
		is.PanicsWithValue("assertion failed: user defined message", func() {
			Assert(false, "user defined message")
		})
	})

	// checks that the examples in `README.md` compile
	t.Run("README example compiles", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		age := 20
		is.NotPanics(func() {
			Assert(age >= 15)
		})
		is.NotPanics(func() {
			Assert(age >= 15, "user age must be >= 15")
		})
	})
}

func TestAssertf(t *testing.T) { //nolint:paralleltest
	// t.Parallel()

	t.Run("does not panic when true", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		is.NotPanics(func() {
			Assertf(true, "user defined message")
		})
		is.NotPanics(func() {
			Assertf(true, "user defined message %d %d", 1, 2)
		})
	})

	t.Run("panics when false", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		is.PanicsWithValue("assertion failed: user defined message", func() {
			Assertf(false, "user defined message")
		})
		is.PanicsWithValue("assertion failed: user defined message 1 2", func() {
			Assertf(false, "user defined message %d %d", 1, 2)
		})
	})

	// checks that the example in `README.md` compiles
	t.Run("README example compiles", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		age := 7
		is.PanicsWithValue("assertion failed: user age must be >= 15, got 7", func() {
			Assertf(age >= 15, "user age must be >= 15, got %d", age)
		})
	})
}

func TestAssertfWithCustom(t *testing.T) { //nolint:paralleltest
	oldAssertf := Assertf
	Assertf = func(condition bool, format string, args ...any) {
		if !condition {
			panic(fmt.Errorf("%s: %s", "customErr", fmt.Sprintf(format, args...)))
		}
	}
	defer func() {
		Assertf = oldAssertf
	}()

	e, ok := TryWithErrorValue(func() error {
		Assertf(false, "user defined message")
		return nil
	})
	assert.False(t, ok)
	assert.NotNil(t, e)
	err, ok := e.(error)
	assert.True(t, ok)
	assert.Equal(t, "customErr: user defined message", err.Error())
}
