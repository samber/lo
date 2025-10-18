package lo

import (
	"time"
)

// Duration returns the time taken to execute a function.
// Play: https://go.dev/play/p/HQfbBbAXaFP
func Duration(cb func()) time.Duration {
	return Duration0(cb)
}

// Duration0 returns the time taken to execute a function.
// Play: https://go.dev/play/p/HQfbBbAXaFP
func Duration0(cb func()) time.Duration {
	start := time.Now()
	cb()
	return time.Since(start)
}

// Duration1 returns the time taken to execute a function.
// Play: https://go.dev/play/p/HQfbBbAXaFP
func Duration1[A any](cb func() A) (A, time.Duration) {
	start := time.Now()
	a := cb()
	return a, time.Since(start)
}

// Duration2 returns the time taken to execute a function.
// Play: https://go.dev/play/p/HQfbBbAXaFP
func Duration2[A, B any](cb func() (A, B)) (A, B, time.Duration) {
	start := time.Now()
	a, b := cb()
	return a, b, time.Since(start)
}

// Duration3 returns the time taken to execute a function.
// Play: https://go.dev/play/p/xr863iwkAxQ
func Duration3[A, B, C any](cb func() (A, B, C)) (A, B, C, time.Duration) {
	start := time.Now()
	a, b, c := cb()
	return a, b, c, time.Since(start)
}

// Duration4 returns the time taken to execute a function.
// Play: https://go.dev/play/p/xr863iwkAxQ
func Duration4[A, B, C, D any](cb func() (A, B, C, D)) (A, B, C, D, time.Duration) {
	start := time.Now()
	a, b, c, d := cb()
	return a, b, c, d, time.Since(start)
}

// Duration5 returns the time taken to execute a function.
// Play: https://go.dev/play/p/xr863iwkAxQ
func Duration5[A, B, C, D, E any](cb func() (A, B, C, D, E)) (A, B, C, D, E, time.Duration) {
	start := time.Now()
	a, b, c, d, e := cb()
	return a, b, c, d, e, time.Since(start)
}

// Duration6 returns the time taken to execute a function.
// Play: https://go.dev/play/p/mR4bTQKO-Tf
func Duration6[A, B, C, D, E, F any](cb func() (A, B, C, D, E, F)) (A, B, C, D, E, F, time.Duration) {
	start := time.Now()
	a, b, c, d, e, f := cb()
	return a, b, c, d, e, f, time.Since(start)
}

// Duration7 returns the time taken to execute a function.
// Play: https://go.dev/play/p/jgIAcBWWInS
func Duration7[A, B, C, D, E, F, G any](cb func() (A, B, C, D, E, F, G)) (A, B, C, D, E, F, G, time.Duration) {
	start := time.Now()
	a, b, c, d, e, f, g := cb()
	return a, b, c, d, e, f, g, time.Since(start)
}

// Duration8 returns the time taken to execute a function.
// Play: https://go.dev/play/p/T8kxpG1c5Na
func Duration8[A, B, C, D, E, F, G, H any](cb func() (A, B, C, D, E, F, G, H)) (A, B, C, D, E, F, G, H, time.Duration) {
	start := time.Now()
	a, b, c, d, e, f, g, h := cb()
	return a, b, c, d, e, f, g, h, time.Since(start)
}

// Duration9 returns the time taken to execute a function.
// Play: https://go.dev/play/p/bg9ix2VrZ0j
func Duration9[A, B, C, D, E, F, G, H, I any](cb func() (A, B, C, D, E, F, G, H, I)) (A, B, C, D, E, F, G, H, I, time.Duration) {
	start := time.Now()
	a, b, c, d, e, f, g, h, i := cb()
	return a, b, c, d, e, f, g, h, i, time.Since(start)
}

// Duration10 returns the time taken to execute a function.
// Play: https://go.dev/play/p/Y3n7oJXqJbk
func Duration10[A, B, C, D, E, F, G, H, I, J any](cb func() (A, B, C, D, E, F, G, H, I, J)) (A, B, C, D, E, F, G, H, I, J, time.Duration) {
	start := time.Now()
	a, b, c, d, e, f, g, h, i, j := cb()
	return a, b, c, d, e, f, g, h, i, j, time.Since(start)
}
