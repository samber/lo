//go:build goexperiment.rangefunc

package lo

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FilterI(ToIterator(1, 2, 3, 4), func(x int, _ int) bool {
		return x%2 == 0
	})

	is.Equal(r1.Slice(), []int{2, 4})

	r2 := FilterI(ToIterator("", "foo", "", "bar", ""), func(x string, _ int) bool {
		return len(x) > 0
	})

	is.Equal(r2.Slice(), []string{"foo", "bar"})
}

func TestMapI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MapI(ToIterator(1, 2, 3, 4), func(x int, _ int) string {
		return "Hello"
	})
	result2 := MapI(ToIterator(1, 2, 3, 4), func(x int, _ int) string {
		return strconv.FormatInt(int64(x), 10)
	})

	is.Equal(result1.Len(), 4)
	is.Equal(result2.Len(), 4)
	is.Equal(result1.Slice(), []string{"Hello", "Hello", "Hello", "Hello"})
	is.Equal(result2.Slice(), []string{"1", "2", "3", "4"})
}

func TestFilterMapI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FilterMapI(ToIterator(1, 2, 3, 4), func(x int, _ int) (string, bool) {
		if x%2 == 0 {
			return strconv.FormatInt(int64(x), 10), true
		}
		return "", false
	})
	r2 := FilterMapI(ToIterator("cpu", "gpu", "mouse", "keyboard"), func(x string, _ int) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", true
		}
		return "", false
	})

	is.Equal(r1.Len(), 2)
	is.Equal(r2.Len(), 2)
	is.Equal(r1.Slice(), []string{"2", "4"})
	is.Equal(r2.Slice(), []string{"xpu", "xpu"})
}
