package lo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAsync(t *testing.T) {
	is := assert.New(t)

	sync := make(chan struct{})

	ch := Async(func() int {
		<-sync
		return 10
	})

	sync <- struct{}{}

	select {
	case result := <-ch:
		is.Equal(result, 10)
	case <-time.After(time.Millisecond):
		is.Fail("Async should not block")
	}
}

func TestAsyncX(t *testing.T) {
	is := assert.New(t)

	{
		sync := make(chan struct{})

		ch := Async0(func() {
			<-sync
		})

		sync <- struct{}{}

		select {
		case <-ch:
		case <-time.After(time.Millisecond):
			is.Fail("Async0 should not block")
		}
	}

	{
		sync := make(chan struct{})

		ch := Async1(func() int {
			<-sync
			return 10
		})

		sync <- struct{}{}

		select {
		case result := <-ch:
			is.Equal(result, 10)
		case <-time.After(time.Millisecond):
			is.Fail("Async1 should not block")
		}
	}
}
