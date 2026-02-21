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
type myInt8 int8
type myInt16 int16
type myInt32 int32
type myInt64 int64
type myUint8 uint8
type myUint16 uint16
type myUint32 uint32
type myUint64 uint64
type myFloat32 float32
type myFloat64 float64
