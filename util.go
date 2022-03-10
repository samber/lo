package lo

// Range creates an array of numbers (positive and/or negative) with given length.
func Range(elementNum int) []int {
	if elementNum >= 0 {
		result := make([]int, elementNum)
		for i := 0; i < elementNum; i++ {
			result[i] = i
		}
		return result
	}
	result := make([]int, -elementNum)
	for i := 0; i < -elementNum; i++ {
		result[i] = -i
	}
	return result
}

// RangeFrom creates an array of numbers from start with specified length.
func RangeFrom(start, elementNum int) []int {
	if elementNum >= 0 {
		result := make([]int, elementNum)
		for i := 0; i < elementNum; i++ {
			result[i] = i + start
		}
		return result
	}
	result := make([]int, -elementNum)
	for i := 0; i < -elementNum; i++ {
		result[i] = start - i
	}
	return result
}

// RangeOpen creates an array of numbers (positive and/or negative) progressing from start up to, but not including end.
// step is default 1 when set to zero or can not reach end.
func RangeOpen(start, end, step int) []int {
	var result []int
	if start == end {
		return result
	}
	if start < end {
		if step <= 0 {
			step = 1
		}
		for i := start; i < end; i += step {
			result = append(result, i)
		}
		return result
	}
	if step >= 0 {
		step = -1
	}
	for i := start; i > end; i += step {
		result = append(result, i)
	}
	return result
}
