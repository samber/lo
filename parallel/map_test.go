package parallel

import (
	"sync"
	"testing"
	"time"
)

func TestForEachKeyValue(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		input := map[string]int{}
		ForEachKeyValue(input, func(_ string, _ int) {
			t.Errorf("unexpected fn call for empty map")
		})
	})
	t.Run("call fn for each key value concurrently", func(t *testing.T) {
		input := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}
		result := map[string]struct {
			value    int
			calledAt time.Time
		}{}
		mu := sync.Mutex{}
		startTime := time.Now()
		allowedCallDelta := time.Millisecond * 2
		ForEachKeyValue(input, func(key string, value int) {
			mu.Lock()
			result[key] = struct {
				value    int
				calledAt time.Time
			}{value, time.Now()}
			mu.Unlock()
			time.Sleep(time.Millisecond * 50)
		})

		if resultLen, inputLen := len(result), len(input); resultLen != inputLen {
			t.Errorf("result len of %d is not equal to input len of %d", resultLen, inputLen)
		}

		for k, v := range input {
			resultValue, ok := result[k]
			if !ok {
				t.Errorf("fn was not called with key %v", k)
			}
			if resultValue.value != input[k] {
				t.Errorf("fn was called with unexpected value %v for key %v.. expected value %v", resultValue.value, k, v)
			}
			if !withinDuration(startTime, resultValue.calledAt, allowedCallDelta) {
				t.Errorf("expected fn to be called within %v of startTime %v but was called at %v", allowedCallDelta, startTime, resultValue.calledAt)
			}
		}
	})
}

func TestForEachKeyValueMax(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		input := map[string]int{}
		ForEachKeyValueMax(input, 1, func(_ string, _ int) {
			t.Errorf("unexpected fn call for empty map")
		})
	})
	t.Run("call fn for each key value with max concurrency", func(t *testing.T) {
		input := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
			"d": 4,
			"e": 5,
		}
		result := orderedMap{}
		mu := sync.Mutex{}
		startTime := time.Now()
		allowedCallDelta := time.Millisecond * 10
		fnExecutionTime := time.Millisecond * 50
		ForEachKeyValueMax(input, 1, func(key string, value int) {
			mu.Lock()
			result.Add(key, orderedMapValue{
				value:    value,
				calledAt: time.Now(),
			})
			mu.Unlock()
			time.Sleep(fnExecutionTime)
		})

		if resultLen, inputLen := len(result.m), len(input); resultLen != inputLen {
			t.Errorf("result len of %d is not equal to input len of %d", resultLen, inputLen)
		}

		for k, v := range input {
			resultValue, ok := result.m[k]
			if !ok {
				t.Errorf("fn was not called with key %v", k)
			}
			if resultValue.value != input[k] {
				t.Errorf("fn was called with unexpected value %v for key %v.. expected value %v", resultValue.value, k, v)
			}
		}

		for idx, value := range result.s {
			timeToAdd := time.Duration(idx) * fnExecutionTime
			expectedCallTime := startTime.Add(timeToAdd)
			if !withinDuration(expectedCallTime, value.value.calledAt, allowedCallDelta) {
				t.Errorf("expected fn to be called within %v of expectedTime %v but was called at %v.. difference of %v", allowedCallDelta, expectedCallTime, value.value.calledAt, expectedCallTime.Sub(value.value.calledAt).String())
			}
		}
	})
}
