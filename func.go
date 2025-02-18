package lo

// Func is a function with 0 argument and 1 return value (null-ary function).
type Func[R any] func() R

// Func1 is a function with 1 argument and 1 return value (unary function),
// which supports partial application.
type Func1[T1, R any] func(T1) R

// MFunc1 casts function f with 1 argument and 1 return value to Func1.
func MFunc1[T1, R any](f Func1[T1, R]) Func1[T1, R] {
	return f
}

// Partial binds the first argument (from left to right) of Func1 f to value t1,
// producing Func of smaller arity.
func (f Func1[T1, R]) Partial(t1 T1) Func[R] {
	return func() R {
		return f(t1)
	}
}

// PartialR binds the first argument (from right to left) of Func1 f to value t1,
// producing Func of smaller arity.
func (f Func1[T1, R]) PartialR(t1 T1) Func[R] {
	return func() R {
		return f(t1)
	}
}

// Func2 is a function with 2 arguments and 1 return value (2-ary function),
// which supports partial application.
type Func2[T1, T2, R any] func(T1, T2) R

// MFunc2 casts function f with 2 arguments and 1 return value to Func2
func MFunc2[T1, T2, R any](f Func2[T1, T2, R]) Func2[T1, T2, R] {
	return f
}

// Partial binds the first argument (from left to right) of Func2 f to value t1,
// producing Func1 of smaller arity.
func (f Func2[T1, T2, R]) Partial(t1 T1) Func1[T2, R] {
	return func(t2 T2) R {
		return f(t1, t2)
	}
}

// PartialR binds the first argument (from right to left) of Func2 f to value t2,
// producing Func1 of smaller arity.
func (f Func2[T1, T2, R]) PartialR(t2 T2) Func1[T1, R] {
	return func(t1 T1) R {
		return f(t1, t2)
	}
}

// Func3 is a function with 3 arguments and 1 return value (3-ary function),
// which supports partial application.
type Func3[T1, T2, T3, R any] func(T1, T2, T3) R

// MFunc3 casts function f with 3 arguments and 1 return value to Func3
func MFunc3[T1, T2, T3, R any](f Func3[T1, T2, T3, R]) Func3[T1, T2, T3, R] {
	return f
}

// Partial binds the first argument (from left to right) of Func3 f to value t1,
// producing Func2 of smaller arity.
func (f Func3[T1, T2, T3, R]) Partial(t1 T1) Func2[T2, T3, R] {
	return func(t2 T2, t3 T3) R {
		return f(t1, t2, t3)
	}
}

// PartialR binds the first argument (from right to left) of Func3 f to value t3,
// producing Func2 of smaller arity.
func (f Func3[T1, T2, T3, R]) PartialR(t3 T3) Func2[T1, T2, R] {
	return func(t1 T1, t2 T2) R {
		return f(t1, t2, t3)
	}
}

// Func4 is a function with 4 arguments and 1 return value (4-ary function),
// which supports partial application.
type Func4[T1, T2, T3, T4, R any] func(T1, T2, T3, T4) R

// MFunc4 casts function f with 4 arguments and 1 return value to Func4
func MFunc4[T1, T2, T3, T4, R any](f Func4[T1, T2, T3, T4, R]) Func4[T1, T2, T3, T4, R] {
	return f
}

// Partial binds the first argument (from left to right) of Func4 f to value t1,
// producing Func3 of smaller arity.
func (f Func4[T1, T2, T3, T4, R]) Partial(t1 T1) Func3[T2, T3, T4, R] {
	return func(t2 T2, t3 T3, t4 T4) R {
		return f(t1, t2, t3, t4)
	}
}

// PartialR binds the first argument (from right to left) of Func4 f to value t4,
// producing Func3 of smaller arity.
func (f Func4[T1, T2, T3, T4, R]) PartialR(t4 T4) Func3[T1, T2, T3, R] {
	return func(t1 T1, t2 T2, t3 T3) R {
		return f(t1, t2, t3, t4)
	}
}

// Func5 is a function with 5 arguments and 1 return value (5-ary function),
// which supports partial application.
type Func5[T1, T2, T3, T4, T5, R any] func(T1, T2, T3, T4, T5) R

// MFunc5 casts function f with 5 arguments and 1 return value to Func5
func MFunc5[T1, T2, T3, T4, T5, R any](f Func5[T1, T2, T3, T4, T5, R]) Func5[T1, T2, T3, T4, T5, R] {
	return f
}

// Partial binds the first argument (from left to right) of Func5 f to value t1,
// producing Func4 of smaller arity.
func (f Func5[T1, T2, T3, T4, T5, R]) Partial(t1 T1) Func4[T2, T3, T4, T5, R] {
	return func(t2 T2, t3 T3, t4 T4, t5 T5) R {
		return f(t1, t2, t3, t4, t5)
	}
}

// PartialR binds the first argument (from right to left) of Func5 f to value t5,
// producing Func4 of smaller arity.
func (f Func5[T1, T2, T3, T4, T5, R]) PartialR(t5 T5) Func4[T1, T2, T3, T4, R] {
	return func(t1 T1, t2 T2, t3 T3, t4 T4) R {
		return f(t1, t2, t3, t4, t5)
	}
}

// Func6 is a function with 6 arguments and 1 return value (6-ary function),
// which supports partial application.
type Func6[T1, T2, T3, T4, T5, T6, R any] func(T1, T2, T3, T4, T5, T6) R

// MFunc6 casts function f with 6 arguments and 1 return value to Func6
func MFunc6[T1, T2, T3, T4, T5, T6, R any](f Func6[T1, T2, T3, T4, T5, T6, R]) Func6[T1, T2, T3, T4, T5, T6, R] {
	return f
}

// Partial binds the first argument (from left to right) of Func6 f to value t1,
// producing Func5 of smaller arity.
func (f Func6[T1, T2, T3, T4, T5, T6, R]) Partial(t1 T1) Func5[T2, T3, T4, T5, T6, R] {
	return func(t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) R {
		return f(t1, t2, t3, t4, t5, t6)
	}
}

// PartialR binds the first argument (from right to left) of Func6 f to value t6,
// producing Func5 of smaller arity.
func (f Func6[T1, T2, T3, T4, T5, T6, R]) PartialR(t6 T6) Func5[T1, T2, T3, T4, T5, R] {
	return func(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5) R {
		return f(t1, t2, t3, t4, t5, t6)
	}
}

// Func7 is a function with 7 arguments and 1 return value (7-ary function),
// which supports partial application.
type Func7[T1, T2, T3, T4, T5, T6, T7, R any] func(T1, T2, T3, T4, T5, T6, T7) R

// MFunc7 casts function f with 7 arguments and 1 return value to Func7
func MFunc7[T1, T2, T3, T4, T5, T6, T7, R any](f Func7[T1, T2, T3, T4, T5, T6, T7, R]) Func7[T1, T2, T3, T4, T5, T6, T7, R] {
	return f
}

// Partial binds the first argument (from left to right) of Func7 f to value t1,
// producing Func6 of smaller arity.
func (f Func7[T1, T2, T3, T4, T5, T6, T7, R]) Partial(t1 T1) Func6[T2, T3, T4, T5, T6, T7, R] {
	return func(t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7) R {
		return f(t1, t2, t3, t4, t5, t6, t7)
	}
}

// PartialR binds the first argument (from right to left) of Func7 f to value t7,
// producing Func6 of smaller arity.
func (f Func7[T1, T2, T3, T4, T5, T6, T7, R]) PartialR(t7 T7) Func6[T1, T2, T3, T4, T5, T6, R] {
	return func(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) R {
		return f(t1, t2, t3, t4, t5, t6, t7)
	}
}

// Func8 is a function with 8 arguments and 1 return value (8-ary function),
// which supports partial application.
type Func8[T1, T2, T3, T4, T5, T6, T7, T8, R any] func(T1, T2, T3, T4, T5, T6, T7, T8) R

// MFunc8 casts function f with 8 arguments and 1 return value to Func8
func MFunc8[T1, T2, T3, T4, T5, T6, T7, T8, R any](f Func8[T1, T2, T3, T4, T5, T6, T7, T8, R]) Func8[T1, T2, T3, T4, T5, T6, T7, T8, R] {
	return f
}

// Partial binds the first argument (from left to right) of Func8 f to value t1,
// producing Func7 of smaller arity.
func (f Func8[T1, T2, T3, T4, T5, T6, T7, T8, R]) Partial(t1 T1) Func7[T2, T3, T4, T5, T6, T7, T8, R] {
	return func(t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8) R {
		return f(t1, t2, t3, t4, t5, t6, t7, t8)
	}
}

// PartialR binds the first argument (from right to left) of Func8 f to value t8,
// producing Func7 of smaller arity.
func (f Func8[T1, T2, T3, T4, T5, T6, T7, T8, R]) PartialR(t8 T8) Func7[T1, T2, T3, T4, T5, T6, T7, R] {
	return func(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7) R {
		return f(t1, t2, t3, t4, t5, t6, t7, t8)
	}
}

// Func9 is a function with 9 arguments and 1 return value (9-ary function),
// which supports partial application.
type Func9[T1, T2, T3, T4, T5, T6, T7, T8, T9, R any] func(T1, T2, T3, T4, T5, T6, T7, T8, T9) R

// MFunc9 casts function f with 9 arguments and 1 return value to Func9
func MFunc9[T1, T2, T3, T4, T5, T6, T7, T8, T9, R any](f Func9[T1, T2, T3, T4, T5, T6, T7, T8, T9, R]) Func9[T1, T2, T3, T4, T5, T6, T7, T8, T9, R] {
	return f
}

// Partial binds the first argument (from left to right) of Func9 f to value t1,
// producing Func8 of smaller arity.
func (f Func9[T1, T2, T3, T4, T5, T6, T7, T8, T9, R]) Partial(t1 T1) Func8[T2, T3, T4, T5, T6, T7, T8, T9, R] {
	return func(t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9) R {
		return f(t1, t2, t3, t4, t5, t6, t7, t8, t9)
	}
}

// PartialR binds the first argument (from right to left) of Func9 f to value t9,
// producing Func8 of smaller arity.
func (f Func9[T1, T2, T3, T4, T5, T6, T7, T8, T9, R]) PartialR(t9 T9) Func8[T1, T2, T3, T4, T5, T6, T7, T8, R] {
	return func(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8) R {
		return f(t1, t2, t3, t4, t5, t6, t7, t8, t9)
	}
}

// Func10 is a function with 10 arguments and 1 return value (10-ary function),
// which supports partial application.
type Func10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, R any] func(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) R

// MFunc10 casts function f with 10 arguments and 1 return value to Func10
func MFunc10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, R any](f Func10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, R]) Func10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, R] {
	return f
}

// Partial binds the first argument (from left to right) of Func10 f to value t1,
// producing Func9 of smaller arity.
func (f Func10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, R]) Partial(t1 T1) Func9[T2, T3, T4, T5, T6, T7, T8, T9, T10, R] {
	return func(t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10) R {
		return f(t1, t2, t3, t4, t5, t6, t7, t8, t9, t10)
	}
}

// PartialR binds the first argument (from right to left) of Func10 f to value t10,
// producing Func9 of smaller arity.
func (f Func10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, R]) PartialR(t10 T10) Func9[T1, T2, T3, T4, T5, T6, T7, T8, T9, R] {
	return func(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9) R {
		return f(t1, t2, t3, t4, t5, t6, t7, t8, t9, t10)
	}
}
