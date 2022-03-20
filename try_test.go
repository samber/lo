package lo

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
	is.False(Try(func() error {
		return fmt.Errorf("fail")
	}))
}

func TestTryFunctions(t *testing.T) {
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
		return "", nil
	}))

	is.False(Try3(func() (string, string, error) {
		panic("error")
		return "", "", nil
	}))

	is.False(Try4(func() (string, string, string, error) {
		panic("error")
		return "", "", "", nil
	}))

	is.False(Try5(func() (string, string, string, string, error) {
		panic("error")
		return "", "", "", "", nil
	}))

	is.False(Try6(func() (string, string, string, string, string, error) {
		panic("error")
		return "", "", "", "", "", nil
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
