package groutine_test

import (
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			t.Log(i)
		}(i)
		// // 错误示例
		// go func() {
		// 	t.Log(i)
		// }()
	}
	time.Sleep(time.Millisecond * 50)
}
