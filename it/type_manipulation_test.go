//go:build go1.23

package it

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPtr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str1 := "foo"
	str2 := "bar"
	result1 := ToPtr(values(str1, str2))

	is.Equal([]*string{&str1, &str2}, slices.Collect(result1))
}

func TestFromPtr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str1 := "foo"
	str2 := "bar"
	result1 := FromPtr(values(&str1, &str2, nil))

	is.Equal([]string{str1, str2, ""}, slices.Collect(result1))
}

func TestFromPtrOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str1 := "foo"
	str2 := "bar"
	result1 := FromPtrOr(values(&str1, &str2, nil), "fallback")

	is.Equal([]string{str1, str2, "fallback"}, slices.Collect(result1))
}

func TestToAny(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in1 := values(0, 1, 2, 3)
	in2 := values[int]()
	out1 := ToAny(in1)
	out2 := ToAny(in2)

	is.Equal([]any{0, 1, 2, 3}, slices.Collect(out1))
	is.Empty(slices.Collect(out2))
}

func TestFromAny(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	out1 := FromAny[string](values[any]("foobar", 42))
	out2 := FromAny[string](values[any]("foobar", "42"))

	is.PanicsWithValue("it.FromAny: type conversion failed", func() { _ = slices.Collect(out1) })
	is.Equal([]string{"foobar", "42"}, slices.Collect(out2))
}

func TestEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Empty(slices.Collect(Empty[string]()))
}

func TestIsEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.True(IsEmpty(values[string]()))
	is.False(IsEmpty(values("foo")))
}

func TestIsNotEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.False(IsNotEmpty(values[string]()))
	is.True(IsNotEmpty(values("foo")))
}

func TestCoalesce(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	seq0 := values[int]()
	seq1 := values(1)
	seq2 := values(1, 2)

	result1, ok1 := Coalesce[int]()
	result4, ok4 := Coalesce(seq0)
	result6, ok6 := Coalesce(seq2)
	result7, ok7 := Coalesce(seq1)
	result8, ok8 := Coalesce(seq1, seq2)
	result9, ok9 := Coalesce(seq2, seq1)
	result10, ok10 := Coalesce(seq0, seq1, seq2)

	is.NotNil(result1)
	is.Empty(slices.Collect(result1))
	is.False(ok1)

	is.NotNil(result4)
	is.Empty(slices.Collect(result4))
	is.False(ok4)

	is.NotNil(result6)
	is.Equal(slices.Collect(seq2), slices.Collect(result6))
	is.True(ok6)

	is.NotNil(result7)
	is.Equal(slices.Collect(seq1), slices.Collect(result7))
	is.True(ok7)

	is.NotNil(result8)
	is.Equal(slices.Collect(seq1), slices.Collect(result8))
	is.True(ok8)

	is.NotNil(result9)
	is.Equal(slices.Collect(seq2), slices.Collect(result9))
	is.True(ok9)

	is.NotNil(result10)
	is.Equal(slices.Collect(seq1), slices.Collect(result10))
	is.True(ok10)
}

func TestCoalesceOrEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	seq0 := values[int]()
	seq1 := values(1)
	seq2 := values(1, 2)

	result1 := CoalesceOrEmpty[int]()
	result4 := CoalesceOrEmpty(seq0)
	result6 := CoalesceOrEmpty(seq2)
	result7 := CoalesceOrEmpty(seq1)
	result8 := CoalesceOrEmpty(seq1, seq2)
	result9 := CoalesceOrEmpty(seq2, seq1)
	result10 := CoalesceOrEmpty(seq0, seq1, seq2)

	is.NotNil(result1)
	is.Empty(slices.Collect(result1))
	is.NotNil(result4)
	is.Empty(slices.Collect(result4))
	is.NotNil(result6)
	is.Equal(slices.Collect(seq2), slices.Collect(result6))
	is.NotNil(result7)
	is.Equal(slices.Collect(seq1), slices.Collect(result7))
	is.NotNil(result8)
	is.Equal(slices.Collect(seq1), slices.Collect(result8))
	is.NotNil(result9)
	is.Equal(slices.Collect(seq2), slices.Collect(result9))
	is.NotNil(result10)
	is.Equal(slices.Collect(seq1), slices.Collect(result10))
}
