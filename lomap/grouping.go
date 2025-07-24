package lomap

func Group[Slice ~[]T, T any, K comparable](xs Slice, fkey func(T) K) map[K]Slice {
	m := make(map[K]Slice, len(xs))

	for _, x := range xs {
		k := fkey(x)
		m[k] = append(m[k], x)
	}

	return m
}

func IGroup[Slice ~[]T, T any, K comparable](xs Slice, ikey func(int, T) K) map[K]Slice {
	m := make(map[K]Slice, len(xs))

	for i, x := range xs {
		k := ikey(i, x)
		m[k] = append(m[k], x)
	}

	return m
}

func GroupMap[Slice ~[]T, T, V any, K comparable](xs Slice, fmap func(T) (K, V)) map[K][]V {
	m := make(map[K][]V, len(xs))

	for _, x := range xs {
		k, v := fmap(x)
		m[k] = append(m[k], v)
	}

	return m
}

func IGroupMap[Slice ~[]T, T, V any, K comparable](xs Slice, ifmap func(int, T) (K, V)) map[K][]V {
	m := make(map[K][]V, len(xs))

	for i, x := range xs {
		k, v := ifmap(i, x)
		m[k] = append(m[k], v)
	}

	return m
}

func GroupFilter[Slice ~[]T, T any, K comparable](xs Slice, fkey func(T) (K, bool)) map[K]Slice {
	m := make(map[K]Slice, len(xs))

	for _, x := range xs {
		if k, ok := fkey(x); ok {
			m[k] = append(m[k], x)
		}
	}

	return m
}

func IGroupFilter[Slice ~[]T, T any, K comparable](xs Slice, ifkey func(int, T) (K, bool)) map[K]Slice {
	m := make(map[K]Slice, len(xs))

	for i, x := range xs {
		if k, ok := ifkey(i, x); ok {
			m[k] = append(m[k], x)
		}
	}

	return m
}

func GroupFilterMap[Slice ~[]T, T, V any, K comparable](xs Slice, fmap func(T) (K, V, bool)) map[K][]V {
	m := make(map[K][]V, len(xs))

	for _, x := range xs {
		if k, v, ok := fmap(x); ok {
			m[k] = append(m[k], v)
		}
	}

	return m
}

func IGroupFilterMap[Slice ~[]T, T, V any, K comparable](xs Slice, ifmap func(int, T) (K, V, bool)) map[K][]V {
	m := make(map[K][]V, len(xs))

	for i, x := range xs {
		if k, v, ok := ifmap(i, x); ok {
			m[k] = append(m[k], v)
		}
	}

	return m
}
