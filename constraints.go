package lo

type Clonable[T any] interface {
	Clone() T
}
