package share_memory_test

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second)
	t.Log("counter =", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	counter := 0
	var lock sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			lock.Lock()
			defer func() {
				wg.Done()
				lock.Unlock()
			}()
			counter++
		}()
	}
	wg.Wait()
	t.Log("counter =", counter)
}
