//go:build go1.21

package constraints

import (
	"cmp"
)

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
type Ordered = cmp.Ordered

// Less orders floating-point NaNs before other values, as in the standard library.
func Less[T Ordered](x, y T) bool {
	return cmp.Less(x, y)
}
