package lomap

import (
	"cmp"
	"github.com/samber/lo/lotup"
)

func MinVal[Map ~map[K]V, K comparable, V cmp.Ordered](m Map) V {
	if len(m) == 0 {
		var zero V
		return zero
	}

	res := lotup.Second(Any(m))
	for _, v := range m {
		if v < res {
			res = v
		}
	}

	return res
}
