//go:build go1.23

package iter_test

import (
	"fmt"
	"slices"
	"testing"

	li "github.com/samber/lo/iter"
	"github.com/stretchr/testify/assert"
)

func ExampleIndexOf() {
	numbers := slices.Values([]int{0, 1, 2, 1, 2, 3})
	fmt.Println(li.IndexOf(numbers, 1))
	// Output: 1
}

func TestIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := li.IndexOf(slices.Values([]int{0, 1, 2, 1, 2, 3}), 2)
	result2 := li.IndexOf(slices.Values([]int{0, 1, 2, 1, 2, 3}), 6)

	is.Equal(result1, 2)
	is.Equal(result2, -1)
}
