package package_test

import (
	"testing"

	"github.com/u8x9/go-learning/ch8/series"
)

func TestGetFibonacciSerie(t *testing.T) {
	fib, err := series.GetFibonacciSerie(10)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(fib)
}
