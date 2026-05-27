package loslice

import "slices"

// Flatten returns an array a single level deep.
// Play: https://go.dev/play/p/rbp9ORaMpjw
func Flatten[Slice ~[]T, T any](table ...Slice) Slice {
	if len(table) == 0 || Every(table, IsNil) {
		return nil
	}

	totalLen := MapSum(table, Len)
	result := make(Slice, 0, totalLen)
	for i := range table {
		result = append(result, table[i]...)
	}

	return result
}

func Transpose[Slice ~[]T, T any](table ...Slice) []Slice {
	n := len(table)
	switch n {
	case 0:
		return nil
	case 1:
		xs := table[0]

		if xs == nil {
			return nil
		}

		result := make([]Slice, len(xs))
		for i := range xs {
			result[i] = Slice{xs[i]}
		}

		return result
	}

	sizes := Uniq(Map(table, Len))
	if len(sizes) == 1 {
		// all table have the same size
		sz := sizes[0]

		if sz == 0 && Every(table, IsNil) {
			return nil
		}

		result := make([]Slice, sz)
		for i := range result {
			result[i] = make(Slice, n)
			for j := range result[i] {
				result[i][j] = table[j][i]
			}
		}

		return result
	}

	slices.Sort(sizes)
	j := 0
	result := make([]Slice, sizes[len(sizes)-1])
	for i := range result {
		if i == sizes[j] {
			j++

			// drop short rows
			last, size := 0, len(table)
			for last < size && len(table[last]) > i {
				last++
			}

			for k := last + 1; k < size; k++ {
				if len(table[k]) > i {
					table[last] = table[k]
					last++
				}
			}

			table = table[:last]
		}

		result[i] = make(Slice, len(table))
		for k := range table {
			result[i][k] = table[k][i]
		}
	}

	return result
}

// Interleave round-robin alternating input slices and sequentially appending value at index into result
func Interleave[Slice ~[]T, T any](table ...Slice) Slice {
	return Flatten(Transpose(table...)...)
}
