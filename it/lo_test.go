//go:build go1.23

package it

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// assertSeqSupportBreak checks whether it is possible to break iteration over a [iter.Seq]
func assertSeqSupportBreak[T any](t *testing.T, seq iter.Seq[T]) iter.Seq[T] {
	assert.NotPanics(t, func() {
		for range seq {
			break
		}

		for range seq {
			return
		}
	})

	return seq
}

func values[T any](v ...T) iter.Seq[T] { return slices.Values(v) }

type foo struct {
	bar string
}

func (f foo) Clone() foo {
	return foo{f.bar}
}
