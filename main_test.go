package lo

import (
	"testing"

	"github.com/samber/lo/internal/xtime"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	xtime.SetClock(xtime.NewFakeClock())
	goleak.VerifyTestMain(m)
}
