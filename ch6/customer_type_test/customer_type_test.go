package customer_type_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type intConv func(int) int

func timeSpent(inner intConv) intConv {
	return func(i int) int {
		start := time.Now()
		ret := inner(i)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func TestTimeSpent(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	go t.Log(timeSpent(func(i int) int {
		defer wg.Done()
		time.Sleep(time.Second)
		return i
	})(123))
	go t.Log(timeSpent(func(i int) int {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		return i * i
	})(123))
	go t.Log(timeSpent(func(i int) int {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		return i * i * i
	})(123))
	wg.Wait()
}
