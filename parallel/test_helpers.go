package parallel

import "time"

func withinDuration(expected, actual time.Time, delta time.Duration) bool {
	dt := expected.Sub(actual)

	return dt > -delta && dt < delta
}

type orderedMapValue struct {
	value    int
	calledAt time.Time
}

type orderedMap struct {
	m map[string]orderedMapValue
	s []struct {
		key   string
		value orderedMapValue
	}
}

func (o *orderedMap) Add(key string, value orderedMapValue) {
	if o.m == nil {
		o.m = make(map[string]orderedMapValue)
	}
	o.m[key] = value
	o.s = append(o.s, struct {
		key   string
		value orderedMapValue
	}{key, value})
}
