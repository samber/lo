//go:build go1.23

package it

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		size     int
		expected []string
	}{
		{
			name:     "size smaller than length, remainder",
			input:    "12345",
			size:     2,
			expected: []string{"12", "34", "5"},
		},
		{
			name:     "size smaller than length, exact",
			input:    "123456",
			size:     2,
			expected: []string{"12", "34", "56"},
		},
		{
			name:     "size equal to length",
			input:    "123456",
			size:     6,
			expected: []string{"123456"},
		},
		{
			name:     "size greater than length",
			input:    "123456",
			size:     10,
			expected: []string{"123456"},
		},
		{
			name:     "empty string",
			input:    "",
			size:     2,
			expected: []string{""}, // @TODO: should be [] - see https://github.com/samber/lo/issues/788
		},
		{
			name:     "multi-byte runes",
			input:    "明1好休2林森",
			size:     2,
			expected: []string{"明1", "好休", "2林", "森"},
		},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			result := ChunkString(tt.input, tt.size)
			is.Equal(tt.expected, slices.Collect(result))
		})
	}

	t.Run("panics on non-positive size", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.PanicsWithValue("it.ChunkString: size must be greater than 0", func() {
			ChunkString("12345", 0)
		})
	})
}
