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

func Min[Slice ~[]T, T any, R cmp.Ordered](xs Slice, frank func(T) R) (t T, r R) {
	if len(xs) == 0 {
		return
	}

	t, r = xs[0], frank(xs[0])
	for _, x := range xs[1:] {
		if v := frank(x); v < r {
			t, r = x, v
		}
	}

	return
}

func ArgMin[Slice ~[]T, T any, R cmp.Ordered](xs Slice, frank func(T) R) (index int, r R) {
	if len(xs) == 0 {
		return 0, r
	}

	index = 0
	r = frank(xs[0])
	for i, x := range xs[1:] {
		if v := frank(x); v < r {
			index, r = i+1, v // +1 because we skipped the first element
		}
	}

	return
}

func IMin[Slice ~[]T, T any, R cmp.Ordered](xs Slice, irank func(int, T) R) (t T, r R) {
	if len(xs) == 0 {
		return
	}

	t, r = xs[0], irank(0, xs[0])
	for i, x := range xs[1:] {
		if v := irank(i+1, x); v < r { // +1 because we skipped the first element
			t, r = x, v
		}
	}

	return
}

func IArgMin[Slice ~[]T, T any, R cmp.Ordered](xs Slice, irank func(int, T) R) (index int, r R) {
	if len(xs) == 0 {
		return 0, r
	}

	index = 0
	r = irank(0, xs[0])
	for i, x := range xs[1:] {
		if v := irank(i+1, x); v < r { // +1 because we skipped the first element
			index, r = i+1, v
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

func RMin[Slice ~[]T, T any, R cmp.Ordered](xs Slice, frank func(T) R) (t T, r R) {
	if len(xs) == 0 {
		return
	}

	last := len(xs) - 1
	t, r = xs[last], frank(xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := frank(xs[i]); v < r {
			t, r = xs[i], v
		}
	}

	return
}

func RArgMin[Slice ~[]T, T any, R cmp.Ordered](xs Slice, frank func(T) R) (index int, r R) {
	if len(xs) == 0 {
		return
	}

	last := len(xs) - 1
	index, r = last, frank(xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := frank(xs[i]); v < r {
			index, r = i, v
		}
	}

	return
}

func IRMin[Slice ~[]T, T any, R cmp.Ordered](xs Slice, frank func(int, T) R) (t T, r R) {
	if len(xs) == 0 {
		return
	}

	last := len(xs) - 1
	t, r = xs[last], frank(last, xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := frank(i, xs[i]); v < r {
			t, r = xs[i], v
		}
	}

	return
}

func IRArgMin[Slice ~[]T, T any, R cmp.Ordered](xs Slice, irank func(int, T) R) (index int, r R) {
	if len(xs) == 0 {
		return
	}

	last := len(xs) - 1
	index, r = last, irank(last, xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := irank(i, xs[i]); v < r {
			index, r = i, v
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

func Max[Slice ~[]T, T any, R cmp.Ordered](xs Slice, frank func(T) R) (t T, r R) {
	if len(xs) == 0 {
		return
	}

	t, r = xs[0], frank(xs[0])
	for _, x := range xs[1:] {
		if v := frank(x); v > r {
			t, r = x, v
		}
	}

	return
}

func ArgMax[Slice ~[]T, T any, R cmp.Ordered](xs Slice, frank func(T) R) (index int, r R) {
	if len(xs) == 0 {
		return 0, r
	}

	index = 0
	r = frank(xs[0])
	for i, x := range xs[1:] {
		if v := frank(x); v > r {
			index, r = i+1, v
		}
	}

	return
}

func IMax[Slice ~[]T, T any, R cmp.Ordered](xs Slice, irank func(int, T) R) (t T, r R) {
	if len(xs) == 0 {
		return
	}

	t, r = xs[0], irank(0, xs[0])
	for i, x := range xs[1:] {
		if v := irank(i+1, x); v > r {
			t, r = x, v
		}
	}

	return
}

func IArgMax[Slice ~[]T, T any, R cmp.Ordered](xs Slice, irank func(int, T) R) (index int, r R) {
	if len(xs) == 0 {
		return 0, r
	}

	index = 0
	r = irank(0, xs[0])
	for i, x := range xs[1:] {
		if v := irank(i+1, x); v > r {
			index, r = i+1, v
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

func RMax[Slice ~[]T, T any, R cmp.Ordered](xs Slice, frank func(T) R) (t T, r R) {
	if len(xs) == 0 {
		return
	}

	last := len(xs) - 1
	t, r = xs[last], frank(xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := frank(xs[i]); v > r {
			t, r = xs[i], v
		}
	}

	return
}

func RArgMax[Slice ~[]T, T any, R cmp.Ordered](xs Slice, frank func(T) R) (index int, r R) {
	if len(xs) == 0 {
		return
	}

	last := len(xs) - 1
	index, r = last, frank(xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := frank(xs[i]); v > r {
			index, r = i, v
		}
	}

	return
}

func IRMax[Slice ~[]T, T any, R cmp.Ordered](xs Slice, irank func(int, T) R) (t T, r R) {
	if len(xs) == 0 {
		return
	}

	last := len(xs) - 1
	t, r = xs[last], irank(last, xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := irank(i, xs[i]); v > r {
			t, r = xs[i], v
		}
	}

	return
}

func IRArgMax[Slice ~[]T, T any, R cmp.Ordered](xs Slice, irank func(int, T) R) (index int, r R) {
	if len(xs) == 0 {
		return
	}

	last := len(xs) - 1
	index, r = last, irank(last, xs[last])
	for i := last - 1; i >= 0; i-- {
		if v := irank(i, xs[i]); v > r {
			index, r = i, v
		}
	}

	return
}

// MinCmp forward min using fcmp
func MinCmp[Slice ~[]T, T any](xs Slice, fcmp func(a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}

	res := xs[0]
	for _, x := range xs[1:] {
		if fcmp(x, res) < 0 {
			res = x
		}
	}

	return res
}

// MaxCmp forward max using fcmp
func MaxCmp[Slice ~[]T, T any](xs Slice, fcmp func(a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}

	res := xs[0]
	for _, x := range xs[1:] {
		if fcmp(x, res) > 0 {
			res = x
		}
	}

	return res
}

// RMinCmp reverse min using fcmp
func RMinCmp[Slice ~[]T, T any](xs Slice, fcmp func(a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}

	last := len(xs) - 1
	res := xs[last]
	for i := last - 1; i >= 0; i-- {
		if fcmp(xs[i], res) < 0 {
			res = xs[i]
		}
	}

	return res
}

// RMaxCmp reverse max using fcmp
func RMaxCmp[Slice ~[]T, T any](xs Slice, fcmp func(a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}

	last := len(xs) - 1
	res := xs[last]
	for i := last - 1; i >= 0; i-- {
		if fcmp(xs[i], res) > 0 {
			res = xs[i]
		}
	}

	return res
}

// IMinCmp forward min using index-aware icmp
func IMinCmp[Slice ~[]T, T any](xs Slice, icmp func(i, j int, a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}

	index := 0
	res := xs[0]
	for i, x := range xs[1:] {
		if icmp(i+1, index, x, res) < 0 {
			index = i + 1 // +1 because we skipped the first element
			res = x
		}
	}

	return res
}

// IMaxCmp forward max using index-aware icmp
func IMaxCmp[Slice ~[]T, T any](xs Slice, icmp func(i, j int, a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}

	index := 0
	res := xs[0]
	for i, x := range xs[1:] {
		if icmp(i+1, index, x, res) > 0 {
			index = i + 1 // +1 because we skipped the first element
			res = x
		}
	}

	return res
}

// IRMinCmp reverse min using index-aware icmp
func IRMinCmp[Slice ~[]T, T any](xs Slice, icmp func(i, j int, a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}

	index := len(xs) - 1
	res := xs[index]
	for i := index - 1; i >= 0; i-- {
		if icmp(i, index, xs[i], res) < 0 {
			index = i
			res = xs[i]
		}
	}

	return res
}

// IRMaxCmp reverse max using index-aware icmp
func IRMaxCmp[Slice ~[]T, T any](xs Slice, icmp func(i, j int, a, b T) int) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}

	index := len(xs) - 1
	res := xs[index]
	for i := index - 1; i >= 0; i-- {
		if icmp(i, index, xs[i], res) > 0 {
			index = i
			res = xs[i]
		}
	}

	return res
}

// ArgMinCmp forward min index using fcmp
func ArgMinCmp[Slice ~[]T, T any](xs Slice, fcmp func(a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}

	index := 0
	for i, x := range xs[1:] {
		if fcmp(x, xs[index]) < 0 {
			index = i + 1
		}
	}

	return index
}

// ArgMaxCmp forward max index using fcmp
func ArgMaxCmp[Slice ~[]T, T any](xs Slice, fcmp func(a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}

	index := 0
	for i, x := range xs[1:] {
		if fcmp(x, xs[index]) > 0 {
			index = i + 1
		}
	}

	return index
}

// RArgMinCmp reverse min index using fcmp
func RArgMinCmp[Slice ~[]T, T any](xs Slice, fcmp func(a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}

	last := len(xs) - 1
	index := last
	for i := last - 1; i >= 0; i-- {
		if fcmp(xs[i], xs[index]) < 0 {
			index = i
		}
	}

	return index
}

// RArgMaxCmp reverse max index using fcmp
func RArgMaxCmp[Slice ~[]T, T any](xs Slice, fcmp func(a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}

	last := len(xs) - 1
	index := last
	for i := last - 1; i >= 0; i-- {
		if fcmp(xs[i], xs[index]) > 0 {
			index = i
		}
	}

	return index
}

// IArgMinCmp forward min index using index-aware icmp
func IArgMinCmp[Slice ~[]T, T any](xs Slice, icmp func(i, j int, a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}

	index := 0
	for i, x := range xs[1:] {
		if icmp(i+1, index, x, xs[index]) < 0 {
			index = i + 1
		}
	}

	return index
}

// IArgMaxCmp forward max index using index-aware icmp
func IArgMaxCmp[Slice ~[]T, T any](xs Slice, icmp func(i, j int, a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}

	index := 0
	for i, x := range xs[1:] {
		if icmp(i+1, index, x, xs[index]) > 0 {
			index = i + 1
		}
	}

	return index
}

// IRArgMinCmp reverse min index using index-aware icmp
func IRArgMinCmp[Slice ~[]T, T any](xs Slice, icmp func(i, j int, a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}

	last := len(xs) - 1
	index := last
	for i := last - 1; i >= 0; i-- {
		if icmp(i, index, xs[i], xs[index]) < 0 {
			index = i
		}
	}

	return index
}

// IRArgMaxCmp reverse max index using index-aware icmp
func IRArgMaxCmp[Slice ~[]T, T any](xs Slice, icmp func(i, j int, a, b T) int) int {
	if len(xs) == 0 {
		return 0
	}

	last := len(xs) - 1
	index := last
	for i := last - 1; i >= 0; i-- {
		if icmp(i, index, xs[i], xs[index]) > 0 {
			index = i
		}
	}

	return index
}
