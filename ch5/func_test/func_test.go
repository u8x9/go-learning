package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10), rand.Intn(20)
}

func timeSpend(inner func(int) int) func(int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(a int) int {
	time.Sleep(time.Second * 2)
	return a
}

func TestReturnMultiValues(t *testing.T) {
	x, y := returnMultiValues()
	t.Log(x, y)

	a := timeSpend(slowFunc)(123)
	t.Log(a)
}
