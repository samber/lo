package lomap

import (
	"cmp"
	"github.com/samber/lo/lotup"
)

func MinKey[Map ~map[K]V, K cmp.Ordered, V any](m Map) K {
	if len(m) == 0 {
		var zero K
		return zero
	}

	res := lotup.First(Any(m))
	for k := range m {
		if k < res {
			res = k
		}
	}

	return res
}

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

func ArgMinVal[Map ~map[K]V, K comparable, V cmp.Ordered](m Map) K {
	if len(m) == 0 {
		var zero K
		return zero
	}

	key, res := Any(m)
	for k, v := range m {
		if v < res {
			key, res = k, v
		}
	}

	return key
}

func Min[Map ~map[K]V, K comparable, V any, R cmp.Ordered](m Map, frank func(V) R) (V, R) {
	if len(m) == 0 {
		var zero V
		var zeroR R
		return zero, zeroR
	}

	val := lotup.Second(Any(m))
	rank := frank(val)
	for _, v := range m {
		r := frank(v)
		if r < rank {
			val, rank = v, r
		}
	}

	return val, rank
}

func ArgMin[Map ~map[K]V, K comparable, V any, R cmp.Ordered](m Map, frank func(V) R) (K, R) {
	if len(m) == 0 {
		var zero K
		var zeroR R
		return zero, zeroR
	}

	key, v := Any(m)
	rank := frank(v)
	for k, v := range m {
		r := frank(v)
		if r < rank {
			key, rank = k, r
		}
	}

	return key, rank
}

func KMin[Map ~map[K]V, K comparable, V any, R cmp.Ordered](m Map, krank func(K, V) R) (V, R) {
	if len(m) == 0 {
		var zero V
		var zeroR R
		return zero, zeroR
	}

	k, val := Any(m)
	rank := krank(k, val)
	for k, v := range m {
		r := krank(k, v)
		if r < rank {
			val, rank = v, r
		}
	}

	return val, rank
}

func KArgMin[Map ~map[K]V, K comparable, V any, R cmp.Ordered](m Map, krank func(K, V) R) (K, R) {
	if len(m) == 0 {
		var zero K
		var zeroR R
		return zero, zeroR
	}

	key, v := Any(m)
	rank := krank(key, v)
	for k, v := range m {
		r := krank(k, v)
		if r < rank {
			key, rank = k, r
		}
	}

	return key, rank
}

func MaxKey[Map ~map[K]V, K cmp.Ordered, V any](m Map) K {
	if len(m) == 0 {
		var zero K
		return zero
	}

	res := lotup.First(Any(m))
	for k := range m {
		if k > res {
			res = k
		}
	}

	return res
}

func MaxVal[Map ~map[K]V, K comparable, V cmp.Ordered](m Map) V {
	if len(m) == 0 {
		var zero V
		return zero
	}

	res := lotup.Second(Any(m))
	for _, v := range m {
		if v > res {
			res = v
		}
	}

	return res
}

func ArgMaxVal[Map ~map[K]V, K comparable, V cmp.Ordered](m Map) K {
	if len(m) == 0 {
		var zero K
		return zero
	}

	key, res := Any(m)
	for k, v := range m {
		if v > res {
			key, res = k, v
		}
	}

	return key
}

func Max[Map ~map[K]V, K comparable, V any, R cmp.Ordered](m Map, frank func(V) R) (V, R) {
	if len(m) == 0 {
		var zero V
		var zeroR R
		return zero, zeroR
	}

	val := lotup.Second(Any(m))
	rank := frank(val)
	for _, v := range m {
		r := frank(v)
		if r > rank {
			val, rank = v, r
		}
	}

	return val, rank
}

func ArgMax[Map ~map[K]V, K comparable, V any, R cmp.Ordered](m Map, frank func(V) R) (K, R) {
	if len(m) == 0 {
		var zero K
		var zeroR R
		return zero, zeroR
	}

	key, v := Any(m)
	rank := frank(v)
	for k, v := range m {
		r := frank(v)
		if r > rank {
			key, rank = k, r
		}
	}

	return key, rank
}

func KMax[Map ~map[K]V, K comparable, V any, R cmp.Ordered](m Map, krank func(K, V) R) (V, R) {
	if len(m) == 0 {
		var zero V
		var zeroR R
		return zero, zeroR
	}

	k, val := Any(m)
	rank := krank(k, val)
	for k, v := range m {
		r := krank(k, v)
		if r > rank {
			val, rank = v, r
		}
	}

	return val, rank
}

func KArgMax[Map ~map[K]V, K comparable, V any, R cmp.Ordered](m Map, krank func(K, V) R) (K, R) {
	if len(m) == 0 {
		var zero K
		var zeroR R
		return zero, zeroR
	}

	key, v := Any(m)
	rank := krank(key, v)
	for k, v := range m {
		r := krank(k, v)
		if r > rank {
			key, rank = k, r
		}
	}

	return key, rank
}
