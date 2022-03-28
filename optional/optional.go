package optional

type Value[T any] struct {
	set   bool
	value T
}

type IfPresent[T any] struct {
	value *Value[T]
}

func (v Value[T]) IfPresent() IfPresent[T] {
	return IfPresent[T]{
		value: &v,
	}
}

func (v Value[T]) IfPresentOrElse(elseVal T) T {
	if v.set {
		return v.value
	}

	return elseVal
}

func (v Value[T]) IsPresent() bool {
	return v.set
}

func (v Value[T]) Get() T {
	return v.value
}

func (p IfPresent[T]) OrElse(elseVal T) T {
	if p.value.set {
		return p.value.value
	}

	return elseVal
}

func (p IfPresent[T]) OrElseThrow() {
	if !p.value.set {
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
