package parallel

import "sync"

//ForEachKeyValue call iteratee concurrently for each key value pair in m with no limit on concurrency
func ForEachKeyValue[K comparable, V any](m map[K]V, iteratee func(key K, value V)) {
	ForEachKeyValueMax(m, 0, iteratee)
}

//ForEachKeyValueMax call iteratee concurrently for each key value pair in m with max limit on concurrency
func ForEachKeyValueMax[K comparable, V any](m map[K]V, max int, iteratee func(key K, value V)) {
	var wg sync.WaitGroup
	wg.Add(len(m))

	var ch chan struct{}
	if max != 0 {
		ch = make(chan struct{}, max)
	}

	for k, v := range m {

		if max != 0 {
			ch <- struct{}{}
		}

		go func(_k K, _v V) {
			defer func() {
				if max != 0 {
					<-ch
				}
				wg.Done()
			}()

			iteratee(_k, _v)
		}(k, v)

	}

	wg.Wait()
}
