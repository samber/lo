//go:build go1.23

package it

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSeqPtr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str1 := "foo"
	str2 := "bar"
	result1 := ToSeqPtr(values(str1, str2))

	is.Equal([]*string{&str1, &str2}, slices.Collect(result1))
}

func TestFromSeqPtr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str1 := "foo"
	str2 := "bar"
	result1 := FromSeqPtr(values(&str1, &str2, nil))

	is.Equal([]string{str1, str2, ""}, slices.Collect(result1))
}

func TestFromSeqPtrOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str1 := "foo"
	str2 := "bar"
	result1 := FromSeqPtrOr(values(&str1, &str2, nil), "fallback")

	is.Equal([]string{str1, str2, "fallback"}, slices.Collect(result1))
}

func TestToAnySeq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in1 := values(0, 1, 2, 3)
	in2 := values[int]()
	out1 := ToAnySeq(in1)
	out2 := ToAnySeq(in2)

	is.Equal([]any{0, 1, 2, 3}, slices.Collect(out1))
	is.Empty(slices.Collect(out2))
}

func TestFromAnySeq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	out1 := FromAnySeq[string](values[any]("foobar", 42))
	out2 := FromAnySeq[string](values[any]("foobar", "42"))

	is.PanicsWithValue("it.FromAnySeq: type conversion failed", func() { _ = slices.Collect(out1) })
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

func TestCoalesceSeq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	seq0 := values[int]()
	seq1 := values(1)
	seq2 := values(1, 2)

	result1, ok1 := CoalesceSeq[int]()
	result4, ok4 := CoalesceSeq(seq0)
	result6, ok6 := CoalesceSeq(seq2)
	result7, ok7 := CoalesceSeq(seq1)
	result8, ok8 := CoalesceSeq(seq1, seq2)
	result9, ok9 := CoalesceSeq(seq2, seq1)
	result10, ok10 := CoalesceSeq(seq0, seq1, seq2)

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

func TestCoalesceSeqOrEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	seq0 := values[int]()
	seq1 := values(1)
	seq2 := values(1, 2)

	result1 := CoalesceSeqOrEmpty[int]()
	result4 := CoalesceSeqOrEmpty(seq0)
	result6 := CoalesceSeqOrEmpty(seq2)
	result7 := CoalesceSeqOrEmpty(seq1)
	result8 := CoalesceSeqOrEmpty(seq1, seq2)
	result9 := CoalesceSeqOrEmpty(seq2, seq1)
	result10 := CoalesceSeqOrEmpty(seq0, seq1, seq2)

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
