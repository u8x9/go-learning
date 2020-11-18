package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "[SRV] 完毕"
}

func asyncService() chan string {
	retch := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("[ASGR-1] 准备返回结果")
		retch <- ret
		fmt.Println("[ASGR-2] service退出")
	}()
	return retch
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-asyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Fatal("超时")
	}
}
