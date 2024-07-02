package lo

import (
	"testing"
	"time"
)

// https://github.com/stretchr/testify/issues/1101
func testWithTimeout(t *testing.T, timeout time.Duration) {
	t.Helper()

	testFinished := make(chan struct{})
	t.Cleanup(func() { close(testFinished) })

	go func() {
		select {
		case <-testFinished:
		case <-time.After(timeout):
			t.Errorf("test timed out after %s", timeout)
			t.FailNow()
		}
	}()
}

type foo struct {
	bar string
}

func (f foo) Clone() foo {
	return foo{f.bar}
}
