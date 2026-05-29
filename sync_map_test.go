package lo

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncMap(t *testing.T) {

	mp := NewSyncMap[string, []string]()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			k := strconv.FormatInt(int64(idx), 10)
			mp.Store("key"+k, []string{"value" + k})
		}(i)
	}
	wg.Wait()

	size := 0
	mp.Range(func(key string, value []string) bool {
		size++
		assert.Equal(t, "value"+key[3:], value[0])
		return true
	})
	assert.Equal(t, 10, size)

	snapshot := mp.Snapshot()

	assert.Equal(t, 10, len(snapshot))
}
