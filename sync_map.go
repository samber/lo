package lo

import "sync"

// SyncMap is a thread-safe map, which is a wrapper around sync.Map, but with a more convenient generic API.
type SyncMap[K comparable, V any] struct {
	data sync.Map
}

// Load returns the value stored in the map for a key, or default value if no value is present.
func (m *SyncMap[K, V]) Load(key K) (V, bool) {
	r, ok := m.data.Load(key)
	if ok {
		return r.(V), true
	}
	var rv V
	return rv, false
}

// Store sets the value for a key.
func (m *SyncMap[K, V]) Store(key K, value V) {
	m.data.Store(key, value)
}

// LoadOrStore returns the existing value for the key if present.
func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	r, ok := m.data.LoadOrStore(key, value)
	return r.(V), ok
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
func (m *SyncMap[K, V]) LoadAndDelete(key K) (V, bool) {
	r, ok := m.data.LoadAndDelete(key)
	if v1, ok1 := r.(V); ok1 {
		return v1, ok
	}
	var rv V
	return rv, false
}

// Delete deletes the value for a key.
func (m *SyncMap[K, V]) Delete(key K) {
	m.data.Delete(key)
}

// Range calls f sequentially for each key and value present in the map.
func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.data.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

// Snapshot returns a copy of the map based on Range method.
func (m *SyncMap[K, V]) Snapshot() map[K]V {
	snapshot := make(map[K]V)
	m.Range(func(key K, value V) bool {
		snapshot[key] = value
		return true
	})
	return snapshot
}

// NewSyncMap returns a new SyncMap.
func NewSyncMap[K comparable, V any]() SyncMap[K, V] {
	return SyncMap[K, V]{
		data: sync.Map{},
	}
}
