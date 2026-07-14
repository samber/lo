//go:build go1.23

package it

import (
	"iter"
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

	tests := []struct {
		name     string
		input    iter.Seq[int]
		expected []any
	}{
		{name: "with elements", input: values(0, 1, 2, 3), expected: []any{0, 1, 2, 3}},
		{name: "empty", input: values[int](), expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(ToAnySeq(tt.input)))
		})
	}
}

func TestFromAnySeq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name        string
		input       iter.Seq[any]
		expectPanic bool
		expected    []string
	}{
		{name: "type conversion failure panics", input: values[any]("foobar", 42), expectPanic: true},
		{name: "successful conversion", input: values[any]("foobar", "42"), expected: []string{"foobar", "42"}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			out := FromAnySeq[string](tt.input)

			if tt.expectPanic {
				is.PanicsWithValue("it.FromAnySeq: type conversion failed", func() { _ = slices.Collect(out) })
			} else {
				is.Equal(tt.expected, slices.Collect(out))
			}
		})
	}
}

func TestEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Empty(slices.Collect(Empty[string]()))
}

func TestIsEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    iter.Seq[string]
		expected bool
	}{
		{name: "empty sequence", input: values[string](), expected: true},
		{name: "non-empty sequence", input: values("foo"), expected: false},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, IsEmpty(tt.input))
		})
	}
}

func TestIsNotEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    iter.Seq[string]
		expected bool
	}{
		{name: "empty sequence", input: values[string](), expected: false},
		{name: "non-empty sequence", input: values("foo"), expected: true},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, IsNotEmpty(tt.input))
		})
	}
}

func TestCoalesceSeq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	seq0 := values[int]()
	seq1 := values(1)
	seq2 := values(1, 2)

	tests := []struct {
		name       string
		args       []iter.Seq[int]
		expected   []int
		expectedOk bool
	}{
		{name: "no sequences", args: nil, expectedOk: false},
		{name: "single empty sequence", args: []iter.Seq[int]{seq0}, expectedOk: false},
		{name: "single sequence with two elements", args: []iter.Seq[int]{seq2}, expected: []int{1, 2}, expectedOk: true},
		{name: "single sequence with one element", args: []iter.Seq[int]{seq1}, expected: []int{1}, expectedOk: true},
		{name: "first non-empty sequence wins (one then two elements)", args: []iter.Seq[int]{seq1, seq2}, expected: []int{1}, expectedOk: true},
		{name: "first non-empty sequence wins (two then one elements)", args: []iter.Seq[int]{seq2, seq1}, expected: []int{1, 2}, expectedOk: true},
		{name: "skips leading empty sequence", args: []iter.Seq[int]{seq0, seq1, seq2}, expected: []int{1}, expectedOk: true},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, ok := CoalesceSeq(tt.args...)

			is.NotNil(result)
			is.Equal(tt.expectedOk, ok)

			if tt.expectedOk {
				is.Equal(tt.expected, slices.Collect(result))
			} else {
				is.Empty(slices.Collect(result))
			}
		})
	}
}

func TestCoalesceSeqOrEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	seq0 := values[int]()
	seq1 := values(1)
	seq2 := values(1, 2)

	tests := []struct {
		name       string
		args       []iter.Seq[int]
		expected   []int
		expectHits bool
	}{
		{name: "no sequences", args: nil, expectHits: false},
		{name: "single empty sequence", args: []iter.Seq[int]{seq0}, expectHits: false},
		{name: "single sequence with two elements", args: []iter.Seq[int]{seq2}, expected: []int{1, 2}, expectHits: true},
		{name: "single sequence with one element", args: []iter.Seq[int]{seq1}, expected: []int{1}, expectHits: true},
		{name: "first non-empty sequence wins (one then two elements)", args: []iter.Seq[int]{seq1, seq2}, expected: []int{1}, expectHits: true},
		{name: "first non-empty sequence wins (two then one elements)", args: []iter.Seq[int]{seq2, seq1}, expected: []int{1, 2}, expectHits: true},
		{name: "skips leading empty sequence", args: []iter.Seq[int]{seq0, seq1, seq2}, expected: []int{1}, expectHits: true},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CoalesceSeqOrEmpty(tt.args...)

			is.NotNil(result)

			if tt.expectHits {
				is.Equal(tt.expected, slices.Collect(result))
			} else {
				is.Empty(slices.Collect(result))
			}
		})
	}
}
