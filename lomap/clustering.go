package lomap

func ClusterValues[Map ~map[K]V, C, K comparable, V any](m Map, fcluster func(V) C) (clusters map[C]Map) {
	clusters = make(map[C]Map)

	for k, v := range m {
		c := fcluster(v)
		if _, exists := clusters[c]; !exists {
			clusters[c] = make(Map)
		}

		clusters[c][k] = v
	}

	return clusters
}

func ClusterKeys[Map ~map[K]V, C, K comparable, V any](m Map, fcluster func(K) C) (clusters map[C]Map) {
	clusters = make(map[C]Map)

	for k, v := range m {
		c := fcluster(k)
		if _, exists := clusters[c]; !exists {
			clusters[c] = make(Map)
		}

		clusters[c][k] = v
	}

	return clusters
}

func ClusterPairs[Map ~map[K]V, C, K comparable, V any](m Map, fcluster func(K, V) C) (clusters map[C]Map) {
	clusters = make(map[C]Map)

	for k, v := range m {
		c := fcluster(k, v)
		if _, exists := clusters[c]; !exists {
			clusters[c] = make(Map)
		}

		clusters[c][k] = v
	}

	return clusters
}

func FilterClusterValues[Map ~map[K]V, C, K comparable, V any](m Map, fcluster func(V) (C, bool)) (clusters map[C]Map) {
	clusters = make(map[C]Map)

	for k, v := range m {
		c, ok := fcluster(v)
		if !ok {
			continue
		}

		if _, exists := clusters[c]; !exists {
			clusters[c] = make(Map)
		}

		clusters[c][k] = v
	}

	return clusters
}

func FilterClusterKeys[Map ~map[K]V, C, K comparable, V any](m Map, fcluster func(K) (C, bool)) (clusters map[C]Map) {
	clusters = make(map[C]Map)

	for k, v := range m {
		c, ok := fcluster(k)
		if !ok {
			continue
		}

		if _, exists := clusters[c]; !exists {
			clusters[c] = make(Map)
		}

		clusters[c][k] = v
	}

	return clusters
}

func FilterClusterPairs[Map ~map[K]V, C, K comparable, V any](m Map, fcluster func(K, V) (C, bool)) (clusters map[C]Map) {
	clusters = make(map[C]Map)

	for k, v := range m {
		c, ok := fcluster(k, v)
		if !ok {
			continue
		}

		if _, exists := clusters[c]; !exists {
			clusters[c] = make(Map)
		}

		clusters[c][k] = v
	}

	return clusters
}

func FilterMapClusterValues[Map ~map[K]V, C, K comparable, V, R any](m Map, fmapcluster func(V) (C, R, bool)) (clusters map[C]map[K][]R) {
	clusters = make(map[C]map[K][]R)

	for k, v := range m {
		c, r, ok := fmapcluster(v)
		if !ok {
			continue
		}

		if _, exists := clusters[c]; !exists {
			clusters[c] = make(map[K][]R)
		}

		clusters[c][k] = append(clusters[c][k], r)
	}

	return clusters
}

func FilterMapClusterKeys[Map ~map[K]V, C, K comparable, V, R any](m Map, fmapcluster func(K) (C, R, bool)) (clusters map[C]map[K][]R) {
	clusters = make(map[C]map[K][]R)

	for k := range m {
		c, r, ok := fmapcluster(k)
		if !ok {
			continue
		}

		if _, exists := clusters[c]; !exists {
			clusters[c] = make(map[K][]R)
		}

		clusters[c][k] = append(clusters[c][k], r)
	}

	return clusters
}

func FilterMapClusterPairs[Map ~map[K]V, C, K comparable, V, R any](m Map, fmapcluster func(K, V) (C, R, bool)) (clusters map[C]map[K][]R) {
	clusters = make(map[C]map[K][]R)

	for k, v := range m {
		c, r, ok := fmapcluster(k, v)
		if !ok {
			continue
		}

		if _, exists := clusters[c]; !exists {
			clusters[c] = make(map[K][]R)
		}

		clusters[c][k] = append(clusters[c][k], r)
	}

	return clusters
}

func MapClusterValues[Map ~map[K]V, C, K comparable, V, R any](m Map, fmapcluster func(V) (C, R)) (clusters map[C]map[K][]R) {
	clusters = make(map[C]map[K][]R)

	for k, v := range m {
		c, r := fmapcluster(v)

		if _, exists := clusters[c]; !exists {
			clusters[c] = make(map[K][]R)
		}

		clusters[c][k] = append(clusters[c][k], r)
	}

	return clusters
}

func MapClusterKeys[Map ~map[K]V, C, K comparable, V, R any](m Map, fmapcluster func(K) (C, R)) (clusters map[C]map[K][]R) {
	clusters = make(map[C]map[K][]R)

	for k := range m {
		c, r := fmapcluster(k)

		if _, exists := clusters[c]; !exists {
			clusters[c] = make(map[K][]R)
		}

		clusters[c][k] = append(clusters[c][k], r)
	}

	return clusters
}

func MapClusterPairs[Map ~map[K]V, C, K comparable, V, R any](m Map, fmapcluster func(K, V) (C, R)) (clusters map[C]map[K][]R) {
	clusters = make(map[C]map[K][]R)

	for k, v := range m {
		c, r := fmapcluster(k, v)

		if _, exists := clusters[c]; !exists {
			clusters[c] = make(map[K][]R)
		}

		clusters[c][k] = append(clusters[c][k], r)
	}

	return clusters
}
