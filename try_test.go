package lo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTry(t *testing.T) {
	is := assert.New(t)

	is.False(Try(func() error {
		panic("error")
		return nil
	}))
	is.True(Try(func() error {
		return nil
	}))
}

func TestTryWithErrorValue(t *testing.T) {
	is := assert.New(t)

	err, ok := TryWithErrorValue(func() error {
		panic("error")
		return nil
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
		return nil
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
		return nil
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
