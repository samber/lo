package loset

import (
	"github.com/samber/lo/lomap"
	"github.com/samber/lo/loslice"
	"github.com/samber/lo/mutmap"
	"maps"
)

func KeepCommon[S ~Set[T], T comparable](self, other S, more ...S) {
	if len(other) == 0 || loslice.Contains(more, lomap.IsEmpty) {
		mutmap.Clear(self) // Clear the set if the other set is empty or if any of the additional sets are empty
		return
	}

	for k := range self {
		if _, ok := other[k]; !ok {
			delete(self, k)
		}
	}

	if len(more) == 0 {
		return // If no additional sets are provided, we are done
	}

	for k := range self {
		for _, s := range more {
			if _, ok := s[k]; !ok {
				delete(self, k)
				break
			}
		}
	}
}

func Subtract[S ~Set[T], T comparable](self, other S, more ...S) {
	for k := range self {
		if _, ok := other[k]; ok {
			delete(self, k)
		}
	}

	if len(more) == 0 || loslice.Every(more, lomap.IsEmpty) {
		return // If no additional sets are provided, or they have no elements, we are done
	}

	for k := range self {
		for _, s := range more {
			if _, ok := s[k]; ok {
				delete(self, k)
				break
			}
		}
	}
}

func Add[S ~Set[T], T comparable](self, other S, more ...S) {
	maps.Copy(self, other)
	for _, s := range more {
		maps.Copy(self, s)
	}
}
