package cancel_by_close

import (
	"testing"
	"time"
)

type notifyChan chan struct{}

func isCancelled(ch notifyChan) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
func cancel1(ch notifyChan) {
	ch <- struct{}{}
}
func cancel2(ch notifyChan) {
	close(ch)
}

func TestCancel(t *testing.T) {
	ch := make(notifyChan, 0)
	for i := 0; i < 5; i++ {
		go func(i int, ch notifyChan) {
			for {
				if isCancelled(ch) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			t.Log(i, "Cancelled")
		}(i, ch)
	}
	//cancel1(ch)
	cancel2(ch)
	time.Sleep(time.Second)
}
