package lo

import (
	"fmt"
	"strings"
)

// form query convert string to query string
func ToQuery[K comparable, V any](in map[K]V) string {
	res := ""

	firstIteration := true
	for k, v := range in {
		if firstIteration {
			res += fmt.Sprintf("%v=%v", k, v)
			firstIteration = false
		} else {
			res += fmt.Sprintf("&%v=%v", k, v)
		}
	}

	return res
}

// form query convert query string to string
func FromQuery(s string) map[string]string {
	result := make(map[string]string)

	pairs := strings.Split(s, "&")

	for _, pair := range pairs {
		parts := strings.Split(pair, "=")
		if len(parts) == 2 {
			result[parts[0]] = parts[1]
		}
	}

	return result
}
