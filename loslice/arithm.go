package loslice

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

func Sum[Slice ~[]T, T Number](xs Slice) (sum T) {
	for _, x := range xs {
		sum += x
	}

	return sum
}

func Product[Slice ~[]T, T Number](xs Slice) (prod T) {
	prod = 1
	for _, x := range xs {
		prod *= x
	}

	return prod
}

func FilterSum[Slice ~[]T, T Number](xs Slice, pred func(T) bool) (sum T) {
	for _, x := range xs {
		if pred(x) {
			sum += x
		}
	}

	return sum
}

func FilterProduct[Slice ~[]T, T Number](xs Slice, pred func(T) bool) (prod T) {
	prod = 1
	for _, x := range xs {
		if pred(x) {
			prod *= x
		}
	}

	return prod
}

func IFilterSum[Slice ~[]T, T Number](xs Slice, ipred func(int, T) bool) (sum T) {
	for i, x := range xs {
		if ipred(i, x) {
			sum += x
		}
	}

	return sum
}

func IFilterProduct[Slice ~[]T, T Number](xs Slice, ipred func(int, T) bool) (prod T) {
	prod = 1
	for i, x := range xs {
		if ipred(i, x) {
			prod *= x
		}
	}

	return prod
}

func MapSum[Slice ~[]T, T any, V Number](xs Slice, fmap func(T) V) (sum V) {
	for _, x := range xs {
		sum += fmap(x)
	}

	return sum
}

func MapProduct[Slice ~[]T, T any, V Number](xs Slice, fmap func(T) V) (prod V) {
	prod = 1
	for _, x := range xs {
		prod *= fmap(x)
	}

	return prod
}

func IMapSum[Slice ~[]T, T any, V Number](xs Slice, imap func(int, T) V) (sum V) {
	for i, x := range xs {
		sum += imap(i, x)
	}

	return sum
}

func IMapProduct[Slice ~[]T, T any, V Number](xs Slice, imap func(int, T) V) (prod V) {
	prod = 1
	for i, x := range xs {
		prod *= imap(i, x)
	}

	return prod
}

func FilterMapSum[Slice ~[]T, T any, V Number](xs Slice, fmap func(T) (V, bool)) (sum V) {
	for _, x := range xs {
		if v, ok := fmap(x); ok {
			sum += v
		}
	}

	return sum
}

func FilterMapProduct[Slice ~[]T, T any, V Number](xs Slice, fmap func(T) (V, bool)) (prod V) {
	prod = 1
	for _, x := range xs {
		if v, ok := fmap(x); ok {
			prod *= v
		}
	}

	return prod
}

func IFilterMapSum[Slice ~[]T, T any, V Number](xs Slice, imap func(int, T) (V, bool)) (sum V) {
	for i, x := range xs {
		if v, ok := imap(i, x); ok {
			sum += v
		}
	}

	return sum
}

func IFilterMapProduct[Slice ~[]T, T any, V Number](xs Slice, imap func(int, T) (V, bool)) (prod V) {
	prod = 1
	for i, x := range xs {
		if v, ok := imap(i, x); ok {
			prod *= v
		}
	}

	return prod
}
