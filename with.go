package lo

//With modifies a value for the duration of a callback
func With[T any](ptr *T, val T, callback func()) {
	oldVal := *ptr
	*ptr = val
	callback()
	*ptr = oldVal
}
