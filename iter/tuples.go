package iter

import (
	"iter"

	"github.com/samber/lo"
)

// Zip2 creates a sequence of grouped elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func Zip2[A, B any](a iter.Seq[A], b iter.Seq[B]) iter.Seq[lo.Tuple2[A, B]] {
	return func(yield func(lo.Tuple2[A, B]) bool) {
		var next lo.Tuple2[func() (A, bool), func() (B, bool)]
		var stop func()
		next.A, stop = iter.Pull(a)
		defer stop()
		next.B, stop = iter.Pull(b)
		defer stop()

		for {
			var item lo.Tuple2[A, B]
			var ok [2]bool
			item.A, ok[0] = next.A()
			item.B, ok[1] = next.B()
			if ok == [2]bool{} {
				return
			}
			yield(item)
		}
	}
}

// Zip3 creates a sequence of grouped elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func Zip3[A, B, C any](a iter.Seq[A], b iter.Seq[B], c iter.Seq[C]) iter.Seq[lo.Tuple3[A, B, C]] {
	return func(yield func(lo.Tuple3[A, B, C]) bool) {
		var next lo.Tuple3[func() (A, bool), func() (B, bool), func() (C, bool)]
		var stop func()
		next.A, stop = iter.Pull(a)
		defer stop()
		next.B, stop = iter.Pull(b)
		defer stop()
		next.C, stop = iter.Pull(c)
		defer stop()

		for {
			var item lo.Tuple3[A, B, C]
			var ok [3]bool
			item.A, ok[0] = next.A()
			item.B, ok[1] = next.B()
			item.C, ok[2] = next.C()
			if ok == [3]bool{} {
				return
			}
			yield(item)
		}
	}
}

// Zip4 creates a sequence of grouped elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func Zip4[A, B, C, D any](a iter.Seq[A], b iter.Seq[B], c iter.Seq[C], d iter.Seq[D]) iter.Seq[lo.Tuple4[A, B, C, D]] {
	return func(yield func(lo.Tuple4[A, B, C, D]) bool) {
		var next lo.Tuple4[func() (A, bool), func() (B, bool), func() (C, bool), func() (D, bool)]
		var stop func()
		next.A, stop = iter.Pull(a)
		defer stop()
		next.B, stop = iter.Pull(b)
		defer stop()
		next.C, stop = iter.Pull(c)
		defer stop()
		next.D, stop = iter.Pull(d)
		defer stop()

		for {
			var item lo.Tuple4[A, B, C, D]
			var ok [4]bool
			item.A, ok[0] = next.A()
			item.B, ok[1] = next.B()
			item.C, ok[2] = next.C()
			item.D, ok[3] = next.D()
			if ok == [4]bool{} {
				return
			}
			yield(item)
		}
	}
}

// Zip5 creates a sequence of grouped elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func Zip5[A, B, C, D, E any](a iter.Seq[A], b iter.Seq[B], c iter.Seq[C], d iter.Seq[D], e iter.Seq[E]) iter.Seq[lo.Tuple5[A, B, C, D, E]] {
	return func(yield func(lo.Tuple5[A, B, C, D, E]) bool) {
		var next lo.Tuple5[func() (A, bool), func() (B, bool), func() (C, bool), func() (D, bool), func() (E, bool)]
		var stop func()
		next.A, stop = iter.Pull(a)
		defer stop()
		next.B, stop = iter.Pull(b)
		defer stop()
		next.C, stop = iter.Pull(c)
		defer stop()
		next.D, stop = iter.Pull(d)
		defer stop()
		next.E, stop = iter.Pull(e)
		defer stop()

		for {
			var item lo.Tuple5[A, B, C, D, E]
			var ok [5]bool
			item.A, ok[0] = next.A()
			item.B, ok[1] = next.B()
			item.C, ok[2] = next.C()
			item.D, ok[3] = next.D()
			item.E, ok[4] = next.E()
			if ok == [5]bool{} {
				return
			}
			yield(item)
		}
	}
}

// Zip6 creates a sequence of grouped elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func Zip6[A, B, C, D, E, F any](a iter.Seq[A], b iter.Seq[B], c iter.Seq[C], d iter.Seq[D], e iter.Seq[E], f iter.Seq[F]) iter.Seq[lo.Tuple6[A, B, C, D, E, F]] {
	return func(yield func(lo.Tuple6[A, B, C, D, E, F]) bool) {
		var next lo.Tuple6[func() (A, bool), func() (B, bool), func() (C, bool), func() (D, bool), func() (E, bool), func() (F, bool)]
		var stop func()
		next.A, stop = iter.Pull(a)
		defer stop()
		next.B, stop = iter.Pull(b)
		defer stop()
		next.C, stop = iter.Pull(c)
		defer stop()
		next.D, stop = iter.Pull(d)
		defer stop()
		next.E, stop = iter.Pull(e)
		defer stop()
		next.F, stop = iter.Pull(f)
		defer stop()

		for {
			var item lo.Tuple6[A, B, C, D, E, F]
			var ok [6]bool
			item.A, ok[0] = next.A()
			item.B, ok[1] = next.B()
			item.C, ok[2] = next.C()
			item.D, ok[3] = next.D()
			item.E, ok[4] = next.E()
			item.F, ok[5] = next.F()
			if ok == [6]bool{} {
				return
			}
			yield(item)
		}
	}
}

// Zip7 creates a sequence of grouped elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func Zip7[A, B, C, D, E, F, G any](a iter.Seq[A], b iter.Seq[B], c iter.Seq[C], d iter.Seq[D], e iter.Seq[E], f iter.Seq[F], g iter.Seq[G]) iter.Seq[lo.Tuple7[A, B, C, D, E, F, G]] {
	return func(yield func(lo.Tuple7[A, B, C, D, E, F, G]) bool) {
		var next lo.Tuple7[func() (A, bool), func() (B, bool), func() (C, bool), func() (D, bool), func() (E, bool), func() (F, bool), func() (G, bool)]
		var stop func()
		next.A, stop = iter.Pull(a)
		defer stop()
		next.B, stop = iter.Pull(b)
		defer stop()
		next.C, stop = iter.Pull(c)
		defer stop()
		next.D, stop = iter.Pull(d)
		defer stop()
		next.E, stop = iter.Pull(e)
		defer stop()
		next.F, stop = iter.Pull(f)
		defer stop()
		next.G, stop = iter.Pull(g)
		defer stop()

		for {
			var item lo.Tuple7[A, B, C, D, E, F, G]
			var ok [7]bool
			item.A, ok[0] = next.A()
			item.B, ok[1] = next.B()
			item.C, ok[2] = next.C()
			item.D, ok[3] = next.D()
			item.E, ok[4] = next.E()
			item.F, ok[5] = next.F()
			item.G, ok[6] = next.G()
			if ok == [7]bool{} {
				return
			}
			yield(item)
		}
	}
}

// Zip8 creates a sequence of grouped elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func Zip8[A, B, C, D, E, F, G, H any](a iter.Seq[A], b iter.Seq[B], c iter.Seq[C], d iter.Seq[D], e iter.Seq[E], f iter.Seq[F], g iter.Seq[G], h iter.Seq[H]) iter.Seq[lo.Tuple8[A, B, C, D, E, F, G, H]] {
	return func(yield func(lo.Tuple8[A, B, C, D, E, F, G, H]) bool) {
		var next lo.Tuple8[func() (A, bool), func() (B, bool), func() (C, bool), func() (D, bool), func() (E, bool), func() (F, bool), func() (G, bool), func() (H, bool)]
		var stop func()
		next.A, stop = iter.Pull(a)
		defer stop()
		next.B, stop = iter.Pull(b)
		defer stop()
		next.C, stop = iter.Pull(c)
		defer stop()
		next.D, stop = iter.Pull(d)
		defer stop()
		next.E, stop = iter.Pull(e)
		defer stop()
		next.F, stop = iter.Pull(f)
		defer stop()
		next.G, stop = iter.Pull(g)
		defer stop()
		next.H, stop = iter.Pull(h)
		defer stop()

		for {
			var item lo.Tuple8[A, B, C, D, E, F, G, H]
			var ok [8]bool
			item.A, ok[0] = next.A()
			item.B, ok[1] = next.B()
			item.C, ok[2] = next.C()
			item.D, ok[3] = next.D()
			item.E, ok[4] = next.E()
			item.F, ok[5] = next.F()
			item.G, ok[6] = next.G()
			item.H, ok[7] = next.H()
			if ok == [8]bool{} {
				return
			}
			yield(item)
		}
	}
}

// Zip9 creates a sequence of grouped elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func Zip9[A, B, C, D, E, F, G, H, I any](a iter.Seq[A], b iter.Seq[B], c iter.Seq[C], d iter.Seq[D], e iter.Seq[E], f iter.Seq[F], g iter.Seq[G], h iter.Seq[H], i iter.Seq[I]) iter.Seq[lo.Tuple9[A, B, C, D, E, F, G, H, I]] {
	return func(yield func(lo.Tuple9[A, B, C, D, E, F, G, H, I]) bool) {
		var next lo.Tuple9[func() (A, bool), func() (B, bool), func() (C, bool), func() (D, bool), func() (E, bool), func() (F, bool), func() (G, bool), func() (H, bool), func() (I, bool)]
		var stop func()
		next.A, stop = iter.Pull(a)
		defer stop()
		next.B, stop = iter.Pull(b)
		defer stop()
		next.C, stop = iter.Pull(c)
		defer stop()
		next.D, stop = iter.Pull(d)
		defer stop()
		next.E, stop = iter.Pull(e)
		defer stop()
		next.F, stop = iter.Pull(f)
		defer stop()
		next.G, stop = iter.Pull(g)
		defer stop()
		next.H, stop = iter.Pull(h)
		defer stop()
		next.I, stop = iter.Pull(i)
		defer stop()

		for {
			var item lo.Tuple9[A, B, C, D, E, F, G, H, I]
			var ok [9]bool
			item.A, ok[0] = next.A()
			item.B, ok[1] = next.B()
			item.C, ok[2] = next.C()
			item.D, ok[3] = next.D()
			item.E, ok[4] = next.E()
			item.F, ok[5] = next.F()
			item.G, ok[6] = next.G()
			item.H, ok[7] = next.H()
			item.I, ok[8] = next.I()
			if ok == [9]bool{} {
				return
			}
			yield(item)
		}
	}
}

// ZipBy2 creates a sequence of transformed elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func ZipBy2[A, B, Out any](a iter.Seq[A], b iter.Seq[B], iteratee func(a A, b B) Out) iter.Seq[Out] {
	return Map(Zip2(a, b), func(item lo.Tuple2[A, B]) Out {
		return iteratee(item.A, item.B)
	})
}

// ZipBy3 creates a sequence of transformed elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func ZipBy3[A, B, C, Out any](a iter.Seq[A], b iter.Seq[B], c iter.Seq[C], iteratee func(a A, b B, c C) Out) iter.Seq[Out] {
	return Map(Zip3(a, b, c), func(item lo.Tuple3[A, B, C]) Out {
		return iteratee(item.A, item.B, item.C)
	})
}

// ZipBy4 creates a sequence of transformed elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func ZipBy4[A, B, C, D, Out any](a iter.Seq[A], b iter.Seq[B], c iter.Seq[C], d iter.Seq[D], iteratee func(a A, b B, c C, d D) Out) iter.Seq[Out] {
	return Map(Zip4(a, b, c, d), func(item lo.Tuple4[A, B, C, D]) Out {
		return iteratee(item.A, item.B, item.C, item.D)
	})
}

// ZipBy5 creates a sequence of transformed elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func ZipBy5[A, B, C, D, E, Out any](a iter.Seq[A], b iter.Seq[B], c iter.Seq[C], d iter.Seq[D], e iter.Seq[E], iteratee func(a A, b B, c C, d D, e E) Out) iter.Seq[Out] {
	return Map(Zip5(a, b, c, d, e), func(item lo.Tuple5[A, B, C, D, E]) Out {
		return iteratee(item.A, item.B, item.C, item.D, item.E)
	})
}

// ZipBy6 creates a sequence of transformed elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func ZipBy6[A, B, C, D, E, F, Out any](a iter.Seq[A], b iter.Seq[B], c iter.Seq[C], d iter.Seq[D], e iter.Seq[E], f iter.Seq[F], iteratee func(a A, b B, c C, d D, e E, f F) Out) iter.Seq[Out] {
	return Map(Zip6(a, b, c, d, e, f), func(item lo.Tuple6[A, B, C, D, E, F]) Out {
		return iteratee(item.A, item.B, item.C, item.D, item.E, item.F)
	})
}

// ZipBy7 creates a sequence of transformed elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func ZipBy7[A, B, C, D, E, F, G, Out any](a iter.Seq[A], b iter.Seq[B], c iter.Seq[C], d iter.Seq[D], e iter.Seq[E], f iter.Seq[F], g iter.Seq[G], iteratee func(a A, b B, c C, d D, e E, f F, g G) Out) iter.Seq[Out] {
	return Map(Zip7(a, b, c, d, e, f, g), func(item lo.Tuple7[A, B, C, D, E, F, G]) Out {
		return iteratee(item.A, item.B, item.C, item.D, item.E, item.F, item.G)
	})
}

// ZipBy8 creates a sequence of transformed elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func ZipBy8[A, B, C, D, E, F, G, H, Out any](a iter.Seq[A], b iter.Seq[B], c iter.Seq[C], d iter.Seq[D], e iter.Seq[E], f iter.Seq[F], g iter.Seq[G], h iter.Seq[H], iteratee func(a A, b B, c C, d D, e E, f F, g G, h H) Out) iter.Seq[Out] {
	return Map(Zip8(a, b, c, d, e, f, g, h), func(item lo.Tuple8[A, B, C, D, E, F, G, H]) Out {
		return iteratee(item.A, item.B, item.C, item.D, item.E, item.F, item.G, item.H)
	})
}

// ZipBy9 creates a sequence of transformed elements, the first of which contains the first elements
// of the given sequences, the second of which contains the second elements of the given sequences, and so on.
// When collections are different sizes, the Tuple attributes are filled with zero value.
func ZipBy9[A, B, C, D, E, F, G, H, I, Out any](a iter.Seq[A], b iter.Seq[B], c iter.Seq[C], d iter.Seq[D], e iter.Seq[E], f iter.Seq[F], g iter.Seq[G], h iter.Seq[H], i iter.Seq[I], iteratee func(a A, b B, c C, d D, e E, f F, g G, h H, i I) Out) iter.Seq[Out] {
	return Map(Zip9(a, b, c, d, e, f, g, h, i), func(item lo.Tuple9[A, B, C, D, E, F, G, H, I]) Out {
		return iteratee(item.A, item.B, item.C, item.D, item.E, item.F, item.G, item.H, item.I)
	})
}

// CrossJoin2 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
func CrossJoin2[A, B any](listA iter.Seq[A], listB iter.Seq[B]) iter.Seq[lo.Tuple2[A, B]] {
	return CrossJoinBy2(listA, listB, lo.T2[A, B])
}

// CrossJoin3 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
func CrossJoin3[A, B, C any](listA iter.Seq[A], listB iter.Seq[B], listC iter.Seq[C]) iter.Seq[lo.Tuple3[A, B, C]] {
	return CrossJoinBy3(listA, listB, listC, lo.T3[A, B, C])
}

// CrossJoin4 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
func CrossJoin4[A, B, C, D any](listA iter.Seq[A], listB iter.Seq[B], listC iter.Seq[C], listD iter.Seq[D]) iter.Seq[lo.Tuple4[A, B, C, D]] {
	return CrossJoinBy4(listA, listB, listC, listD, lo.T4[A, B, C, D])
}

// CrossJoin5 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
func CrossJoin5[A, B, C, D, E any](listA iter.Seq[A], listB iter.Seq[B], listC iter.Seq[C], listD iter.Seq[D], listE iter.Seq[E]) iter.Seq[lo.Tuple5[A, B, C, D, E]] {
	return CrossJoinBy5(listA, listB, listC, listD, listE, lo.T5[A, B, C, D, E])
}

// CrossJoin6 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
func CrossJoin6[A, B, C, D, E, F any](listA iter.Seq[A], listB iter.Seq[B], listC iter.Seq[C], listD iter.Seq[D], listE iter.Seq[E], listF iter.Seq[F]) iter.Seq[lo.Tuple6[A, B, C, D, E, F]] {
	return CrossJoinBy6(listA, listB, listC, listD, listE, listF, lo.T6[A, B, C, D, E, F])
}

// CrossJoin7 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
func CrossJoin7[A, B, C, D, E, F, G any](listA iter.Seq[A], listB iter.Seq[B], listC iter.Seq[C], listD iter.Seq[D], listE iter.Seq[E], listF iter.Seq[F], listG iter.Seq[G]) iter.Seq[lo.Tuple7[A, B, C, D, E, F, G]] {
	return CrossJoinBy7(listA, listB, listC, listD, listE, listF, listG, lo.T7[A, B, C, D, E, F, G])
}

// CrossJoin8 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
func CrossJoin8[A, B, C, D, E, F, G, H any](listA iter.Seq[A], listB iter.Seq[B], listC iter.Seq[C], listD iter.Seq[D], listE iter.Seq[E], listF iter.Seq[F], listG iter.Seq[G], listH iter.Seq[H]) iter.Seq[lo.Tuple8[A, B, C, D, E, F, G, H]] {
	return CrossJoinBy8(listA, listB, listC, listD, listE, listF, listG, listH, lo.T8[A, B, C, D, E, F, G, H])
}

// CrossJoin9 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments.
// Returns an empty list if a list is empty.
func CrossJoin9[A, B, C, D, E, F, G, H, I any](listA iter.Seq[A], listB iter.Seq[B], listC iter.Seq[C], listD iter.Seq[D], listE iter.Seq[E], listF iter.Seq[F], listG iter.Seq[G], listH iter.Seq[H], listI iter.Seq[I]) iter.Seq[lo.Tuple9[A, B, C, D, E, F, G, H, I]] {
	return CrossJoinBy9(listA, listB, listC, listD, listE, listF, listG, listH, listI, lo.T9[A, B, C, D, E, F, G, H, I])
}

// CrossJoinBy2 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The project function
// is used to create the output values.
// Returns an empty list if a list is empty.
func CrossJoinBy2[A, B, Out any](listA iter.Seq[A], listB iter.Seq[B], project func(a A, b B) Out) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for a := range listA {
			for b := range listB {
				if !yield(project(a, b)) {
					return
				}
			}
		}
	}
}

// CrossJoinBy3 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The project function
// is used to create the output values.
// Returns an empty list if a list is empty.
func CrossJoinBy3[A, B, C, Out any](listA iter.Seq[A], listB iter.Seq[B], listC iter.Seq[C], project func(a A, b B, c C) Out) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for a := range listA {
			for b := range listB {
				for c := range listC {
					if !yield(project(a, b, c)) {
						return
					}
				}
			}
		}
	}
}

// CrossJoinBy4 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The project function
// is used to create the output values.
// Returns an empty list if a list is empty.
func CrossJoinBy4[A, B, C, D, Out any](listA iter.Seq[A], listB iter.Seq[B], listC iter.Seq[C], listD iter.Seq[D], project func(a A, b B, c C, d D) Out) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for a := range listA {
			for b := range listB {
				for c := range listC {
					for d := range listD {
						if !yield(project(a, b, c, d)) {
							return
						}
					}
				}
			}
		}
	}
}

// CrossJoinBy5 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The project function
// is used to create the output values.
// Returns an empty list if a list is empty.
func CrossJoinBy5[A, B, C, D, E, Out any](listA iter.Seq[A], listB iter.Seq[B], listC iter.Seq[C], listD iter.Seq[D], listE iter.Seq[E], project func(a A, b B, c C, d D, e E) Out) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for a := range listA {
			for b := range listB {
				for c := range listC {
					for d := range listD {
						for e := range listE {
							if !yield(project(a, b, c, d, e)) {
								return
							}
						}
					}
				}
			}
		}
	}
}

// CrossJoinBy6 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The project function
// is used to create the output values.
// Returns an empty list if a list is empty.
func CrossJoinBy6[A, B, C, D, E, F, Out any](listA iter.Seq[A], listB iter.Seq[B], listC iter.Seq[C], listD iter.Seq[D], listE iter.Seq[E], listF iter.Seq[F], project func(a A, b B, c C, d D, e E, f F) Out) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for a := range listA {
			for b := range listB {
				for c := range listC {
					for d := range listD {
						for e := range listE {
							for f := range listF {
								if !yield(project(a, b, c, d, e, f)) {
									return
								}
							}
						}
					}
				}
			}
		}
	}
}

// CrossJoinBy7 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The project function
// is used to create the output values.
// Returns an empty list if a list is empty.
func CrossJoinBy7[A, B, C, D, E, F, G, Out any](listA iter.Seq[A], listB iter.Seq[B], listC iter.Seq[C], listD iter.Seq[D], listE iter.Seq[E], listF iter.Seq[F], listG iter.Seq[G], project func(a A, b B, c C, d D, e E, f F, g G) Out) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for a := range listA {
			for b := range listB {
				for c := range listC {
					for d := range listD {
						for e := range listE {
							for f := range listF {
								for g := range listG {
									if !yield(project(a, b, c, d, e, f, g)) {
										return
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

// CrossJoinBy8 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The project function
// is used to create the output values.
// Returns an empty list if a list is empty.
func CrossJoinBy8[A, B, C, D, E, F, G, H, Out any](listA iter.Seq[A], listB iter.Seq[B], listC iter.Seq[C], listD iter.Seq[D], listE iter.Seq[E], listF iter.Seq[F], listG iter.Seq[G], listH iter.Seq[H], project func(a A, b B, c C, d D, e E, f F, g G, h H) Out) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for a := range listA {
			for b := range listB {
				for c := range listC {
					for d := range listD {
						for e := range listE {
							for f := range listF {
								for g := range listG {
									for h := range listH {
										if !yield(project(a, b, c, d, e, f, g, h)) {
											return
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

// CrossJoinBy9 combines every item from one list with every item from others.
// It is the cartesian product of lists received as arguments. The project function
// is used to create the output values.
// Returns an empty list if a list is empty.
func CrossJoinBy9[A, B, C, D, E, F, G, H, I, Out any](listA iter.Seq[A], listB iter.Seq[B], listC iter.Seq[C], listD iter.Seq[D], listE iter.Seq[E], listF iter.Seq[F], listG iter.Seq[G], listH iter.Seq[H], listI iter.Seq[I], project func(a A, b B, c C, d D, e E, f F, g G, h H, i I) Out) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for a := range listA {
			for b := range listB {
				for c := range listC {
					for d := range listD {
						for e := range listE {
							for f := range listF {
								for g := range listG {
									for h := range listH {
										for i := range listI {
											if !yield(project(a, b, c, d, e, f, g, h, i)) {
												return
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
