package loset

import (
	"github.com/samber/lo/loslice"
	"github.com/samber/lo/lotup"
	"maps"
)

func Intersection[S ~Set[T], T comparable](a, b S) (intersection S) {
	if len(b) < len(a) {
		// swap a and b to ensure a is the smaller set
		a, b = b, a
	}

	intersection = make(S, len(a))
	for k := range a {
		if _, ok := b[k]; ok {
			intersection[k] = struct{}{}
		}
	}

	return
}

func GroupIntersection[S ~Set[T], T comparable](sets ...S) (intersection S) {
	if len(sets) == 0 {
		return nil
	}

	// look for the smallest set to reduce iterations
	i := lotup.First(loslice.ArgMin(sets, func(s S) int { return len(s) }))
	intersection = maps.Clone(sets[i])

	copy(sets[i:], sets[i+1:])
	sets = sets[:len(sets)-1]

	for it := range intersection {
		for _, set := range sets {
			if _, ok := set[it]; !ok {
				delete(intersection, it)
				break
			}
		}
	}

	return intersection
}

func Union[S ~Set[T], T comparable](a, b S) (union S) {
	if len(b) > len(a) {
		// swap a and b to ensure a is the larger set
		a, b = b, a
	}

	union = maps.Clone(a)
	maps.Copy(union, b)

	return union
}

func GroupUnion[S ~Set[T], T comparable](sets ...S) (union S) {
	if len(sets) == 0 {
		return nil
	}

	union = make(S)
	for _, set := range sets {
		maps.Copy(union, set)
	}

	return union
}

func Difference[S ~Set[T], T comparable](a, b S) (difference S) {
	difference = make(S, len(a))
	for k := range a {
		if _, ok := b[k]; !ok {
			difference[k] = struct{}{}
		}
	}

	return
}

func SymmetricDifference[S ~Set[T], T comparable](a, b S) (symmetricDifference S) {
	if len(b) > len(a) {
		// swap a and b to ensure a is the larger set
		a, b = b, a
	}

	symmetricDifference = maps.Clone(a)

	for k := range b {
		if _, ok := a[k]; ok {
			delete(symmetricDifference, k)
		} else {
			symmetricDifference[k] = struct{}{}
		}
	}

	return symmetricDifference
}

func GroupSymmetricDifference[S ~Set[T], T comparable](sets ...S) (symmetricDifference S) {
	if len(sets) == 0 {
		return nil
	}

	rejected := make(S)
	symmetricDifference = make(S)
	for _, s := range sets {
		for k := range s {
			if _, ok := rejected[k]; ok {
				continue
			} else if _, ok := symmetricDifference[k]; ok {
				rejected[k] = struct{}{}
				delete(symmetricDifference, k)
			} else {
				symmetricDifference[k] = struct{}{}
			}
		}
	}

	return symmetricDifference
}
