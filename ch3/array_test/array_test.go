package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int // 初始化为零值
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 3, 4, 5}
	t.Log(arr[0], arr[1], arr[2])
	t.Log(arr1)
	t.Log(arr2)
}
func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 3, 5, 7, 9}
	for idx, v := range arr {
		t.Logf("arr[%d] = %v\n", idx, v)
	}
}

func TestArraySection(t *testing.T) {
	a := [...]int{111, 222, 333, 444}
	t.Log(a[1:2])
	t.Log(a[:2])
	t.Log(a[2:])
	t.Log(a[:])
	t.Log(a[1:3])
}
