package loslice

import "cmp"

func MinVal[Slice ~[]T, T cmp.Ordered](xs Slice) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}

	res := xs[0]
	for _, x := range xs[1:] {
		if x < res {
			res = x
		}
	}

	return res
}

func ArgMinVal[Slice ~[]T, T cmp.Ordered](xs Slice) int {
	if len(xs) == 0 {
		return 0
	}

	index := 0
	for i, x := range xs[1:] {
		if x < xs[index] {
			index = i + 1
		}
	}

	return index
}

func Min[Slice ~[]T, T any, K cmp.Ordered](xs Slice, fkey func(T) K) (t T, k K) {
	if len(xs) == 0 {
		return
	}

	t, k = xs[0], fkey(xs[0])
	for _, x := range xs[1:] {
		if v := fkey(x); v < k {
			t, k = x, v
		}
	}

	return
}

func ArgMin[Slice ~[]T, T any, K cmp.Ordered](xs Slice, fkey func(T) K) (index int, k K) {
	if len(xs) == 0 {
		return 0, k
	}

	index = 0
	k = fkey(xs[0])
	for i, x := range xs[1:] {
		if v := fkey(x); v < k {
			index, k = i+1, v // +1 because we skipped the first element
		}
	}

	return
}

func IMin[Slice ~[]T, T any, K cmp.Ordered](xs Slice, ikey func(int, T) K) (t T, k K) {
	if len(xs) == 0 {
		return
	}

	t, k = xs[0], ikey(0, xs[0])
	for i, x := range xs[1:] {
		if v := ikey(i+1, x); v < k { // +1 because we skipped the first element
			t, k = x, v
		}
	}

	return
}

func IArgMin[Slice ~[]T, T any, K cmp.Ordered](xs Slice, ikey func(int, T) K) (index int, k K) {
	if len(xs) == 0 {
		return 0, k
	}

	index = 0
	k = ikey(0, xs[0])
	for i, x := range xs[1:] {
		if v := ikey(i+1, x); v < k { // +1 because we skipped the first element
			index, k = i+1, v
		}
	}

	return
}

func RMinVal[Slice ~[]T, T cmp.Ordered](xs Slice) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}

	last := len(xs) - 1
	res := xs[last]
	for i := last - 1; i >= 0; i-- {
		if x := xs[i]; x < res {
			res = x
		}
	}

	return res
}

func RArgMinVal[Slice ~[]T, T cmp.Ordered](xs Slice) int {
	if len(xs) == 0 {
		return 0
	}

	last := len(xs) - 1
	index := last
	for i := last - 1; i >= 0; i-- {
		if xs[i] < xs[index] {
			index = i
		}
	}

	return index
}

func RMin[Slice ~[]T, T any, K cmp.Ordered](xs Slice, fkey func(T) K) (t T, k K) {
	if len(xs) == 0 {
		return
	}

	last := len(xs) - 1
	t, k = xs[last], fkey(xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := fkey(xs[i]); v < k {
			t, k = xs[i], v
		}
	}

	return
}

func RArgMin[Slice ~[]T, T any, K cmp.Ordered](xs Slice, fkey func(T) K) (index int, k K) {
	if len(xs) == 0 {
		return
	}

	last := len(xs) - 1
	index, k = last, fkey(xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := fkey(xs[i]); v < k {
			index, k = i, v
		}
	}

	return
}

func IRMin[Slice ~[]T, T any, K cmp.Ordered](xs Slice, fkey func(int, T) K) (t T, k K) {
	if len(xs) == 0 {
		return
	}

	last := len(xs) - 1
	t, k = xs[last], fkey(last, xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := fkey(i, xs[i]); v < k {
			t, k = xs[i], v
		}
	}

	return
}

func IRArgMin[Slice ~[]T, T any, K cmp.Ordered](xs Slice, ikey func(int, T) K) (index int, k K) {
	if len(xs) == 0 {
		return
	}

	last := len(xs) - 1
	index, k = last, ikey(last, xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := ikey(i, xs[i]); v < k {
			index, k = i, v
		}
	}

	return
}

func MaxVal[Slice ~[]T, T cmp.Ordered](xs Slice) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}
	res := xs[0]
	for _, x := range xs[1:] {
		if x > res {
			res = x
		}
	}
	return res
}

func ArgMaxVal[Slice ~[]T, T cmp.Ordered](xs Slice) int {
	if len(xs) == 0 {
		return 0
	}
	index := 0
	for i, x := range xs[1:] {
		if x > xs[index] {
			index = i + 1
		}
	}
	return index
}

func Max[Slice ~[]T, T any, K cmp.Ordered](xs Slice, fkey func(T) K) (t T, k K) {
	if len(xs) == 0 {
		return
	}
	t, k = xs[0], fkey(xs[0])
	for _, x := range xs[1:] {
		if v := fkey(x); v > k {
			t, k = x, v
		}
	}
	return
}

func ArgMax[Slice ~[]T, T any, K cmp.Ordered](xs Slice, fkey func(T) K) (index int, k K) {
	if len(xs) == 0 {
		return 0, k
	}
	index = 0
	k = fkey(xs[0])
	for i, x := range xs[1:] {
		if v := fkey(x); v > k {
			index, k = i+1, v
		}
	}
	return
}

func IMax[Slice ~[]T, T any, K cmp.Ordered](xs Slice, ikey func(int, T) K) (t T, k K) {
	if len(xs) == 0 {
		return
	}
	t, k = xs[0], ikey(0, xs[0])
	for i, x := range xs[1:] {
		if v := ikey(i+1, x); v > k {
			t, k = x, v
		}
	}
	return
}

func IArgMax[Slice ~[]T, T any, K cmp.Ordered](xs Slice, ikey func(int, T) K) (index int, k K) {
	if len(xs) == 0 {
		return 0, k
	}
	index = 0
	k = ikey(0, xs[0])
	for i, x := range xs[1:] {
		if v := ikey(i+1, x); v > k {
			index, k = i+1, v
		}
	}
	return
}

func RMaxVal[Slice ~[]T, T cmp.Ordered](xs Slice) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}
	last := len(xs) - 1
	res := xs[last]
	for i := last - 1; i >= 0; i-- {
		if x := xs[i]; x > res {
			res = x
		}
	}
	return res
}

func RArgMaxVal[Slice ~[]T, T cmp.Ordered](xs Slice) int {
	if len(xs) == 0 {
		return 0
	}
	last := len(xs) - 1
	index := last
	for i := last - 1; i >= 0; i-- {
		if xs[i] > xs[index] {
			index = i
		}
	}
	return index
}

func RMax[Slice ~[]T, T any, K cmp.Ordered](xs Slice, fkey func(T) K) (t T, k K) {
	if len(xs) == 0 {
		return
	}
	last := len(xs) - 1
	t, k = xs[last], fkey(xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := fkey(xs[i]); v > k {
			t, k = xs[i], v
		}
	}
	return
}

func RArgMax[Slice ~[]T, T any, K cmp.Ordered](xs Slice, fkey func(T) K) (index int, k K) {
	if len(xs) == 0 {
		return
	}
	last := len(xs) - 1
	index, k = last, fkey(xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := fkey(xs[i]); v > k {
			index, k = i, v
		}
	}
	return
}

func IRMax[Slice ~[]T, T any, K cmp.Ordered](xs Slice, ikey func(int, T) K) (t T, k K) {
	if len(xs) == 0 {
		return
	}
	last := len(xs) - 1
	t, k = xs[last], ikey(last, xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := ikey(i, xs[i]); v > k {
			t, k = xs[i], v
		}
	}
	return
}

func IRArgMax[Slice ~[]T, T any, K cmp.Ordered](xs Slice, ikey func(int, T) K) (index int, k K) {
	if len(xs) == 0 {
		return
	}
	last := len(xs) - 1
	index, k = last, ikey(last, xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := ikey(i, xs[i]); v > k {
			index, k = i, v
		}
	}
	return
}

// MinCmp forward min using cmpFn
func MinCmp[Slice ~[]T, T any](xs Slice, cmpFn func(a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}
	res := xs[0]
	for _, x := range xs[1:] {
		if cmpFn(x, res) < 0 {
			res = x
		}
	}
	return res
}

// MaxCmp forward max using cmpFn
func MaxCmp[Slice ~[]T, T any](xs Slice, cmpFn func(a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}
	res := xs[0]
	for _, x := range xs[1:] {
		if cmpFn(x, res) > 0 {
			res = x
		}
	}
	return res
}

// RMinCmp reverse min using cmpFn
func RMinCmp[Slice ~[]T, T any](xs Slice, cmpFn func(a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}
	last := len(xs) - 1
	res := xs[last]
	for i := last - 1; i >= 0; i-- {
		if cmpFn(xs[i], res) < 0 {
			res = xs[i]
		}
	}
	return res
}

// RMaxCmp reverse max using cmpFn
func RMaxCmp[Slice ~[]T, T any](xs Slice, cmpFn func(a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}
	last := len(xs) - 1
	res := xs[last]
	for i := last - 1; i >= 0; i-- {
		if cmpFn(xs[i], res) > 0 {
			res = xs[i]
		}
	}
	return res
}

// IMinCmp forward min using index-aware cmpFn
func IMinCmp[Slice ~[]T, T any](xs Slice, cmpFn func(i, j int, a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}
	index := 0
	res := xs[0]
	for i, x := range xs[1:] {
		if cmpFn(i+1, index, x, res) < 0 {
			index = i + 1 // +1 because we skipped the first element
			res = x
		}
	}
	return res
}

// IMaxCmp forward max using index-aware cmpFn
func IMaxCmp[Slice ~[]T, T any](xs Slice, cmpFn func(i, j int, a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}

	index := 0
	res := xs[0]
	for i, x := range xs[1:] {
		if cmpFn(i+1, index, x, res) > 0 {
			index = i + 1 // +1 because we skipped the first element
			res = x
		}
	}
	return res
}

// IRMinCmp reverse min using index-aware cmpFn
func IRMinCmp[Slice ~[]T, T any](xs Slice, cmpFn func(i, j int, a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}
	index := len(xs) - 1
	res := xs[index]
	for i := index - 1; i >= 0; i-- {
		if cmpFn(i, index, xs[i], res) < 0 {
			index = i
			res = xs[i]
		}
	}
	return res
}

// IRMaxCmp reverse max using index-aware cmpFn
func IRMaxCmp[Slice ~[]T, T any](xs Slice, cmpFn func(i, j int, a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}
	index := len(xs) - 1
	res := xs[index]
	for i := index - 1; i >= 0; i-- {
		if cmpFn(i, index, xs[i], res) > 0 {
			index = i
			res = xs[i]
		}
	}
	return res
}

// ArgMinCmp forward min index using cmpFn
func ArgMinCmp[Slice ~[]T, T any](xs Slice, cmpFn func(a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}
	index := 0
	for i, x := range xs[1:] {
		if cmpFn(x, xs[index]) < 0 {
			index = i + 1
		}
	}
	return index
}

// ArgMaxCmp forward max index using cmpFn
func ArgMaxCmp[Slice ~[]T, T any](xs Slice, cmpFn func(a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}
	index := 0
	for i, x := range xs[1:] {
		if cmpFn(x, xs[index]) > 0 {
			index = i + 1
		}
	}
	return index
}

// RArgMinCmp reverse min index using cmpFn
func RArgMinCmp[Slice ~[]T, T any](xs Slice, cmpFn func(a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}
	last := len(xs) - 1
	index := last
	for i := last - 1; i >= 0; i-- {
		if cmpFn(xs[i], xs[index]) < 0 {
			index = i
		}
	}
	return index
}

// RArgMaxCmp reverse max index using cmpFn
func RArgMaxCmp[Slice ~[]T, T any](xs Slice, cmpFn func(a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}
	last := len(xs) - 1
	index := last
	for i := last - 1; i >= 0; i-- {
		if cmpFn(xs[i], xs[index]) > 0 {
			index = i
		}
	}
	return index
}

// IArgMinCmp forward min index using index-aware cmpFn
func IArgMinCmp[Slice ~[]T, T any](xs Slice, cmpFn func(i, j int, a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}
	index := 0
	for i, x := range xs[1:] {
		if cmpFn(i+1, index, x, xs[index]) < 0 {
			index = i + 1
		}
	}
	return index
}

// IArgMaxCmp forward max index using index-aware cmpFn
func IArgMaxCmp[Slice ~[]T, T any](xs Slice, cmpFn func(i, j int, a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}
	index := 0
	for i, x := range xs[1:] {
		if cmpFn(i+1, index, x, xs[index]) > 0 {
			index = i + 1
		}
	}
	return index
}

// IRArgMinCmp reverse min index using index-aware cmpFn
func IRArgMinCmp[Slice ~[]T, T any](xs Slice, cmpFn func(i, j int, a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}
	last := len(xs) - 1
	index := last
	for i := last - 1; i >= 0; i-- {
		if cmpFn(i, index, xs[i], xs[index]) < 0 {
			index = i
		}
	}
	return index
}

// IRArgMaxCmp reverse max index using index-aware cmpFn
func IRArgMaxCmp[Slice ~[]T, T any](xs Slice, cmpFn func(i, j int, a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}
	last := len(xs) - 1
	index := last
	for i := last - 1; i >= 0; i-- {
		if cmpFn(i, index, xs[i], xs[index]) > 0 {
			index = i
		}
	}
	return index
}
