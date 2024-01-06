package lo

import (
	"fmt"
	"strings"
)

func ToQuery[K comparable, V any](in map[K]V) string {
	res := ""

	// Use a flag to check if it's the first iteration
	firstIteration := true
	for k, v := range in {
		// Check if it's the first iteration
		if firstIteration {
			res += fmt.Sprintf("%v=%v", k, v)
			firstIteration = false
		} else {
			res += fmt.Sprintf("&%v=%v", k, v)
		}
	}

	return res
}

func FromQuery(s string) map[string]string {
	result := make(map[string]string)

	// Split the input string by "&" to get individual key-value pairs
	pairs := strings.Split(s, "&")

	for _, pair := range pairs {
		// Split each pair by "=" to get the key and value
		parts := strings.Split(pair, "=")
		if len(parts) == 2 {
			result[parts[0]] = parts[1]
		}
	}

	return result
}
