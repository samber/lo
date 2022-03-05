package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	is := assert.New(t)

	r1 := Zip2[string, int]([]string{"a", "b"}, []int{1, 2})
	r2 := Zip2[string, int]([]string{"a", "b"}, []int{})

	is.Equal(r1, []Tuple2[string, int]{{A: "a", B: 1}, {A: "b", B: 2}})
	is.Equal(r2, []Tuple2[string, int]{{A: "a", B: 0}, {A: "b", B: 0}})
}

func TestUnzip(t *testing.T) {
	is := assert.New(t)

	r1, r2 := Unzip2[string, int]([]Tuple2[string, int]{{A: "a", B: 1}, {A: "b", B: 2}})

	is.Equal(r1, []string{"a", "b"})
	is.Equal(r2, []int{1, 2})
}
