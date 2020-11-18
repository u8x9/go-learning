package series

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}

func GetFibonacciSerie(n int) ([]int, error) {
	if n < 2 || n > 100 {
		return nil, errors.New("n必须介于2~100之间")
	}
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret, nil
}
