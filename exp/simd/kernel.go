//go:build goexperiment.simd

package simd

import (
	"simd"
	"slices"
)

// maxVectorBits is the widest vector archMaxVectorSize() can return today (AVX-512 / SVE).
// It bounds the fixed-size stack buffers below, so lane counts never require a heap
// allocation even though they are only known at run time.
const maxVectorBits = 512

// maxLanesN is the largest number of N-bit lanes a single vector can hold.
const (
	maxLanes8  = maxVectorBits / 8
	maxLanes16 = maxVectorBits / 16
	maxLanes32 = maxVectorBits / 32
	maxLanes64 = maxVectorBits / 64
)

// numeric is the constraint shared by every element type this package accelerates.
type numeric interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// clampScalar mirrors the vector kernels' Min(mx).Max(mn) composition. It matches lo.Clamp
// when mn <= mx. When mn > mx (a malformed precondition), it consistently returns mn for
// every value, rather than the value-dependent result of naively evaluating lo.Clamp's
// branches with swapped bounds — this is what keeps the SIMD and scalar-fallback code paths
// producing identical output for every input, on every platform.
func clampScalar[T numeric](v, mn, mx T) T {
	if v > mx {
		v = mx
	}
	if v < mn {
		v = mn
	}
	return v
}

// containsBlock is the number of vector compares accumulated (via Mask.Or) between checks
// of the running mask. The simd package exposes no bitmask extraction or horizontal
// reduction, so every check costs a Store plus a fold over at most maxLanes64 words; batching
// compares amortizes that cost while still allowing Contains to exit early on a hit.
const containsBlock = 8

// useSIMD reports whether the vector kernels in this package are worth using. Neither
// simd.Emulated nor simd.VectorBitSize mentions a simd type, so this initializer is not
// subject to the "SIMD-dependent var initializers don't work" limitation documented in the
// simd package (that limitation applies to vars whose declared type or initializer expression
// depends on a simd type, and is worked around here by never storing one in a package var).
//
// The emulated implementation is a scalar loop dressed up as vector code (shifts and masks
// per element), so it is always slower than github.com/samber/lo's plain scalar loop; this
// also covers GODEBUG=simd=0 and pre-AVX amd64 hardware.
var useSIMD = !simd.Emulated() && simd.VectorBitSize() <= maxVectorBits

// lanes8 returns the number of 8-bit lanes in the current vector width. Exported to tests
// so table sizes stay correct at every vector width instead of hardcoding one.
func lanes8() int {
	var v simd.Int8s
	return v.Len()
}

// lanes16 is lanes8 for 16-bit lanes.
func lanes16() int {
	var v simd.Int16s
	return v.Len()
}

// lanes32 is lanes8 for 32-bit lanes.
func lanes32() int {
	var v simd.Int32s
	return v.Len()
}

// lanes64 is lanes8 for 64-bit lanes.
func lanes64() int {
	var v simd.Int64s
	return v.Len()
}

// anyMask8 reports whether any lane of m is set. There is no bitmask extraction or
// horizontal reduction for masks in the simd package, so the mask is reinterpreted as
// 64-bit words (a bitwise cast, not a comparison) and OR-folded through a stack buffer.
func anyMask8(m simd.Mask8s, w *[maxLanes64]uint64) bool {
	u := m.ToInt8s().ToBits().ReshapeToUint64s()
	u.Store(w[:])
	var acc uint64
	for k := 0; k < u.Len(); k++ {
		acc |= w[k]
	}
	return acc != 0
}

// anyMask16 is anyMask8 for 16-bit lanes.
func anyMask16(m simd.Mask16s, w *[maxLanes64]uint64) bool {
	u := m.ToInt16s().ToBits().ReshapeToUint64s()
	u.Store(w[:])
	var acc uint64
	for k := 0; k < u.Len(); k++ {
		acc |= w[k]
	}
	return acc != 0
}

// anyMask32 is anyMask8 for 32-bit lanes.
func anyMask32(m simd.Mask32s, w *[maxLanes64]uint64) bool {
	u := m.ToInt32s().ToBits().ReshapeToUint64s()
	u.Store(w[:])
	var acc uint64
	for k := 0; k < u.Len(); k++ {
		acc |= w[k]
	}
	return acc != 0
}

// anyMask64 is anyMask8 for 64-bit lanes.
func anyMask64(m simd.Mask64s, w *[maxLanes64]uint64) bool {
	u := m.ToInt64s().ToBits()
	u.Store(w[:])
	var acc uint64
	for k := 0; k < u.Len(); k++ {
		acc |= w[k]
	}
	return acc != 0
}

// ============================================================================
// int8
// ============================================================================

// sumInt8 sums s, which the caller guarantees is non-empty.
func sumInt8(s []int8) int8 {
	var acc simd.Int8s
	lanes := acc.Len()

	i := 0
	for ; i+lanes <= len(s); i += lanes {
		acc = acc.Add(simd.LoadInt8s(s[i : i+lanes]))
	}
	if i < len(s) {
		// LoadPart zero-fills the unused lanes; 0 is the identity for Add, so the
		// tail needs no masking.
		v, _ := simd.LoadInt8sPart(s[i:])
		acc = acc.Add(v)
	}

	var buf [maxLanes8]int8
	acc.Store(buf[:])
	var total int8
	for k := range lanes {
		total += buf[k]
	}
	return total
}

// minInt8 returns the minimum of s, which the caller guarantees is non-empty.
func minInt8(s []int8) int8 {
	var acc simd.Int8s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v < m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadInt8s(s) // seed: zero is not neutral for Min
	for i := lanes; i+lanes <= len(s); i += lanes {
		acc = acc.Min(simd.LoadInt8s(s[i : i+lanes]))
	}
	if len(s)%lanes != 0 {
		// Overlapping final load: re-reading up to lanes-1 already-seen elements is
		// harmless because Min is idempotent, and it avoids a zero-filled LoadPart.
		acc = acc.Min(simd.LoadInt8s(s[len(s)-lanes:]))
	}

	var buf [maxLanes8]int8
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] < m {
			m = buf[k]
		}
	}
	return m
}

// maxInt8 returns the maximum of s, which the caller guarantees is non-empty.
func maxInt8(s []int8) int8 {
	var acc simd.Int8s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v > m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadInt8s(s)
	for i := lanes; i+lanes <= len(s); i += lanes {
		acc = acc.Max(simd.LoadInt8s(s[i : i+lanes]))
	}
	if len(s)%lanes != 0 {
		acc = acc.Max(simd.LoadInt8s(s[len(s)-lanes:]))
	}

	var buf [maxLanes8]int8
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] > m {
			m = buf[k]
		}
	}
	return m
}

// clampInt8 writes the clamped src into dst. The caller guarantees len(dst) == len(src) > 0
// and mn <= mx.
func clampInt8(dst, src []int8, mn, mx int8) {
	var v simd.Int8s
	lanes := v.Len()
	mnV := simd.BroadcastInt8s(mn)
	mxV := simd.BroadcastInt8s(mx)

	i := 0
	for ; i+lanes <= len(src); i += lanes {
		v = simd.LoadInt8s(src[i : i+lanes])
		v.Min(mxV).Max(mnV).Store(dst[i : i+lanes])
	}
	if i < len(src) {
		// LoadPart zero-fills the unused lanes, but StorePart writes only the
		// len(dst[i:]) < lanes valid lanes, so the padding never reaches dst.
		p, _ := simd.LoadInt8sPart(src[i:])
		p.Min(mxV).Max(mnV).StorePart(dst[i:])
	}
}

// containsInt8 reports whether s contains target.
func containsInt8(s []int8, target int8) bool {
	var v simd.Int8s
	lanes := v.Len()
	if len(s) < lanes {
		return slices.Contains(s, target)
	}

	t := simd.BroadcastInt8s(target)
	var w [maxLanes64]uint64

	i := 0
	for i+lanes <= len(s) {
		var m simd.Mask8s // zero value: all lanes false
		for b := 0; b < containsBlock && i+lanes <= len(s); b, i = b+1, i+lanes {
			m = m.Or(simd.LoadInt8s(s[i : i+lanes]).Equal(t))
		}
		if anyMask8(m, &w) {
			return true
		}
	}
	if i < len(s) {
		// Overlapping final load, never LoadPart: LoadPart zero-fills, which would
		// report a false positive when target == 0.
		if anyMask8(simd.LoadInt8s(s[len(s)-lanes:]).Equal(t), &w) {
			return true
		}
	}
	return false
}

// ============================================================================
// int16
// ============================================================================

func sumInt16(s []int16) int16 {
	var acc simd.Int16s
	lanes := acc.Len()

	i := 0
	for ; i+lanes <= len(s); i += lanes {
		acc = acc.Add(simd.LoadInt16s(s[i : i+lanes]))
	}
	if i < len(s) {
		v, _ := simd.LoadInt16sPart(s[i:])
		acc = acc.Add(v)
	}

	var buf [maxLanes16]int16
	acc.Store(buf[:])
	var total int16
	for k := range lanes {
		total += buf[k]
	}
	return total
}

func minInt16(s []int16) int16 {
	var acc simd.Int16s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v < m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadInt16s(s)
	for i := lanes; i+lanes <= len(s); i += lanes {
		acc = acc.Min(simd.LoadInt16s(s[i : i+lanes]))
	}
	if len(s)%lanes != 0 {
		acc = acc.Min(simd.LoadInt16s(s[len(s)-lanes:]))
	}

	var buf [maxLanes16]int16
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] < m {
			m = buf[k]
		}
	}
	return m
}

func maxInt16(s []int16) int16 {
	var acc simd.Int16s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v > m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadInt16s(s)
	for i := lanes; i+lanes <= len(s); i += lanes {
		acc = acc.Max(simd.LoadInt16s(s[i : i+lanes]))
	}
	if len(s)%lanes != 0 {
		acc = acc.Max(simd.LoadInt16s(s[len(s)-lanes:]))
	}

	var buf [maxLanes16]int16
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] > m {
			m = buf[k]
		}
	}
	return m
}

func clampInt16(dst, src []int16, mn, mx int16) {
	var v simd.Int16s
	lanes := v.Len()
	mnV := simd.BroadcastInt16s(mn)
	mxV := simd.BroadcastInt16s(mx)

	i := 0
	for ; i+lanes <= len(src); i += lanes {
		v = simd.LoadInt16s(src[i : i+lanes])
		v.Min(mxV).Max(mnV).Store(dst[i : i+lanes])
	}
	if i < len(src) {
		p, _ := simd.LoadInt16sPart(src[i:])
		p.Min(mxV).Max(mnV).StorePart(dst[i:])
	}
}

func containsInt16(s []int16, target int16) bool {
	var v simd.Int16s
	lanes := v.Len()
	if len(s) < lanes {
		return slices.Contains(s, target)
	}

	t := simd.BroadcastInt16s(target)
	var w [maxLanes64]uint64

	i := 0
	for i+lanes <= len(s) {
		var m simd.Mask16s
		for b := 0; b < containsBlock && i+lanes <= len(s); b, i = b+1, i+lanes {
			m = m.Or(simd.LoadInt16s(s[i : i+lanes]).Equal(t))
		}
		if anyMask16(m, &w) {
			return true
		}
	}
	if i < len(s) {
		if anyMask16(simd.LoadInt16s(s[len(s)-lanes:]).Equal(t), &w) {
			return true
		}
	}
	return false
}

// ============================================================================
// int32
// ============================================================================

func sumInt32(s []int32) int32 {
	var acc simd.Int32s
	lanes := acc.Len()

	i := 0
	for ; i+lanes <= len(s); i += lanes {
		acc = acc.Add(simd.LoadInt32s(s[i : i+lanes]))
	}
	if i < len(s) {
		v, _ := simd.LoadInt32sPart(s[i:])
		acc = acc.Add(v)
	}

	var buf [maxLanes32]int32
	acc.Store(buf[:])
	var total int32
	for k := range lanes {
		total += buf[k]
	}
	return total
}

func minInt32(s []int32) int32 {
	var acc simd.Int32s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v < m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadInt32s(s)
	for i := lanes; i+lanes <= len(s); i += lanes {
		acc = acc.Min(simd.LoadInt32s(s[i : i+lanes]))
	}
	if len(s)%lanes != 0 {
		acc = acc.Min(simd.LoadInt32s(s[len(s)-lanes:]))
	}

	var buf [maxLanes32]int32
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] < m {
			m = buf[k]
		}
	}
	return m
}

func maxInt32(s []int32) int32 {
	var acc simd.Int32s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v > m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadInt32s(s)
	for i := lanes; i+lanes <= len(s); i += lanes {
		acc = acc.Max(simd.LoadInt32s(s[i : i+lanes]))
	}
	if len(s)%lanes != 0 {
		acc = acc.Max(simd.LoadInt32s(s[len(s)-lanes:]))
	}

	var buf [maxLanes32]int32
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] > m {
			m = buf[k]
		}
	}
	return m
}

func clampInt32(dst, src []int32, mn, mx int32) {
	var v simd.Int32s
	lanes := v.Len()
	mnV := simd.BroadcastInt32s(mn)
	mxV := simd.BroadcastInt32s(mx)

	i := 0
	for ; i+lanes <= len(src); i += lanes {
		v = simd.LoadInt32s(src[i : i+lanes])
		v.Min(mxV).Max(mnV).Store(dst[i : i+lanes])
	}
	if i < len(src) {
		p, _ := simd.LoadInt32sPart(src[i:])
		p.Min(mxV).Max(mnV).StorePart(dst[i:])
	}
}

func containsInt32(s []int32, target int32) bool {
	var v simd.Int32s
	lanes := v.Len()
	if len(s) < lanes {
		return slices.Contains(s, target)
	}

	t := simd.BroadcastInt32s(target)
	var w [maxLanes64]uint64

	i := 0
	for i+lanes <= len(s) {
		var m simd.Mask32s
		for b := 0; b < containsBlock && i+lanes <= len(s); b, i = b+1, i+lanes {
			m = m.Or(simd.LoadInt32s(s[i : i+lanes]).Equal(t))
		}
		if anyMask32(m, &w) {
			return true
		}
	}
	if i < len(s) {
		if anyMask32(simd.LoadInt32s(s[len(s)-lanes:]).Equal(t), &w) {
			return true
		}
	}
	return false
}

// ============================================================================
// int64 (no hardware Min/Max: emulated via IfElse)
// ============================================================================

func sumInt64(s []int64) int64 {
	var acc simd.Int64s
	lanes := acc.Len()

	i := 0
	for ; i+lanes <= len(s); i += lanes {
		acc = acc.Add(simd.LoadInt64s(s[i : i+lanes]))
	}
	if i < len(s) {
		v, _ := simd.LoadInt64sPart(s[i:])
		acc = acc.Add(v)
	}

	var buf [maxLanes64]int64
	acc.Store(buf[:])
	var total int64
	for k := range lanes {
		total += buf[k]
	}
	return total
}

func minInt64(s []int64) int64 {
	var acc simd.Int64s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v < m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadInt64s(s)
	for i := lanes; i+lanes <= len(s); i += lanes {
		v := simd.LoadInt64s(s[i : i+lanes])
		acc = acc.IfElse(acc.Less(v), v) // min(acc, v)
	}
	if len(s)%lanes != 0 {
		v := simd.LoadInt64s(s[len(s)-lanes:])
		acc = acc.IfElse(acc.Less(v), v)
	}

	var buf [maxLanes64]int64
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] < m {
			m = buf[k]
		}
	}
	return m
}

func maxInt64(s []int64) int64 {
	var acc simd.Int64s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v > m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadInt64s(s)
	for i := lanes; i+lanes <= len(s); i += lanes {
		v := simd.LoadInt64s(s[i : i+lanes])
		acc = acc.IfElse(acc.Greater(v), v) // max(acc, v)
	}
	if len(s)%lanes != 0 {
		v := simd.LoadInt64s(s[len(s)-lanes:])
		acc = acc.IfElse(acc.Greater(v), v)
	}

	var buf [maxLanes64]int64
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] > m {
			m = buf[k]
		}
	}
	return m
}

func clampInt64(dst, src []int64, mn, mx int64) {
	var v simd.Int64s
	lanes := v.Len()
	mnV := simd.BroadcastInt64s(mn)
	mxV := simd.BroadcastInt64s(mx)

	i := 0
	for ; i+lanes <= len(src); i += lanes {
		v = simd.LoadInt64s(src[i : i+lanes])
		c := v.IfElse(v.Less(mxV), mxV)   // min(v, mx)
		c = c.IfElse(c.Greater(mnV), mnV) // max(c, mn)
		c.Store(dst[i : i+lanes])
	}
	if i < len(src) {
		p, _ := simd.LoadInt64sPart(src[i:])
		c := p.IfElse(p.Less(mxV), mxV)
		c = c.IfElse(c.Greater(mnV), mnV)
		c.StorePart(dst[i:])
	}
}

func containsInt64(s []int64, target int64) bool {
	var v simd.Int64s
	lanes := v.Len()
	if len(s) < lanes {
		return slices.Contains(s, target)
	}

	t := simd.BroadcastInt64s(target)
	var w [maxLanes64]uint64

	i := 0
	for i+lanes <= len(s) {
		var m simd.Mask64s
		for b := 0; b < containsBlock && i+lanes <= len(s); b, i = b+1, i+lanes {
			m = m.Or(simd.LoadInt64s(s[i : i+lanes]).Equal(t))
		}
		if anyMask64(m, &w) {
			return true
		}
	}
	if i < len(s) {
		if anyMask64(simd.LoadInt64s(s[len(s)-lanes:]).Equal(t), &w) {
			return true
		}
	}
	return false
}

// ============================================================================
// uint8
// ============================================================================

func sumUint8(s []uint8) uint8 {
	var acc simd.Uint8s
	lanes := acc.Len()

	i := 0
	for ; i+lanes <= len(s); i += lanes {
		acc = acc.Add(simd.LoadUint8s(s[i : i+lanes]))
	}
	if i < len(s) {
		v, _ := simd.LoadUint8sPart(s[i:])
		acc = acc.Add(v)
	}

	var buf [maxLanes8]uint8
	acc.Store(buf[:])
	var total uint8
	for k := range lanes {
		total += buf[k]
	}
	return total
}

func minUint8(s []uint8) uint8 {
	var acc simd.Uint8s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v < m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadUint8s(s)
	for i := lanes; i+lanes <= len(s); i += lanes {
		acc = acc.Min(simd.LoadUint8s(s[i : i+lanes]))
	}
	if len(s)%lanes != 0 {
		acc = acc.Min(simd.LoadUint8s(s[len(s)-lanes:]))
	}

	var buf [maxLanes8]uint8
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] < m {
			m = buf[k]
		}
	}
	return m
}

func maxUint8(s []uint8) uint8 {
	var acc simd.Uint8s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v > m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadUint8s(s)
	for i := lanes; i+lanes <= len(s); i += lanes {
		acc = acc.Max(simd.LoadUint8s(s[i : i+lanes]))
	}
	if len(s)%lanes != 0 {
		acc = acc.Max(simd.LoadUint8s(s[len(s)-lanes:]))
	}

	var buf [maxLanes8]uint8
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] > m {
			m = buf[k]
		}
	}
	return m
}

func clampUint8(dst, src []uint8, mn, mx uint8) {
	var v simd.Uint8s
	lanes := v.Len()
	mnV := simd.BroadcastUint8s(mn)
	mxV := simd.BroadcastUint8s(mx)

	i := 0
	for ; i+lanes <= len(src); i += lanes {
		v = simd.LoadUint8s(src[i : i+lanes])
		v.Min(mxV).Max(mnV).Store(dst[i : i+lanes])
	}
	if i < len(src) {
		p, _ := simd.LoadUint8sPart(src[i:])
		p.Min(mxV).Max(mnV).StorePart(dst[i:])
	}
}

func containsUint8(s []uint8, target uint8) bool {
	var v simd.Uint8s
	lanes := v.Len()
	if len(s) < lanes {
		return slices.Contains(s, target)
	}

	t := simd.BroadcastUint8s(target)
	var w [maxLanes64]uint64

	i := 0
	for i+lanes <= len(s) {
		var m simd.Mask8s
		for b := 0; b < containsBlock && i+lanes <= len(s); b, i = b+1, i+lanes {
			m = m.Or(simd.LoadUint8s(s[i : i+lanes]).Equal(t))
		}
		if anyMask8(m, &w) {
			return true
		}
	}
	if i < len(s) {
		if anyMask8(simd.LoadUint8s(s[len(s)-lanes:]).Equal(t), &w) {
			return true
		}
	}
	return false
}

// ============================================================================
// uint16
// ============================================================================

func sumUint16(s []uint16) uint16 {
	var acc simd.Uint16s
	lanes := acc.Len()

	i := 0
	for ; i+lanes <= len(s); i += lanes {
		acc = acc.Add(simd.LoadUint16s(s[i : i+lanes]))
	}
	if i < len(s) {
		v, _ := simd.LoadUint16sPart(s[i:])
		acc = acc.Add(v)
	}

	var buf [maxLanes16]uint16
	acc.Store(buf[:])
	var total uint16
	for k := range lanes {
		total += buf[k]
	}
	return total
}

func minUint16(s []uint16) uint16 {
	var acc simd.Uint16s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v < m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadUint16s(s)
	for i := lanes; i+lanes <= len(s); i += lanes {
		acc = acc.Min(simd.LoadUint16s(s[i : i+lanes]))
	}
	if len(s)%lanes != 0 {
		acc = acc.Min(simd.LoadUint16s(s[len(s)-lanes:]))
	}

	var buf [maxLanes16]uint16
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] < m {
			m = buf[k]
		}
	}
	return m
}

func maxUint16(s []uint16) uint16 {
	var acc simd.Uint16s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v > m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadUint16s(s)
	for i := lanes; i+lanes <= len(s); i += lanes {
		acc = acc.Max(simd.LoadUint16s(s[i : i+lanes]))
	}
	if len(s)%lanes != 0 {
		acc = acc.Max(simd.LoadUint16s(s[len(s)-lanes:]))
	}

	var buf [maxLanes16]uint16
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] > m {
			m = buf[k]
		}
	}
	return m
}

func clampUint16(dst, src []uint16, mn, mx uint16) {
	var v simd.Uint16s
	lanes := v.Len()
	mnV := simd.BroadcastUint16s(mn)
	mxV := simd.BroadcastUint16s(mx)

	i := 0
	for ; i+lanes <= len(src); i += lanes {
		v = simd.LoadUint16s(src[i : i+lanes])
		v.Min(mxV).Max(mnV).Store(dst[i : i+lanes])
	}
	if i < len(src) {
		p, _ := simd.LoadUint16sPart(src[i:])
		p.Min(mxV).Max(mnV).StorePart(dst[i:])
	}
}

func containsUint16(s []uint16, target uint16) bool {
	var v simd.Uint16s
	lanes := v.Len()
	if len(s) < lanes {
		return slices.Contains(s, target)
	}

	t := simd.BroadcastUint16s(target)
	var w [maxLanes64]uint64

	i := 0
	for i+lanes <= len(s) {
		var m simd.Mask16s
		for b := 0; b < containsBlock && i+lanes <= len(s); b, i = b+1, i+lanes {
			m = m.Or(simd.LoadUint16s(s[i : i+lanes]).Equal(t))
		}
		if anyMask16(m, &w) {
			return true
		}
	}
	if i < len(s) {
		if anyMask16(simd.LoadUint16s(s[len(s)-lanes:]).Equal(t), &w) {
			return true
		}
	}
	return false
}

// ============================================================================
// uint32
// ============================================================================

func sumUint32(s []uint32) uint32 {
	var acc simd.Uint32s
	lanes := acc.Len()

	i := 0
	for ; i+lanes <= len(s); i += lanes {
		acc = acc.Add(simd.LoadUint32s(s[i : i+lanes]))
	}
	if i < len(s) {
		v, _ := simd.LoadUint32sPart(s[i:])
		acc = acc.Add(v)
	}

	var buf [maxLanes32]uint32
	acc.Store(buf[:])
	var total uint32
	for k := range lanes {
		total += buf[k]
	}
	return total
}

func minUint32(s []uint32) uint32 {
	var acc simd.Uint32s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v < m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadUint32s(s)
	for i := lanes; i+lanes <= len(s); i += lanes {
		acc = acc.Min(simd.LoadUint32s(s[i : i+lanes]))
	}
	if len(s)%lanes != 0 {
		acc = acc.Min(simd.LoadUint32s(s[len(s)-lanes:]))
	}

	var buf [maxLanes32]uint32
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] < m {
			m = buf[k]
		}
	}
	return m
}

func maxUint32(s []uint32) uint32 {
	var acc simd.Uint32s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v > m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadUint32s(s)
	for i := lanes; i+lanes <= len(s); i += lanes {
		acc = acc.Max(simd.LoadUint32s(s[i : i+lanes]))
	}
	if len(s)%lanes != 0 {
		acc = acc.Max(simd.LoadUint32s(s[len(s)-lanes:]))
	}

	var buf [maxLanes32]uint32
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] > m {
			m = buf[k]
		}
	}
	return m
}

func clampUint32(dst, src []uint32, mn, mx uint32) {
	var v simd.Uint32s
	lanes := v.Len()
	mnV := simd.BroadcastUint32s(mn)
	mxV := simd.BroadcastUint32s(mx)

	i := 0
	for ; i+lanes <= len(src); i += lanes {
		v = simd.LoadUint32s(src[i : i+lanes])
		v.Min(mxV).Max(mnV).Store(dst[i : i+lanes])
	}
	if i < len(src) {
		p, _ := simd.LoadUint32sPart(src[i:])
		p.Min(mxV).Max(mnV).StorePart(dst[i:])
	}
}

func containsUint32(s []uint32, target uint32) bool {
	var v simd.Uint32s
	lanes := v.Len()
	if len(s) < lanes {
		return slices.Contains(s, target)
	}

	t := simd.BroadcastUint32s(target)
	var w [maxLanes64]uint64

	i := 0
	for i+lanes <= len(s) {
		var m simd.Mask32s
		for b := 0; b < containsBlock && i+lanes <= len(s); b, i = b+1, i+lanes {
			m = m.Or(simd.LoadUint32s(s[i : i+lanes]).Equal(t))
		}
		if anyMask32(m, &w) {
			return true
		}
	}
	if i < len(s) {
		if anyMask32(simd.LoadUint32s(s[len(s)-lanes:]).Equal(t), &w) {
			return true
		}
	}
	return false
}

// ============================================================================
// uint64 (no hardware Min/Max: emulated via IfElse)
// ============================================================================

func sumUint64(s []uint64) uint64 {
	var acc simd.Uint64s
	lanes := acc.Len()

	i := 0
	for ; i+lanes <= len(s); i += lanes {
		acc = acc.Add(simd.LoadUint64s(s[i : i+lanes]))
	}
	if i < len(s) {
		v, _ := simd.LoadUint64sPart(s[i:])
		acc = acc.Add(v)
	}

	var buf [maxLanes64]uint64
	acc.Store(buf[:])
	var total uint64
	for k := range lanes {
		total += buf[k]
	}
	return total
}

func minUint64(s []uint64) uint64 {
	var acc simd.Uint64s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v < m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadUint64s(s)
	for i := lanes; i+lanes <= len(s); i += lanes {
		v := simd.LoadUint64s(s[i : i+lanes])
		acc = acc.IfElse(acc.Less(v), v)
	}
	if len(s)%lanes != 0 {
		v := simd.LoadUint64s(s[len(s)-lanes:])
		acc = acc.IfElse(acc.Less(v), v)
	}

	var buf [maxLanes64]uint64
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] < m {
			m = buf[k]
		}
	}
	return m
}

func maxUint64(s []uint64) uint64 {
	var acc simd.Uint64s
	lanes := acc.Len()

	if len(s) < lanes {
		m := s[0]
		for _, v := range s[1:] {
			if v > m {
				m = v
			}
		}
		return m
	}

	acc = simd.LoadUint64s(s)
	for i := lanes; i+lanes <= len(s); i += lanes {
		v := simd.LoadUint64s(s[i : i+lanes])
		acc = acc.IfElse(acc.Greater(v), v)
	}
	if len(s)%lanes != 0 {
		v := simd.LoadUint64s(s[len(s)-lanes:])
		acc = acc.IfElse(acc.Greater(v), v)
	}

	var buf [maxLanes64]uint64
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] > m {
			m = buf[k]
		}
	}
	return m
}

func clampUint64(dst, src []uint64, mn, mx uint64) {
	var v simd.Uint64s
	lanes := v.Len()
	mnV := simd.BroadcastUint64s(mn)
	mxV := simd.BroadcastUint64s(mx)

	i := 0
	for ; i+lanes <= len(src); i += lanes {
		v = simd.LoadUint64s(src[i : i+lanes])
		c := v.IfElse(v.Less(mxV), mxV)
		c = c.IfElse(c.Greater(mnV), mnV)
		c.Store(dst[i : i+lanes])
	}
	if i < len(src) {
		p, _ := simd.LoadUint64sPart(src[i:])
		c := p.IfElse(p.Less(mxV), mxV)
		c = c.IfElse(c.Greater(mnV), mnV)
		c.StorePart(dst[i:])
	}
}

func containsUint64(s []uint64, target uint64) bool {
	var v simd.Uint64s
	lanes := v.Len()
	if len(s) < lanes {
		return slices.Contains(s, target)
	}

	t := simd.BroadcastUint64s(target)
	var w [maxLanes64]uint64

	i := 0
	for i+lanes <= len(s) {
		var m simd.Mask64s
		for b := 0; b < containsBlock && i+lanes <= len(s); b, i = b+1, i+lanes {
			m = m.Or(simd.LoadUint64s(s[i : i+lanes]).Equal(t))
		}
		if anyMask64(m, &w) {
			return true
		}
	}
	if i < len(s) {
		if anyMask64(simd.LoadUint64s(s[len(s)-lanes:]).Equal(t), &w) {
			return true
		}
	}
	return false
}

// ============================================================================
// float32
// ============================================================================

// sumFloat32 accumulates lanes independently (element i goes into lane i%lanes) and only
// combines them in the final horizontal fold below. Floating-point addition is not
// associative, so this reordering can change rounding relative to a strictly sequential
// sum — the result is numerically equivalent to, but not always bit-identical to, scalarSum.
func sumFloat32(s []float32) float32 {
	var acc simd.Float32s
	lanes := acc.Len()

	i := 0
	for ; i+lanes <= len(s); i += lanes {
		acc = acc.Add(simd.LoadFloat32s(s[i : i+lanes]))
	}
	if i < len(s) {
		v, _ := simd.LoadFloat32sPart(s[i:])
		acc = acc.Add(v)
	}

	var buf [maxLanes32]float32
	acc.Store(buf[:])
	var total float32
	for k := range lanes { // final horizontal fold: another reordering relative to scalarSum
		total += buf[k]
	}
	return total
}

// scalarMinFloat32 is lo.Min's exact algorithm: if s[0] is NaN, the result is NaN (nothing
// ever compares less than NaN); otherwise NaN elsewhere in s can never win the comparison,
// so it is effectively ignored.
func scalarMinFloat32(s []float32) float32 {
	m := s[0]
	for _, v := range s[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

// scalarMaxFloat32 is scalarMinFloat32 with the comparison flipped; see it for NaN handling.
func scalarMaxFloat32(s []float32) float32 {
	m := s[0]
	for _, v := range s[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

// minFloat32 returns the minimum of s, which the caller guarantees is non-empty.
//
// Hardware Min discards a NaN operand in favour of the other one (or, on some
// architectures, propagates it), which does not match scalarMinFloat32's "only s[0] can
// produce a NaN result" rule. Rather than accept that divergence, every vector iteration
// also folds a "was this lane ever NaN" mask (cost: one extra compare per iteration, no
// extra memory traffic); if any NaN was seen anywhere in s, the whole call is redone with
// scalarMinFloat32 instead of trusting the vectorized reduction. NaN is rare in practice, so
// this keeps the common case fast while making the result identical to lo.Min in every case.
func minFloat32(s []float32) float32 {
	var acc simd.Float32s
	lanes := acc.Len()

	if len(s) < lanes {
		return scalarMinFloat32(s)
	}

	acc = simd.LoadFloat32s(s)
	nan := acc.NotEqual(acc) // true only in lanes that are NaN (NaN != NaN)
	for i := lanes; i+lanes <= len(s); i += lanes {
		v := simd.LoadFloat32s(s[i : i+lanes])
		nan = nan.Or(v.NotEqual(v))
		acc = acc.Min(v)
	}
	if len(s)%lanes != 0 {
		v := simd.LoadFloat32s(s[len(s)-lanes:])
		nan = nan.Or(v.NotEqual(v))
		acc = acc.Min(v)
	}

	var w [maxLanes64]uint64
	if anyMask32(nan, &w) {
		return scalarMinFloat32(s)
	}

	var buf [maxLanes32]float32
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] < m {
			m = buf[k]
		}
	}
	return m
}

// maxFloat32 returns the maximum of s, which the caller guarantees is non-empty. See
// minFloat32 for the NaN-detection strategy, which is identical here.
func maxFloat32(s []float32) float32 {
	var acc simd.Float32s
	lanes := acc.Len()

	if len(s) < lanes {
		return scalarMaxFloat32(s)
	}

	acc = simd.LoadFloat32s(s)
	nan := acc.NotEqual(acc)
	for i := lanes; i+lanes <= len(s); i += lanes {
		v := simd.LoadFloat32s(s[i : i+lanes])
		nan = nan.Or(v.NotEqual(v))
		acc = acc.Max(v)
	}
	if len(s)%lanes != 0 {
		v := simd.LoadFloat32s(s[len(s)-lanes:])
		nan = nan.Or(v.NotEqual(v))
		acc = acc.Max(v)
	}

	var w [maxLanes64]uint64
	if anyMask32(nan, &w) {
		return scalarMaxFloat32(s)
	}

	var buf [maxLanes32]float32
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] > m {
			m = buf[k]
		}
	}
	return m
}

// clampFloat32 writes the clamped src into dst. Unlike Min/Max, Clamp is element-wise, so
// NaN handling needs no whole-slice detection: a NaN input has no ordering relative to mn or
// mx, so lo.Clamp's `<`/`>` comparisons are both false and it returns the value unchanged.
// Each lane restores that original (possibly NaN) value wherever the input was NaN, in
// place of whatever the hardware Min/Max composition computed for that lane.
func clampFloat32(dst, src []float32, mn, mx float32) {
	var v simd.Float32s
	lanes := v.Len()
	mnV := simd.BroadcastFloat32s(mn)
	mxV := simd.BroadcastFloat32s(mx)

	i := 0
	for ; i+lanes <= len(src); i += lanes {
		v = simd.LoadFloat32s(src[i : i+lanes])
		clamped := v.Min(mxV).Max(mnV)
		v.IfElse(v.NotEqual(v), clamped).Store(dst[i : i+lanes])
	}
	if i < len(src) {
		p, _ := simd.LoadFloat32sPart(src[i:])
		clamped := p.Min(mxV).Max(mnV)
		p.IfElse(p.NotEqual(p), clamped).StorePart(dst[i:])
	}
}

func containsFloat32(s []float32, target float32) bool {
	var v simd.Float32s
	lanes := v.Len()
	if len(s) < lanes {
		return slices.Contains(s, target)
	}

	t := simd.BroadcastFloat32s(target)
	var w [maxLanes64]uint64

	i := 0
	for i+lanes <= len(s) {
		var m simd.Mask32s
		for b := 0; b < containsBlock && i+lanes <= len(s); b, i = b+1, i+lanes {
			m = m.Or(simd.LoadFloat32s(s[i : i+lanes]).Equal(t))
		}
		if anyMask32(m, &w) {
			return true
		}
	}
	if i < len(s) {
		if anyMask32(simd.LoadFloat32s(s[len(s)-lanes:]).Equal(t), &w) {
			return true
		}
	}
	return false
}

// ============================================================================
// float64
// ============================================================================

// sumFloat64 accumulates lanes independently and only combines them in the final
// horizontal fold below; see sumFloat32 for why this can change rounding relative to a
// strictly sequential sum.
func sumFloat64(s []float64) float64 {
	var acc simd.Float64s
	lanes := acc.Len()

	i := 0
	for ; i+lanes <= len(s); i += lanes {
		acc = acc.Add(simd.LoadFloat64s(s[i : i+lanes]))
	}
	if i < len(s) {
		v, _ := simd.LoadFloat64sPart(s[i:])
		acc = acc.Add(v)
	}

	var buf [maxLanes64]float64
	acc.Store(buf[:])
	var total float64
	for k := range lanes { // final horizontal fold: another reordering relative to scalarSum
		total += buf[k]
	}
	return total
}

// scalarMinFloat64 is scalarMinFloat32 for float64; see it for NaN handling.
func scalarMinFloat64(s []float64) float64 {
	m := s[0]
	for _, v := range s[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

// scalarMaxFloat64 is scalarMaxFloat32 for float64; see scalarMinFloat32 for NaN handling.
func scalarMaxFloat64(s []float64) float64 {
	m := s[0]
	for _, v := range s[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

// minFloat64 returns the minimum of s, which the caller guarantees is non-empty. See
// minFloat32 for the NaN-detection strategy, which is identical here.
func minFloat64(s []float64) float64 {
	var acc simd.Float64s
	lanes := acc.Len()

	if len(s) < lanes {
		return scalarMinFloat64(s)
	}

	acc = simd.LoadFloat64s(s)
	nan := acc.NotEqual(acc)
	for i := lanes; i+lanes <= len(s); i += lanes {
		v := simd.LoadFloat64s(s[i : i+lanes])
		nan = nan.Or(v.NotEqual(v))
		acc = acc.Min(v)
	}
	if len(s)%lanes != 0 {
		v := simd.LoadFloat64s(s[len(s)-lanes:])
		nan = nan.Or(v.NotEqual(v))
		acc = acc.Min(v)
	}

	var w [maxLanes64]uint64
	if anyMask64(nan, &w) {
		return scalarMinFloat64(s)
	}

	var buf [maxLanes64]float64
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] < m {
			m = buf[k]
		}
	}
	return m
}

// maxFloat64 returns the maximum of s, which the caller guarantees is non-empty. See
// minFloat32 for the NaN-detection strategy, which is identical here.
func maxFloat64(s []float64) float64 {
	var acc simd.Float64s
	lanes := acc.Len()

	if len(s) < lanes {
		return scalarMaxFloat64(s)
	}

	acc = simd.LoadFloat64s(s)
	nan := acc.NotEqual(acc)
	for i := lanes; i+lanes <= len(s); i += lanes {
		v := simd.LoadFloat64s(s[i : i+lanes])
		nan = nan.Or(v.NotEqual(v))
		acc = acc.Max(v)
	}
	if len(s)%lanes != 0 {
		v := simd.LoadFloat64s(s[len(s)-lanes:])
		nan = nan.Or(v.NotEqual(v))
		acc = acc.Max(v)
	}

	var w [maxLanes64]uint64
	if anyMask64(nan, &w) {
		return scalarMaxFloat64(s)
	}

	var buf [maxLanes64]float64
	acc.Store(buf[:])
	m := buf[0]
	for k := 1; k < lanes; k++ {
		if buf[k] > m {
			m = buf[k]
		}
	}
	return m
}

// clampFloat64 writes the clamped src into dst. See clampFloat32 for the NaN handling
// strategy, which is identical here.
func clampFloat64(dst, src []float64, mn, mx float64) {
	var v simd.Float64s
	lanes := v.Len()
	mnV := simd.BroadcastFloat64s(mn)
	mxV := simd.BroadcastFloat64s(mx)

	i := 0
	for ; i+lanes <= len(src); i += lanes {
		v = simd.LoadFloat64s(src[i : i+lanes])
		clamped := v.Min(mxV).Max(mnV)
		v.IfElse(v.NotEqual(v), clamped).Store(dst[i : i+lanes])
	}
	if i < len(src) {
		p, _ := simd.LoadFloat64sPart(src[i:])
		clamped := p.Min(mxV).Max(mnV)
		p.IfElse(p.NotEqual(p), clamped).StorePart(dst[i:])
	}
}

func containsFloat64(s []float64, target float64) bool {
	var v simd.Float64s
	lanes := v.Len()
	if len(s) < lanes {
		return slices.Contains(s, target)
	}

	t := simd.BroadcastFloat64s(target)
	var w [maxLanes64]uint64

	i := 0
	for i+lanes <= len(s) {
		var m simd.Mask64s
		for b := 0; b < containsBlock && i+lanes <= len(s); b, i = b+1, i+lanes {
			m = m.Or(simd.LoadFloat64s(s[i : i+lanes]).Equal(t))
		}
		if anyMask64(m, &w) {
			return true
		}
	}
	if i < len(s) {
		if anyMask64(simd.LoadFloat64s(s[len(s)-lanes:]).Equal(t), &w) {
			return true
		}
	}
	return false
}
