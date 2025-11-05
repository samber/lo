package lo

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/samber/lo/internal/xtime"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	xtime.SetClock(xtime.NewFakeClock())
	goleak.VerifyTestMain(m)
}

// https://github.com/stretchr/testify/issues/1101
func testWithTimeout(t *testing.T, timeout time.Duration) {
	t.Helper()

	testFinished := make(chan struct{})

	t.Cleanup(func() {
		close(testFinished)
	})

	line := ""
	funcName := ""

	var pc [1]uintptr
	n := runtime.Callers(2, pc[:])
	if n > 0 {
		frames := runtime.CallersFrames(pc[:])
		frame, _ := frames.Next()
		line = frame.File + ":" + strconv.Itoa(frame.Line)
		funcName = frame.Function
	}

	go func() {
		select {
		case <-testFinished:
		case <-time.After(timeout):
			if line == "" || funcName == "" {
				panic(fmt.Sprintf("Test timed out after: %v", timeout))
			}
			panic(fmt.Sprintf("%s: Test timed out after: %v\n%s", funcName, timeout, line))

			// t.Errorf("Test timed out after: %v", timeout)
			// os.Exit(1)
		}
	}()
}

type foo struct {
	bar string
}

func (f foo) Clone() foo {
	return foo{f.bar}
}
