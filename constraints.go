package lo

type Clonable[T any] interface {
	Clone() T
}

type Ordered interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string
}
