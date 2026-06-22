package lotype

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func CastNum[R, T Number](x T) R {
	return R(x)
}

func CastStr[R, T ~string](x T) R {
	return R(x)
}
