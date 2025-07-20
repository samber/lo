package loset

import "maps"

func KeepCommon[T comparable, Set ~map[T]struct{}](self, other Set) {
	for k := range self {
		if _, ok := other[k]; !ok {
			delete(self, k)
		}
	}
}

func Subtract[T comparable, Set ~map[T]struct{}](self, other Set) {
	for k := range self {
		if _, ok := other[k]; ok {
			delete(self, k)
		}
	}
}

func Add[T comparable, Set ~map[T]struct{}](self Set, other Set) {
	maps.Copy(self, other)
}
