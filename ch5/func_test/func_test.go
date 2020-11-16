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

func sum(ops ...int) int {
	s := 0
	for _, i := range ops {
		s += i
	}
	return s
}
func TestVarParam(t *testing.T) {
	s := sum(12, 34, 56, 78, 90)
	t.Log(s)
}

func clear(t *testing.T) {
	t.Log("clear resources")
}
func TestDefer(t *testing.T) {
	defer clear(t)
	t.Log("Start")
	panic("err")
	t.Log("End")
}
