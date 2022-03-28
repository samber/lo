package optional

type Value[T any] struct {
	set   bool
	value T
}

func (v Value[T]) IsPresent() bool {
	return v.set
}

func (v Value[T]) IsEmpty() bool {
	return !v.set
}

func (v Value[T]) Get() T {
	return v.value
}

func (v Value[T]) IfPresent(consumer func(T)) {
	if v.set {
		consumer(v.value)
	}
}

func (v Value[T]) IfPresentOrElse(consumer func(T), elseFunction func()) {
	if !v.set {
		elseFunction()
		return
	}

	consumer(v.value)
}

func (v Value[T]) OrElse(elseVal T) T {
	if v.set {
		return v.value
	}

	return elseVal
}

func (v Value[T]) OrElseThrow() {
	if !v.set {
		panic("Optional value was expected to be set but wasn't!")
	}
}

func Of[T any](value T) Value[T] {
	return Value[T]{
		set:   true,
		value: value,
	}
}

func Empty[T any]() Value[T] {
	return Value[T]{
		set: false,
	}
}
