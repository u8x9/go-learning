package type_test

import "testing"

type MyInt int64

// TestImplicit 隐式类型转换
func TestImplicit(t *testing.T) {
	// 不支持任何类型的隐式转换，包括类型别名之间的隐式转换
	var a int = 1
	var b int64
	b = int64(a)
	t.Log(a, b)

	var c MyInt
	c = MyInt(b)
	t.Log(b, c)
}

func TestPointer(t *testing.T) {
	// go 语言中的指针不能进行运算
	a := 1
	p := &a
	t.Log(a, p)
	t.Logf("%T %T", a, p)
}

func TestString(t *testing.T) {
	// go 语言中的字符串，不是引用类型，默认值是""，而不是 nil
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
