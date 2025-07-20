package loset

import (
	"github.com/samber/lo/loslice"
	"github.com/samber/lo/lotup"
	"maps"
)

func Intersection[T comparable, Set ~map[T]struct{}](a, b Set) (intersection Set) {
	if len(b) < len(a) {
		// swap a and b to ensure a is the smaller set
		a, b = b, a
	}

	intersection = make(Set, len(a))
	for k := range a {
		if _, ok := b[k]; ok {
			intersection[k] = struct{}{}
		}
	}

	return
}

func GroupIntersection[T comparable, Set ~map[T]struct{}](sets ...Set) (intersection Set) {
	if len(sets) == 0 {
		return nil
	}

	// look for the smallest set to reduce iterations
	i := lotup.First(loslice.ArgMin(sets, func(s Set) int { return len(s) }))
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

func Union[T comparable, Set ~map[T]struct{}](a, b Set) (union Set) {
	if len(b) > len(a) {
		// swap a and b to ensure a is the larger set
		a, b = b, a
	}

	union = maps.Clone(a)
	maps.Copy(union, b)

	return union
}

func GroupUnion[T comparable, Set ~map[T]struct{}](sets ...Set) (union Set) {
	if len(sets) == 0 {
		return nil
	}

	union = make(Set)
	for _, set := range sets {
		maps.Copy(union, set)
	}

	return union
}

func Difference[T comparable, Set ~map[T]struct{}](a, b Set) (difference Set) {
	difference = make(Set, len(a))
	for k := range a {
		if _, ok := b[k]; !ok {
			difference[k] = struct{}{}
		}
	}

	return
}

func SymmetricDifference[T comparable, Set ~map[T]struct{}](a, b Set) (symmetricDifference Set) {
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

func GroupSymmetricDifference[T comparable, Set ~map[T]struct{}](sets ...Set) (symmetricDifference Set) {
	if len(sets) == 0 {
		return nil
	}

	rejected := make(Set)
	symmetricDifference = make(Set)
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
