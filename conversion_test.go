package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToQuery(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := ToQuery(map[string]string{"foo": "bar", "bar": "baz", "k": "v"})
	result2 := ToQuery(map[string]string{"foo": "bar", "bar": "baz"})
	result3 := ToQuery(map[string]string{"key": "value"})

	is.Equal(result1, "foo=bar&bar=baz&k=v")
	is.Equal(result2, "foo=bar&bar=baz")
	is.Equal(result3, "key=value")
}

func TestFromQuery(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FromQuery("foo=bar&bar=baz&k=v")
	result2 := FromQuery("foo=bar&bar=baz")
	result3 := FromQuery("key=value")

	is.Equal(result1, map[string]string{"foo": "bar", "bar": "baz", "k": "v"})
	is.Equal(result2, map[string]string{"foo": "bar", "bar": "baz"})
	is.Equal(result3, map[string]string{"key": "value"})
}
