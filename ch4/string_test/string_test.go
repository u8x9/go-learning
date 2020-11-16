package string_test

import (
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 初始化为默认零值 ""
	s = "hello"
	t.Log(len(s))
	//s[1] = '3' // 错误string 是不可变的byte slice
	s = "\xE4\xB8\xA5" // 可以存储任何二进制
	//s = "\xE4\xB8\xB5\xFF"
	t.Log(s)
	s = "中"
	t.Log(len(s)) // byte 数

	c := []rune(s)
	t.Log(len(c))
	//t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("%s unicode %x", s, c[0])
	t.Logf("%s utf8 %x", s, s)
}
func TestStringToRune(t *testing.T) {
	s := "锄禾日当午"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}
