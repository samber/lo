//go:build goexperiment.simd

package simd

import (
	"fmt"
	"os"
	"strings"
)

// Named types used throughout the test tables to exercise the ~T generic constraints (as
// opposed to the predeclared types themselves).
type (
	myInt8    int8
	myInt16   int16
	myInt32   int32
	myInt64   int64
	myUint8   uint8
	myUint16  uint16
	myUint32  uint32
	myUint64  uint64
	myFloat32 float32
	myFloat64 float64
)

func init() {
	for _, arg := range os.Args {
		bench, ok := strings.CutPrefix(arg, "-test.bench=")
		if ok && bench != "" && bench != "none" {
			_, _ = fmt.Fprintf(os.Stdout, "simd: useSIMD=%v lanes8=%d lanes16=%d lanes32=%d lanes64=%d\n",
				useSIMD, lanes8(), lanes16(), lanes32(), lanes64())
			break
		}
	}
}

// sizesAround returns a spread of slice lengths around lanes: the values that exercise the
// LoadPart/overlap-load boundaries regardless of the current vector width (128/256/512 bits).
func sizesAround(lanes int) []int {
	sizes := []int{0, 1, lanes - 1, lanes, lanes + 1, 2 * lanes, 2*lanes + 1}
	out := make([]int, 0, len(sizes))
	seen := make(map[int]bool, len(sizes))
	for _, s := range sizes {
		if s >= 0 && !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}
