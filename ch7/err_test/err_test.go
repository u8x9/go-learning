package err_test

import (
	"errors"
	"testing"
)

var lessThanTwoError = errors.New("n 必须大于2")
var largerThenHuanderError = errors.New("n 必须小于100")

func getFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, lessThanTwoError
	}
	if n > 100 {
		return nil, largerThenHuanderError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	fib, err := getFibonacci(-10)
	if err != nil {
		if err == lessThanTwoError {
			t.Log("foo")
		}
		if err == largerThenHuanderError {
			t.Log("bar")
		}
		t.Fatal(err)
	}
	t.Log(fib)
}
