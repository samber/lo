package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPtr(t *testing.T) {
	is := assert.New(t)

	result1 := ToPtr[[]int]([]int{1, 2})

	is.Equal(*result1, []int{1, 2})
}

func TestToSlicePtr(t *testing.T) {
	is := assert.New(t)

	str1 := "foo"
	str2 := "bar"
	result1 := ToSlicePtr[string]([]string{str1, str2})

	is.Equal(result1, []*string{&str1, &str2})
}
