package util_anyone_reply_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func firstResponse() string {
	numOfRunner := 10
	ch := make(chan string)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	// runtime.NumGoroutine() 当前系统中的协程数
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(firstResponse())
	time.Sleep(time.Second)
	t.Log("After:", runtime.NumGoroutine())
}
