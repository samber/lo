package lo

import "fmt"

const defaultAssertionFailureMessage = "assertion failed"

// Assert does nothing when the condition is true, otherwise it panics with an optional message.
func Assert(condition bool, message ...string) {
	if condition {
		return
	}

	panicMessage := defaultAssertionFailureMessage
	if len(message) > 0 {
		panicMessage = fmt.Sprintf("%s: %s", defaultAssertionFailureMessage, message[0])
	}
	panic(panicMessage)
}

// Assertf does nothing when the condition is true, otherwise it panics with a formatted message.
func Assertf(condition bool, format string, args ...any) {
	if condition {
		return
	}

	panicMessage := fmt.Sprintf("%s: %s", defaultAssertionFailureMessage, fmt.Sprintf(format, args...))
	panic(panicMessage)
}
