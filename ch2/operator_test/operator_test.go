package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writeable
	Executable
)

// TestCompareArray 用 `==` 比较数组
// 相同维数且含有相同个数元素的数组才可以比较
// 每个元素都相同的才相等
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c) // invalid operation: a == c (mismatched types [4]int and [5]int)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	a := 7
	a = a &^ Readable // 清除a的Readable
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
