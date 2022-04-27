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

func TestAwait(t *testing.T) {
	is := assert.New(t)

	aAsync := Async(func() int { return 1 })
	bAsync := Async(func() string { return "1" })

	a, b := AwaitAll2(aAsync, bAsync)
	is.Equal(1, a)
	is.Equal("1", b)

	aAsync = Async(func() int { return 1 })
	bAsync = Async(func() string { return "1" })
	cAsync := Async(func() int { return 2 })

	a, b, c := AwaitAll3(aAsync, bAsync, cAsync)
	is.Equal(1, a)
	is.Equal("1", b)
	is.Equal(2, c)

	aAsync = Async(func() int { return 1 })
	bAsync = Async(func() string { return "1" })
	cAsync = Async(func() int { return 2 })
	dAsync := Async(func() string { return "2" })

	a, b, c, d := AwaitAll4(aAsync, bAsync, cAsync, dAsync)
	is.Equal(1, a)
	is.Equal("1", b)
	is.Equal(2, c)
	is.Equal("2", d)

	aAsync = Async(func() int { return 1 })
	bAsync = Async(func() string { return "1" })
	cAsync = Async(func() int { return 2 })
	dAsync = Async(func() string { return "2" })
	eAsync := Async(func() int { return 3 })

	a, b, c, d, e := AwaitAll5(aAsync, bAsync, cAsync, dAsync, eAsync)
	is.Equal(1, a)
	is.Equal("1", b)
	is.Equal(2, c)
	is.Equal("2", d)
	is.Equal(3, e)
}
