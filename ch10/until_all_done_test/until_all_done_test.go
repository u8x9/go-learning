package until_all_done_test

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
func allResponse() []string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := make([]string, 0, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		finalRet = append(finalRet, <-ch)
	}
	return finalRet
}

func TestResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	for _, r := range allResponse() {
		t.Log(r)
	}
	time.Sleep(time.Second)
	t.Log("After:", runtime.NumGoroutine())
}
