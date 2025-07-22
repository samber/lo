package loset

import "maps"

func KeepCommon[S ~Set[T], T comparable](self, other S) {
	for k := range self {
		if _, ok := other[k]; !ok {
			delete(self, k)
		}
	}
}

func Subtract[S ~Set[T], T comparable](self, other S) {
	for k := range self {
		if _, ok := other[k]; ok {
			delete(self, k)
		}
	}
}

func Add[S ~Set[T], T comparable](self, other S) {
	maps.Copy(self, other)
}
