package map_ext_test

import "testing"

type mapValue func(int) int

func TestMapWithFunValue(t *testing.T) {
	m := map[int]mapValue{}
	m[1] = func(a int) int { return a }
	m[2] = func(a int) int { return a * a }
	m[3] = func(a int) int { return a * a * a }
	for k, v := range m {
		t.Log(k, v(123))
	}
}
func TestMapForSet(t *testing.T) {
	mySet := map[string]bool{}
	mySet["foo"] = true
	mySet["bar"] = true

	if mySet["foo"] {
		t.Log("foo is exists")
	} else {
		t.Log("foo is NOT exists")
	}

	if mySet["bar"] {
		t.Log("bar is exists")
	} else {
		t.Log("bar is NOT exists")
	}

	if mySet["foobar"] {
		t.Log("foobar is exists")
	} else {
		t.Log("foobar is NOT exists")
	}
}
