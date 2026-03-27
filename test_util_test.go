// test_util_test contains testing utility functions
package lo

import (
	"strconv"
	"strings"
)

func prettyFormatSlice(vals []int) string {
	valsStr := make([]string, len(vals))

	i := 0
	for _, v := range vals {
		valsStr[i] = strconv.Itoa(v)
		i++
	}

	return strings.Join(valsStr, ",")
}
