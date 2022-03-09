package lo

func Drop[T any](collection []T, n int) []T {
	if len(collection) <= n {
		return make([]T, 0)
	}

	result := make([]T, len(collection)-n)
	for i := n; i < len(collection); i++ {
		result[i-n] = collection[i]
	}

	return result
}

func DropWhile[T any](collection []T, predicate func(T) bool) []T {
	result := []T{}

	i := 0
	for ; i < len(collection); i++ {
		if !predicate(collection[i]) {
			break
		}
	}

	for ; i < len(collection); i++ {
		result = append(result, collection[i])
	}

	return result
}

func DropRight[T any](collection []T, n int) []T {
	if len(collection) <= n {
		return make([]T, 0)
	}

	result := make([]T, len(collection)-n)
	for i := len(collection) - 1 - n; i != 0; i-- {
		result[i] = collection[i]
	}

	return result
}

func DropRightWhile[T any](collection []T, predicate func(T) bool) []T {
	result := []T{}

	i := len(collection) - 1
	for ; i >= 0; i-- {
		if !predicate(collection[i]) {
			break
		}
	}

	for ; i >= 0; i-- {
		result = append([]T{collection[i]}, result...)
	}

	return result
}
