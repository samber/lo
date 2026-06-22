//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"fmt"
	"os"
	"strings"

	"simd/archsimd"
)

func init() {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.bench=") {
			bench := strings.TrimPrefix(arg, "-test.bench=")
			if bench != "" && bench != "none" {
				fmt.Fprintf(os.Stdout, "archsimd.X86: AVX=%v AVX2=%v AVX512=%v\n",
					archsimd.X86.AVX(), archsimd.X86.AVX2(), archsimd.X86.AVX512())
				break
			}
		}
	}
}

// Type aliases for testing
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
