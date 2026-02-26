//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"fmt"
	"os"
	"testing"

	"simd/archsimd"
)

// skipHelper is a small interface implemented by both *testing.T and *testing.B
// to allow unified CPU feature requirement checking for both tests and benchmarks.
type skipHelper interface {
	Helper()
	Skipf(format string, args ...any)
}

// How to check if your Linux CPU supports SIMD (avoids SIGILL):
//
//   grep -E 'avx' /proc/cpuinfo
//
// Or:  lscpu | grep -i avx
//
// You need:
//   - AVX tests (128-bit):  avx in flags (baseline on amd64)
//   - AVX2 tests (256-bit):  avx2  in flags
//   - AVX-512 tests:        avx512f (and often avx512bw, avx512vl)
//
// If your CPU lacks AVX or AVX2 or AVX-512, tests that use them will be skipped automatically.

// requireAVX skips the test/benchmark if the CPU does not support AVX (128-bit SIMD).
// Use at the start of each AVX test/benchmark to avoid SIGILL on older or non-x86 systems.
func requireAVX(t skipHelper) {
	t.Helper()
	if !archsimd.X86.AVX() {
		t.Skipf("CPU does not support AVX; skipping. Check compatibility: grep avx /proc/cpuinfo")
	}
}

// requireAVX2 skips the test/benchmark if the CPU does not support AVX2 (256-bit SIMD).
// Use at the start of each AVX2 test/benchmark to avoid SIGILL on older or non-x86 systems.
func requireAVX2(t skipHelper) {
	t.Helper()
	if !archsimd.X86.AVX2() {
		t.Skipf("CPU does not support AVX2; skipping. Check compatibility: grep avx2 /proc/cpuinfo")
	}
}

// requireAVX512 skips the test/benchmark if the CPU does not support AVX-512 Foundation.
// Use at the start of each AVX-512 test/benchmark to avoid SIGILL on CPUs without AVX-512.
func requireAVX512(t skipHelper) {
	t.Helper()
	if !archsimd.X86.AVX512() {
		t.Skipf("CPU does not support AVX-512; skipping. Check compatibility: grep avx512 /proc/cpuinfo")
	}
}

// PrintCPUFeatures prints detected x86 SIMD features (for debugging).
// Run: go test -run PrintCPUFeatures -v
func PrintCPUFeatures(t *testing.T) {
	fmt.Fprintf(os.Stdout, "X86 HasAVX=%v HasAVX2=%v HasAVX512=%v\n",
		archsimd.X86.AVX(), archsimd.X86.AVX2(), archsimd.X86.AVX512())
}
