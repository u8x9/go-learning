package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFunc(t *testing.T) {
	s := "a,b,c"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	i, _ := strconv.Atoi("10")
	t.Log(20 + i)
}
