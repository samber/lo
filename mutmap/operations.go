package mutmap

import (
	"github.com/samber/lo/lomap"
	"github.com/samber/lo/loslice"
)

// Update replaces existing keys in the map with values from another map.
// First map in the list has the highest priority.
// No new keys are added, only existing keys are updated.
func Update[Map ~map[K]V, K comparable, V any](self, other Map, more ...Map) {
	for k := range self {
		if v, exists := other[k]; exists {
			self[k] = v
			continue
		}
	}

	if len(more) == 0 || loslice.Every(more, lomap.IsEmpty) {
		return // If no additional maps are provided, or they have no elements, we are done
	}

	for k := range self {
		for _, m := range more {
			if v, exists := m[k]; exists {
				self[k] = v
				break // Stop at the first match
			}
		}
	}
}

// Add adds new keys and values from another map to the current map.
// First map in the list has the highest priority.
// No existing keys are modified, only new keys are added.
func Add[Map ~map[K]V, K comparable, V any](self, other Map, more ...Map) {
	for k, v := range other {
		if _, exists := self[k]; !exists {
			self[k] = v
		}
	}

	for _, m := range more {
		for k, v := range m {
			if _, exists := self[k]; !exists {
				self[k] = v
			}
		}
	}
}

// Remove deletes keys from the current map that are present in another map.
func Remove[Map ~map[K]V, K comparable, V any](self, other Map, more ...Map) {
	for k := range other {
		delete(self, k)
	}

	if len(more) == 0 || loslice.Every(more, lomap.IsEmpty) {
		return // If no additional maps are provided, or they have no elements, we are done
	}

	for k := range self {
		for _, m := range more {
			if _, exists := m[k]; exists {
				delete(self, k)
				break // Stop at the first match
			}
		}
	}
}

// KeepCommon retains only the keys that are present in both maps.
// Values left unchanged.
func KeepCommon[Map ~map[K]V, K comparable, V any](self, other Map, more ...Map) {
	if len(other) == 0 || loslice.Contains(more, lomap.IsEmpty) {
		Clear(self) // Clear the map if the other map is empty or if any of the additional maps are empty
		return
	}

	for k := range self {
		if _, exists := other[k]; !exists {
			delete(self, k)
		}
	}

	if len(more) == 0 {
		return // If no additional maps are provided, we are done
	}

	for k := range self {
		for _, m := range more {
			if _, exists := m[k]; !exists {
				delete(self, k)
				break // Stop at the first match
			}
		}
	}
}
