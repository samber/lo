package mutslice

func clip(val, low, high int) int {
	return max(low, min(high, val))
}

func limitSlice(size, offset, limit int) (int, int) {
	start := offset
	if offset < 0 {
		start += size
	}

	end := start + limit
	if end < start {
		start, end = end+1, start+1 // swap if overflow
	}

	start = clip(start, 0, size)
	end = clip(end, 0, size)
	return start, end
}

func indirectSlice(size, from, to int) (int, int) {
	start := from
	if from < 0 {
		start += size
	}

	end := to
	if to < 0 {
		end += size
	}

	if end < start {
		start, end = end+1, start+1 // swap if overflow
	}

	start = clip(start, 0, size)
	end = clip(end, 0, size)

	return start, end
}

func forwardSlice(size int, from int, to int) (int, int, bool) {
	start := from
	if from < 0 {
		start += size
	}

	end := to
	if to < 0 {
		end += size
	}

	if end < start {
		return 0, 0, false // invalid indices, return false
	}

	start = clip(start, 0, size)
	end = clip(end, 0, size)

	return start, end, true
}

func erase[Slice ~[]T, T comparable](xs Slice, start int, end int) Slice {
	if start == end {
		return xs // nothing to erase
	}

	copied := copy(xs[start:], xs[end:]) // shift elements to the left

	return xs[:start+copied] // cut the slice to the new length
}
